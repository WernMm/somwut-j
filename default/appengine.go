// +build appengine
package somwutj

import "somwut-j/myserver"

func init() {
	myserver.RegisterHandlers()
}
