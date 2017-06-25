package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/beego/mux"
	"github.com/nanjishidu/gomini"
	"github.com/nanjishidu/gomini/gocrypto"
)

func runWebApp(addr string) {
	mx := mux.New()
	mx.Handler("GET", "/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	mx.Get("/", new(Controllers).index)
	mx.Post("/list", new(Controllers).list)
	mx.Post("/show", new(Controllers).show)
	mx.Post("/edit", new(Controllers).edit)
	mx.Post("/save", new(Controllers).save)
	mx.Post("/export", new(Controllers).exportPost)
	mx.Get("/export/*", new(Controllers).exportGet)
	mx.Post("/rand", new(Controllers).rand)
	err := initSite()
	if err != nil {
		loger.Println("kpass service exception")
		return
	}
	loger.Printf("listen %s", addr)
	loger.Fatal(http.ListenAndServe(addr, mx))
}

type Controllers struct{}

func (c *Controllers) prepare(w http.ResponseWriter, r *http.Request, kpass, kpassfile, kcrypto string) (string, error) {
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}
	loger.Printf("Started %s %s for %s", r.Method, r.URL.Path, addr)
	if kpass == "" {
		return "", errors.New("kpass can not be empty")
	}
	if kpassfile == "" {
		return "", errors.New("kpassfile can not be empty")
	}
	kpassdfile := filepath.Join(kpassDir, kpassfile)
	if !gomini.IsExist(kpassdfile) {
		return "", errors.New(kpassdfile + " is not exist")
	}

	s, err := gomini.FileGetContent(kpassdfile)
	if err != nil {
		return "", errors.New("failed to read file")
	}
	hook, err := NewInstance(kcrypto)
	if err != nil {
		return "", err
	}
	acd, err := hook().Decrypt(kpass, s)
	if err != nil {
		return "", errors.New("password is incorrect")
	}
	return acd, nil
}

//index
func (c *Controllers) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

//list kpassfile
func (c *Controllers) list(w http.ResponseWriter, r *http.Request) {
	fnss, err := ScanDumpFile(kpassDir)
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	serveJson(w, 200, "", fnss)
}

//show kpassfile content
func (c *Controllers) show(w http.ResponseWriter, r *http.Request) {
	acd, err := c.prepare(w, r, r.FormValue("kpass"), r.FormValue("kpassfile"), r.FormValue("kcrypto"))
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	serveJson(w, 200, "success", acd)
}

//edit kpassfile content
func (c *Controllers) edit(w http.ResponseWriter, r *http.Request) {
	acd, err := c.prepare(w, r, r.FormValue("kpass"), r.FormValue("kpassfile"), r.FormValue("kcrypto"))
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	serveJson(w, 200, "success", acd)
}

//save kpassfile content
func (c *Controllers) save(w http.ResponseWriter, r *http.Request) {
	kpass := r.FormValue("kpass")
	kpassfile := r.FormValue("kpassfile")
	kcrypto := r.FormValue("kcrypto")
	_, err := c.prepare(w, r, kpass, kpassfile, kcrypto)
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	kpassBackupDir := filepath.Join(kpassDir, "backup")
	if !gomini.IsExist(kpassBackupDir) {
		gomini.Mkdir(kpassBackupDir)
	}
	hook, err := NewInstance(kcrypto)
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	bse, err := hook().Encrypt(kpass, r.FormValue("content"))
	if err != nil {
		serveJson(w, 500, "password is incorrect", nil)
		return
	}
	kpassdfile := filepath.Join(kpassDir, kpassfile)
	gomini.Rename(kpassdfile, filepath.Join(kpassBackupDir, kpassfile)+"_"+time.Now().Format("2006-01-02_15-04-05"))
	_, err = gomini.FilePutContent(kpassdfile, kpassSign+bse)
	if err != nil {
		serveJson(w, 500, "save failed", nil)
		return
	}
	serveJson(w, 200, "success", nil)

}

func (c *Controllers) rand(w http.ResponseWriter, r *http.Request) {
	var (
		defaultUppercase = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
		defaultLowercase = `abcdefghijklmnopqrstuvwxyz`
		defaultDigital   = `0123456789`
	)
	uppercase, _ := gomini.GetStrInt(r.FormValue("uppercase"), 1)
	lowercase, _ := gomini.GetStrInt(r.FormValue("lowercase"), 1)
	digital, _ := gomini.GetStrInt(r.FormValue("digital"), 1)
	special := r.FormValue("special")
	num, _ := gomini.GetStrInt(r.FormValue("num"), 16)
	count, _ := gomini.GetStrInt(r.FormValue("count"), 10)
	var ralphaNum string
	if uppercase == 1 {
		ralphaNum += defaultUppercase
	}
	if lowercase == 1 {
		ralphaNum += defaultLowercase
	}
	if digital == 1 {
		ralphaNum += defaultDigital
	}
	if special != "" {
		ralphaNum += special
	} else {
		ralphaNum += "!@#$%^&*."
	}
	if count <= 0 || count >= 10 {
		count = 10
	}
	var resp []string
	for i := 0; i <= count; i++ {
		rcb := RandomCreateBytes(num, []byte(ralphaNum)...)
		resp = append(resp, string(rcb))
	}
	serveJson(w, 200, "", "```"+strings.Join(resp, "\r\n")+"```")
}

//export kpassfile
func (c *Controllers) exportPost(w http.ResponseWriter, r *http.Request) {
	kpassfile := r.FormValue("kpassfile")
	_, err := c.prepare(w, r, r.FormValue("kpass"), kpassfile, r.FormValue("kcrypto"))
	if err != nil {
		serveJson(w, 500, err.Error(), nil)
		return
	}
	ekey := gocrypto.Md5(gomini.GetInt64Str(time.Now().UnixNano()))
	emap[ekey] = kpassfile
	serveJson(w, 200, "success", ekey)
	return
}
func (c *Controllers) exportGet(w http.ResponseWriter, r *http.Request) {
	ekey := mux.Param(r, ":splat")
	if kpassfile, ok := emap[ekey]; ok {
		download(w, r, filepath.Join(kpassDir, kpassfile))
		delete(emap, ekey)
	}
	return
}

func serveJson(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	datas := map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	content, err := json.Marshal(datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(content)
}

// Download forces response for download file.
// it prepares the download response header automatically.
// https://github.com/astaxie/beego/blob/master/context/output.go
func download(w http.ResponseWriter, r *http.Request, file string, filename ...string) {
	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Type", "application/octet-stream")
	if len(filename) > 0 && filename[0] != "" {
		w.Header().Set("Content-Disposition", "attachment; filename="+filename[0])
	} else {
		w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(file))
	}
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Pragma", "public")
	http.ServeFile(w, r, file)
	return
}
