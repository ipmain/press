//filename=ipcloud/press.go
//contact=hopley@ipcloud.net
//updatedAt=20160527

package press

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func init() {
	log.Println(" in press.init()")
}

var DEBUG = true

//Press() template File with the path for the specific file and to include  header.tmpl and footer.tmpl
func Press(tF string, v interface{}) (string, error) { //TOOD:(hopley) - handles nil But needs(confirm) a second var ...
	if DEBUG {
		log.Printf("[ipcloud Press(%s,Interface{})] - //not showing Interface{} details ...\n ", tF)
		log.Printf("[ipcloud Press()] - template file=%s\n", tF)
	}
	err := checkFileExists(tF)
	if err != nil {
		log.Printf("[ERROR] File=%s - may not exist.", tF)
		return "[ERROR] File=" + tF + " may not exist.", err
	}
	//dh - here just the name ...
	tT := template.New(filepath.Base(tF))
	tT, err = tT.ParseFiles(tF)
	if err != nil {
		log.Printf("[ERROR] parsing template - %s\n", err)
	}
	var htm bytes.Buffer
	err = tT.Execute(&htm, v)
	if err != nil {
		log.Printf("[ERROR] template execution error - %s\n", err)
	}
	h := htm.String()
	return h, nil
}

//PressPage() template File with content wrapped in header.tmpl and footer.tmpl from path.
func PressPage(tF string, v interface{}) (string, error) {
        //TODO:(hopley) - need a better way to get docRoot, assume heaer and footer .tmpl with tF else a default ...
        tPath := filepath.Dir(tF)
        log.Printf(" ~~ in PressPage() where tPath=%s\n", tPath)
        heade := tPath + "/" + "header.tmpl" //TODO:(hopley) - press.init() check for footer, header ; some help notes ...
        foote := tPath + "/" + "footer.tmpl"
        c, err := Press(heade, "/head")
        chk(err)
        h := c
        c, err = Press(tF, v)
        chk(err)
        h += c
        c, err = Press(foote, "/login")
        chk(err)
        h += c
        return h, nil
}


func checkFileExists(f string) error {
	//TODO:(hopley) - func PressPage(string, interface) string {  ... }
	_, err := os.Stat(f)
	if err != nil {
		return err
	}
	return nil
}

func chk(err error) {
	if err != nil {
		log.Fatal("ERROR: %s\n", err)
	}
}
