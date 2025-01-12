package executor

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/phillip-england/flint/src/generator/asset"
	"github.com/phillip-england/flint/src/generator/config"
	"github.com/phillip-england/flint/src/generator/response"
	"github.com/phillip-england/flint/src/generator/route"

	"github.com/phillip-england/purse"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

type Spark struct {
	Type string
}

func NewSpark() (*Spark, error) {
	executor := &Spark{
		Type: KeySpark,
	}
	return executor, nil
}

func (exe *Spark) Print() { fmt.Println(exe.Type) }
func (exe *Spark) Run() error {
	title := purse.Fmt(`
	#################################################
	##          ##  ###  ####   ######  ##        ###
	##  ##########  #########    #####  #####  ######
	##       #####  ###  ####  #  ####  #####  ######
	##  ##########  ###  ####  ##  ###  #####  ######
	##  ##########  ###  ####  ###  ##  #####  ######
	##  ##########  #########  ####  #  #####  ######
	##  ##########         ##  #####    #####  ######
	-------------------------------------------------
`)
	title = strings.ReplaceAll(title, " ", "$")
	title = strings.ReplaceAll(title, "#", " ")
	title = strings.ReplaceAll(title, "$", "#")
	fmt.Println(title)
	fmt.Println("Language-Agnostic Static Sites")
	fmt.Println("-------------------------------------------------")
	fmt.Println("🔥 sparking flint")
	fmt.Println("🗃️ searching for flint.json")
	conf, err := config.New()
	if err != nil {
		return err
	}
	fmt.Println("🔎 parsing routes from flint.json")
	routes, err := route.NewFromConfig(conf)
	if err != nil {
		return err
	}
	fmt.Println("🏹 making an http request to each route")
	responses, err := response.NewFromRoutes(conf.Host, routes)
	if err != nil {
		return err
	}
	fmt.Println("🔨 generating static html assests")
	assets, err := asset.NewFromResponses(responses, conf.Target)
	if err != nil {
		return err
	}
	fmt.Println("🗑️ removing " + conf.Out)
	if _, err := os.Stat(conf.Out); err == nil {
		err = os.RemoveAll(conf.Out)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		// do nothing
	} else {
		return err
	}
	fmt.Println("✍️ writing minified, static html to " + conf.Out)
	for _, asset := range assets {
		outPath := conf.Out + asset.Path
		dir := filepath.Dir(outPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		file, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.Write([]byte(asset.Html))
		if err != nil {
			return err
		}
	}
	fmt.Println("🖌️ copying over minified, static assests from " + conf.Static)
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("application/xml", xml.Minify)
	err = filepath.Walk(conf.Static, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			outPath := conf.Out + "/" + path
			input, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			ext := filepath.Ext(path)
			var contentType string
			switch ext {
			case ".css":
				contentType = "text/css"
			case ".js":
				contentType = "text/javascript"
			case ".svg":
				contentType = "image/svg+xml"
			case ".html", ".htm":
				contentType = "text/html"
			case ".xml":
				contentType = "application/xml"
			default:
				contentType = ""
			}
			var output []byte
			if contentType != "" {
				output, err = m.Bytes(contentType, input)
				if err != nil {
					return fmt.Errorf("failed to minify %s: %w", path, err)
				}
			} else {
				output = input
			}
			err = os.MkdirAll(filepath.Dir(outPath), 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory for %s: %w", outPath, err)
			}
			err = os.WriteFile(outPath, output, 0644)
			if err != nil {
				return fmt.Errorf("failed to write to %s: %w", outPath, err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = filepath.Walk(conf.Static, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {

		}
		return nil
	})
	fmt.Println("✏️ copying over the favicon from " + conf.Favicon + " to " + conf.Out)
	_, err = os.Stat(conf.Favicon)
	if os.IsNotExist(err) {
		// Favicon doesn't exist, just log a warning and continue to the next task
		fmt.Println("⚠️ Favicon does not exist at: " + conf.Favicon)
		// Do not return or exit, just move to the next step
	} else if err != nil {
		return fmt.Errorf("error checking favicon file: %w", err)
	} else {
		// If the favicon exists, copy it
		outPath := filepath.Join(conf.Out, "favicon.ico")
		sourceFile, err := os.Open(conf.Favicon)
		if err != nil {
			return fmt.Errorf("failed to open favicon file: %w", err)
		}
		defer sourceFile.Close()

		destinationFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("failed to create favicon destination file: %w", err)
		}
		defer destinationFile.Close()

		_, err = io.Copy(destinationFile, sourceFile)
		if err != nil {
			return fmt.Errorf("failed to copy favicon: %w", err)
		}

		fmt.Println("✅ Favicon successfully copied to: " + outPath)
	}

	intro := purse.Fmt(`
		📚 your assets have been bundled at %s
		🙏 thank you for using flint
		⭐ dont forget to give me star at: https://github.com/phillip-england/flint
	`, conf.Out)

	fmt.Println(intro)

	return nil
}
