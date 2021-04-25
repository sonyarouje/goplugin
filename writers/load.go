package writers

import (
	"io/ioutil"
	"os"
	"plugin"
	"strings"
)

func Load() ([]Writer, error) {
	var writers []Writer
	files, _ := ioutil.ReadDir("./writers/plugins/")
	for _, file := range files {
		if notAPlugin(file) {
			continue
		}

		//open the plugins from plugin folder.
		plug, err := plugin.Open("./writers/plugins/" + file.Name())
		if err != nil {
			return writers, err
		}
		p, err := plug.Lookup("P")
		if err != nil {
			return writers, err
		}

		var writer Writer
		writer, ok := p.(Writer)
		if !ok {
			continue
		}
		writers = append(writers, writer)
		// commandName := handler.CommandName()
		// commandMap[commandName] = handler
		// logger.Debug("Loaded " + commandName + " plugin from file " + file.Name())
	}
	return writers, nil
}

//function checks whether a file is a plugin. If the file is a directory then not a plugin
//a go file then not a plugin. Any file contains .so is a valid plugin
func notAPlugin(file os.FileInfo) bool {
	if file.IsDir() {
		return true
	}

	if !strings.Contains(file.Name(), ".so") {
		return true
	}
	return false
}
