package portfolio

import (
	"fmt"
)

type Project struct {
	Id               int
	Title            string
	ShortDescription string
	Description      string
	ImgPath          string
	// 	Url      string
	Year string
	// 	Order    int
}

func NewProject(title string) (*Project, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	Id := 1
	imgPath := ""
	shortDescription := ""
	description := ""
	year := ""
	return &Project{Id, title, shortDescription, description, imgPath, year}, nil
}

type ProjectManager struct {
	projects []*Project
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{}
}

func (m *ProjectManager) Save(project *Project) {
	m.projects = append(m.projects, project)
}

func (m *ProjectManager) All() []*Project {
	projects := []*Project{
		m.Create(1),
		m.Create(2),
		m.Create(3),
		m.Create(4),
		m.Create(5),
		m.Create(6),
		m.Create(7),
		m.Create(8),
		m.Create(9),
		m.Create(10),
	}
	return projects
}

func (m *ProjectManager) Create(Id int) *Project {
	if Id == 1 {
		project := Project{1, "Somboon seafood", "Responsive website, Yii, Bootstrap", "", "somboon.png", "2014"}
		return &project
	} else if Id == 2 {
		project := Project{2, "1Advice Society", "Responsive website, Yii, Bootstrap", "", "1advice.png", "2013"}
		return &project
	} else if Id == 3 {
		project := Project{3, "Timesheet (Infra)", "Web-Based Timesheet, Codeigniter", "", "timesheet.png", "2011"}
		return &project
	} else if Id == 4 {
		project := Project{4, "Molite", "Website, Codeigniter, Javascript, Jquery", "", "molite.png", "2011"}
		return &project
	} else if Id == 5 {
		project := Project{5, "Raffini", "Flash, ActionScript", "", "raffini.png", "2010"}
		return &project
	} else if Id == 6 {
		project := Project{6, "Sinturak", "Flash, ActionScript", "", "sinturak.png", "2010"}
		return &project
	} else if Id == 7 {
		project := Project{7, "RF", "Flash, ActionScript", "", "rf.png", "2010"}
		return &project
	} else if Id == 8 {
		project := Project{8, "3D Race game", "3D Game, C++, Ogre Framework", "", "racegame.png", "2010"}
		return &project
	} else if Id == 9 {
		project := Project{9, "Kenkoon", "Flash, ActionScript", "", "kenkoon.png", "2010"}
		return &project
	} else if Id == 10 {
		project := Project{10, "Music Mixer", "Window App, C++", "", "mixer.png", "2010"}
		return &project
	}
	return nil
}

func (m *ProjectManager) GetProjectById(Id int) (*Project, bool) {
	if Id >= 1 && Id <= 10 {
		project := m.Create(Id)
		project.Description = m.GetDescriptionByProjectId(Id)
		return project, true
	}
	return nil, false
}

