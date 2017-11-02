package myserver

import (
	"appengine"
	"appengine/channel"
	"appengine/datastore"
	"appengine/user"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"somwut-j/portfolio"
	"strconv"
	"time"
)

const PathPrefix = "/rest/portfolio/"
const PathPrefixDict = "/rest/dict/"
const PathPrefixChat = "/rest/channel/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix+"all", portfolioHandler)
	r.HandleFunc(PathPrefix+"{id}", projectHandler)
	r.HandleFunc(PathPrefixDict+"saveword", saveDictHandler)
	r.HandleFunc(PathPrefixDict+"searchword", searchDictHandler)
	r.HandleFunc(PathPrefixChat, getChannelHandler)

	http.Handle("/rest/", r)
}

func getChannelHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	roomkey := r.FormValue("roomkey")
	tok, err := channel.Create(c, u.ID+roomkey)
	if err != nil {
		// http.Error(w, "Couldn't create Channel", http.StatusInternalServerError)
		// c.Errorf("channel.Create: %v", err)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, tok)
	fmt.Fprint(w, "hello", u.ID)
}

type badRequest struct{ error }

type notFound struct{ error }

type DictWord struct {
	Id      int64
	Meaning string
	Word    string
	Created time.Time
}

var projectMgr = portfolio.NewProjectManager()

func decodeDictWord(r io.ReadCloser) (*DictWord, error) {
	defer r.Close()
	var dictWord DictWord
	err := json.NewDecoder(r).Decode(&dictWord)
	return &dictWord, err

}

func defaultDictWord(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "DictWord", "default", 0, nil)
}

func parseID(r *http.Request) (int64, error) {
	txt, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("project id not found")
	}
	return strconv.ParseInt(txt, 10, 0)
}

func search(r *http.Request) ([]DictWord, error) {
	c := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var d DictWord
	decoder.Decode(&d)
	q := datastore.NewQuery("DictWord")
	log.Println(d.Word)
	log.Println(d.Word != "")
	if d.Word != "" {
		q = q.Filter("Word =", d.Word)
	}
	q = q.Order("Word")
	var dictWords []DictWord
	_, err := q.GetAll(c, &dictWords)

	return dictWords, err
}

func (dictWord *DictWord) key(c appengine.Context) *datastore.Key {
	log.Println("key")
	log.Println(dictWord.Id)
	log.Println(dictWord.Word)

	if dictWord.Id != 0 {
		return datastore.NewKey(c, "DictWord", "", dictWord.Id, defaultDictWord(c))
	} else {
		dictWords := []DictWord{}
		ks, err := datastore.NewQuery("DictWord").Filter("Word =", dictWord.Word).Order("Created").GetAll(c, &dictWords)
		if err != nil {
			log.Println(err)
			return nil
		}
		log.Println("len ", len(ks))
		if len(ks) > 0 {
			return datastore.NewKey(c, "DictWord", "", ks[0].IntID(), defaultDictWord(c))
		} else {
			dictWord.Created = time.Now()
			return datastore.NewIncompleteKey(c, "DictWord", defaultDictWord(c))
		}
	}
}

func (dictWord *DictWord) save(c appengine.Context) (*DictWord, error) {
	k, err := datastore.Put(c, dictWord.key(c), dictWord)
	if err != nil {
		return nil, err
	}

	dictWord.Id = k.IntID()
	return dictWord, nil
}

func saveDictHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("saveDictHandler")
	c := appengine.NewContext(r)
	log.Println(r.Body)
	dictWord, err := decodeDictWord(r.Body)
	log.Println(dictWord.Word)
	if err != nil {
		fmt.Println(err)
		return
	}
	dictWord.save(c)

	log.Println(dictWord.Word)

	enc := json.NewEncoder(w)
	enc.Encode(dictWord)
}

func searchDictHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("searchDictHandler")

	dictWords, err := search(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(dictWords)

	// resJson, err := json.Marshal(dictWords)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Fprint(w, resJson)
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Projects []*portfolio.Project
	}
	res.Projects = projectMgr.All()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		log.Println(err)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	log.Println("Project is ", id)
	if err != nil {
	}

	project, ok := projectMgr.GetProjectById(int(id))
	log.Println("Found", ok)

	if !ok {
	}

	var res struct {
		Project *portfolio.Project
	}
	res.Project = project

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		log.Println(err)
	}
}
