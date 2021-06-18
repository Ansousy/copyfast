package tuto

import (
	"io/fs"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if !d.IsDir() {
		println(s)
	}
	return nil
}

func main() {
	filepath.WalkDir("..", walk)
}

/*err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}

libRegEx, e := regexp.Compile("^.+\\.(dylib)$")
if e != nil {
    log.Fatal(e)
}

e = filepath.Walk("/usr/lib", func(path string, info os.FileInfo, err error) error {
    if err == nil && libRegEx.MatchString(info.Name()) {
        println(info.Name())
    }
    return nil
})
if e != nil {
    log.Fatal(e)
}
func processed(fileName string, processedDirectories []string) bool {
    for i := 0; i < len(processedDirectories); i++ {
        if processedDirectories[i] != fileName {
            continue
        }
        return true
    }
    return false
}

func listDirContents(path string, dirs []string) {
    files, _ := ioutil.ReadDir(path)

    for _, f := range files {
        var newPath string
        if path != "/" {
            newPath = fmt.Sprintf("%s/%s", path, f.Name())
        } else {
            newPath = fmt.Sprintf("%s%s", path, f.Name())
        }

        if f.IsDir() {
            if !processed(newPath, dirs) {
                dirs = append(dirs, newPath)
                listDirContents(newPath, dirs)
            }
        } else {
            fmt.Println(newPath)
        }
    }
}
*/
