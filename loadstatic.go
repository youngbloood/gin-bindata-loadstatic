package loadstatic

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

var debug bool

type assetsFS struct {
	engine    *gin.Engine
	assetsDir func(string) ([]string, error)
	asset     func(string) ([]byte, error)
	prefix    string
	htmlMap   map[string][]byte
}

func NewAssetsFS(engine *gin.Engine, AssetsDir func(string) ([]string, error),
	Asset func(string) ([]byte, error),
	Prefix string) *assetsFS {
	return &assetsFS{
		engine:    engine,
		assetsDir: AssetsDir,
		asset:     Asset,
		prefix:    Prefix,
		htmlMap:   make(map[string][]byte, 1000),
	}
}

func (fs *assetsFS) LoadStatic() error {
	if fs == nil {
		return errors.New("staticfs is nil")
	}
	if fs.engine == nil {
		return errors.New("the engine is nil")
	}
	fs.isDebug()
	fs.loadAll("")
	fs.engine.SetHTMLTemplate(template.Must(fs.loadTHMLTemplate(fs.htmlMap)))
	return nil
}
func (fs *assetsFS) loadAll(prefix string) error {
	var prefixPath string
	if prefix == "" {
		files, err := fs.assetsDir(prefix)
		if err != nil {
			return err
		}
		if len(files) != 0 && len(files) == 1 {
			prefixPath = files[0]
			fs.loadAll(prefixPath)
		}
	} else {
		dirs, err := fs.assetsDir(prefix)
		if err != nil {
			return err
		}
		for _, filePath := range dirs {
			prefixPath = prefix + "/" + filePath

			files, _ := fs.assetsDir(prefixPath)
			if len(files) != 0 {
				fs.loadAll(prefixPath)
			} else {

				data, _ := fs.asset(prefixPath)
				if strings.Contains(prefixPath, ".html") {
					fs.htmlMap[prefixPath] = data
				} else {
					fs.engine.Handle("GET", fs.replacePrefix(prefixPath, fs.prefix), func(c *gin.Context) {
						fmt.Fprint(c.Writer, string(data))
					})
				}

			}
		}
	}

	return nil
}
func (fs *assetsFS) isDebug() {
	if ok := gin.IsDebugging(); ok {
		debug = true
	} else {
		debug = false
	}
}

func (fs *assetsFS) loadTHMLTemplate(html map[string][]byte) (*template.Template, error) {

	if debug {
		log.Println("Loading HTML template start...")
	}
	var t *template.Template
	for k, v := range html {
		//name:=k
		//name := filepath.Base(k)
		name := fs.replacePrefix(k, fs.prefix)
		if t == nil {
			t = template.New(name)
		}
		var tmpl = t.New(name)
		_, err := tmpl.Parse(string(v))
		if err != nil {
			return nil, err
		}
		if debug {
			log.Printf("Loading html template :%s\n", name)
		}
	}
	if debug {
		log.Println("Loading HTML template end...")
		log.Print("========================>>>>\n\n")
	}
	return t, nil
}
func (fs *assetsFS) replacePrefix(path, Prefix string) string {
	if Prefix == "" {
		return path
	}
	if strings.HasPrefix(path, Prefix) {
		if strings.HasSuffix(Prefix, "/") {
			return strings.Replace(path, Prefix, "", 1)
		} else {
			return strings.Replace(path, Prefix+"/", "", 1)
		}
	}
	return path
}