func (m *ProjectManager) GetDescriptionByProjectId(Id int) string {
	if Id == 1 {
		description := "<p>somboonseafood.com เว็บนี้ทำงานร่วมกับ <a title='practical' href='http://practicalstudio.wordpress.com/' target='_blank'>Practical Studio</a>&nbsp;ซึ่งทางทีม Practical จะออกแบบเว็บไซต์มาให้</p><p>เว็บนี้ทำให้มีการตั้งคำถามเกี่ยวกับเรื่องของการใช้ Custom Font คือลูกค้าต้องการใช้ font ที่มี license แล้วก็เกิดคำถามว่า ถ้าใช้ custom font แล้วมันจะไปอยู่ในฝั่ง client หรือเปล้่า จะทำให้เหมือนโดนขโมย font หรือเปล่า ยังไม่ได้ลองทำการทดลองดู ก็สรุปกันว่า งั้นใช้รูปแทนตัวอักษรกันไปก่อน&nbsp;</p><p>เว็บนี้ไม่มีการทำงานซับซ้อน ใช้ Yii framework เป็นหลัก มีระบบ backoffice ให้ admin ได้แก้ไขข้อมูลทั่วไป ในส่วน ui ก็ใช้ bootstrap ในการทำ responsive website</p>"
		return description
	} else if Id == 2 {
		description := "<p>เว็บนี้ทำให้กับอลิอันซ์ อยุธยา โดยมีน้องอ๊อฟจากทีม <a href='http://www.mojistudio.com/main/'>Moji</a>&nbsp; มาช่วยโค้ดด้วย และมีทีม <a href='http://www.click-again.com/'>Click-Again</a></p><p>เว็บนี้เป็นเว็บให้ผู้ใช้มาแชร์ข้อมูล แชรรูป แชร์เรื่องดีๆกัน ลักษณะก็เหมือน social website ทั่วไป มีการ follow ผู้ใช้คนอื่น มีการ Share การกด Like และเก็บเข้า Favorite ได้ ในส่วนของ backoffice นี่สุดๆ ลูกค้าต้องการการออก report เยอะมาก การจัดการข้อมูลต่างๆก็แทบไม่ต้องไปเข้า database เลย ยังดีที่ Yii ช่วยไว้เยอะ</p<p>เว็บนี้โจทย์อยู่ที่การทำ responsive website โดยทีมเราใช้ twitter bootstrap แต่ก็ยังดึงความสามารถของ bootstrap ได้ไม่เต็มที่ เพราะมาพบทีหลังว่า compiler ที่ใช้มันมีปัญหา (ใช้ lessphp)</p><p>ปัญหาที่เจอหลักๆเลย คือ เรื่องของ environment ที่ผมไม่ถนัด อย่าง database ลูกค้าใช้ MSSQL ซึ่งผมไม่ค่อยรู้เรื่องเท่าไหร่ ประกอบกับการที่จะต้อง migrate ข้อมูลจากเว็บไซต์เดิมที่เก็บข้อมูลเป็น tis620 มาใช้ในเว็บนี้ให้เป็น utf8 ก็ทำให้เกิดปัญหาหลายอย่างอยู่เหมือนกัน เนื่องจากเราไม่สามารถลง driver สำหรับ database ได้ ต้องใช้ driver เดิม ซึ่งมีข้อจำกัดอยู่มาก</p><p>แต่โดยรวมแล้ว ถือว่าได้ประสบการณ์มากกับงานนี้ และพอใจกับงานที่พัฒนา ถึงจะมีปัญหาบ้าง บั๊กบ้าง แต่ก็ถือว่าพัฒนาได้ไว งานก็ออกมาโอเค (มั้ง)</p>"
		return description
	} else if Id == 3 {
		description := "<p>เว็บนี้ทำให้กับบริษัท INFRA TECHNOLOGY SERVICE&nbsp;</p><p>เป็นเว็บสำหรับลงเวลาการทำงานของพนักงาน จัดการเกี่ยวกับข้อมูอของแต่ละโครงการ (Project) ข้อมูลลูกค้าที่ติดต่อกับบริษัท และสามารถออกรายงานต่างๆได้เป็นตาราง กราฟ รวมไปถึง export เป็น pdf เช่น การออกรายงานการทำงานของพนักงานในแต่ละโครงการ ในแต่ละประเภทการทำงาน การออกรายงานการใช้เวลาในแต่ละขั้นตอนการทำงานของแต่ละโครงการ เป็นต้น</p>"
		return description
	} else if Id == 4 {
		description := "<p>Molite เป็นอีกหนึ่งผลิตภัณฑ์จากบริษัทเคนคูน โดยเว็บไซต์นี้ลูกค่้าไม่ต้องการให้ใช้ flash แต่อยากได้การทำงานที่ใกล้เคียง&nbsp;</p><p>เว็บไซต์ประกอบด้วย 2 ส่วน front-end และ back-end <br> ในส่่วนของ front-end เน้นที่การแสดงผลจาก Javascript (Jquery) และ css เป็นหลัก<br>back-end เป็นส่วนให้ลูกค้าแก้ไขข้อมูล อัพโหลดรูปภาพ&nbsp;ซึ่งเป็นการทำงานของ &nbsp;server-side script เป็นหลัก ผมใช้ codeigniter framework ในการจัดการส่วนนี้</p>"
		return description
	} else if Id == 5 {
		description := "<p dir='ltr'>Raffini เป็นเว็บไซต์แสดงข้อมูลของแบรนด์จิวเวลลี่ของ RF Group โดยงานนี้ออกแบบโดย&nbsp;<a title='Practical Design Studio' href='http://practicalstudio.wordpress.com/' target='_blank'>Practical Studio</a></p><p dir='ltr'>Raffini เป็นเว็บไซต์พัฒนาด้วย Flash เป็น front end และมี back end ที่ให้ admin สามารถแก้ไขข้อมูลต่างๆ ซึ่งพัฒนาด้วย php&nbsp;โดย admin จะสามารถจัดการรูปภาพโดยการเพิ่ม แก้ไข ลบรูป เรียงลำดับรูปในแต่ละคอลเลคชั่น เพิ่มข่าวสาร ข้อมูลงานอีเว้นท์ต่างๆ</p>"
		return description
	} else if Id == 6 {
		description := "<p dir='ltr'>Sinturak เป็นเว็บไซต์แสดงข้อมูลของแบรนด์จิวเวลลี่ของ RF Group โดยงานนี้ออกแบบโดย&nbsp;<a title='Practical Design Studio' href='http://practicalstudio.wordpress.com/' target='_blank'>Practical Studio</a></p><p dir='ltr'>Sinturak&nbsp;เป็นเว็บไซต์พัฒนาด้วย Flash เป็น front end และมี back end ที่ให้ admin สามารถแก้ไขข้อมูลต่างๆ ซึ่งพัฒนาด้วย php&nbsp;โดย admin จะสามารถจัดการรูปภาพโดยการเพิ่ม แก้ไข ลบรูป เรียงลำดับรูปในแต่ละคอลเลคชั่น เพิ่มข่าวสาร ข้อมูลงานอีเว้นท์ต่างๆ</p>"
		return description
	} else if Id == 7 {
		description := "<p dir='ltr'>RF เป็นเว็บไซต์แสดงข้อมูลของแบรนด์จิวเวลลี่ของ RF Group โดยงานนี้ออกแบบโดย&nbsp;<a title='Practical Design Studio' href='http://practicalstudio.wordpress.com/' target='_blank'>Practical Studio</a></p><p dir='ltr'>RF เป็นเว็บไซต์พัฒนาด้วย Flash เป็น front end และมี back end ที่ให้ admin สามารถแก้ไขข้อมูลต่างๆ ซึ่งพัฒนาด้วย php&nbsp;โดย admin จะสามารถจัดการรูปภาพโดยการเพิ่ม แก้ไข ลบรูป เรียงลำดับรูปในแต่ละคอลเลคชั่น เพิ่มข่าวสาร ข้อมูลงานอีเว้นท์ต่างๆ</p>"
		return description
	} else if Id == 8 {
		description := "<p>3D Racegame เป็นเกมแข่งรถสามมิติ ลักษณะเหมือนเกมแข่งม้า ที่ให้ผู้เล่นทายว่าใครจะเป็นผู้ชนะ โดยผู้เล่นไม่สามารถบังคับเองได้ การแข่งกันเป็น AI ทั้งหมด</p><p>เกมนี้ผมพัฒนาด้วย c++ และ OGRE (Ogre ไม่ใช่เกมเอนจิ้นนะครับ เป็นกราฟฟิคเอนจิ้น) ในส่วนของงานออกแบบ ออกแบบโดย <a title='Click-Agian' href='http://www.click-again.com'>Click-Again</a>&nbsp;ทั้งหมด&nbsp;</p><p>สิ่งที่ได้เรียนรู้ใหม่ๆ จากงานนี้ คือเรื่องของ AI และการเขียนโค้ดติดต่อกับ PCB (ผมไม่แน่ใจว่าเรียกแบบนี้ถูกไหม)</p><p>ปล. รูปจะไม่ค่อยชัด ดูเป็น <a title='vdo ที่นี่ได้เลยครับ' href='http://www.youtube.com/watch?v=jnpgPo28qv4' target='_blank'>vdo&nbsp;ที่นี่ได้เลยครับ</a>&nbsp;(ต้องขอบคุณ Click-Again ที่อัพโหลดไว้)</p>"
		return description
	} else if Id == 9 {
		description := "<p>Kenkoon เป็นเว็บไซต์ที่แสดงผลงานและข้อมูลต่างๆของบริษัทเคนคูน งานนี้ผมพัฒนาต่อจากงาน flash เดิม โดยเพิ่ม back-end ให้พนักงานของเคนคูนสามารถเพิ่มคอลเล็คชั่นต่างๆเองได้ และมีการปรับปรุงข้อมูลและส่วนแสดงผลต่างๆตามที่ลูกค้าออกแบบมา</p>"
		return description
	} else if Id == 10 {
		description := "<p>Music Mixer เป็นโปรแกรมที่สามารถสร้างดนตรีด้วยการกำหนดรูปแบบของการเล่นเครื่องดนตรีและในผู้ใช้สามารถ import เสียง หรือ อัดเสียงเพื่อบันทึกได้ โดยโปรแกรมนี้ทำร่วมกับ <a title='Click-Again' href='http://www.click-again.com'>Click-Again</a>&nbsp;ทำให้กับ Cornetto</p><p>โปรแกรมนี้พัฒนาด้วย c++ เป็น window app.</p>"
		return description
	}
	return ""
}
