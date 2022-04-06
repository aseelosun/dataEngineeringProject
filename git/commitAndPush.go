package git

import (
	"fmt"
	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"io/ioutil"
)

var storer *memory.Storage
var fs billy.Filesystem

func CommitAndPush(removedFile string, dbname string) error {
	storer = memory.NewStorage()
	fs = memfs.New()
	auth := &http.BasicAuth{
		Username: "aseelosun",
		Password: "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",
	}

	repository := "https://github.com/aseelosun/data_ddl"
	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  repository,
		Auth: auth,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}
	items, _ := ioutil.ReadDir("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname)
	for _, item := range items {
		if len(removedFile) > 0 {
			remFile := "catalogs/" + dbname + "/" + item.Name() + "/" + removedFile
			w.Remove(remFile)
		}
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname + "\\" + item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {

					txtfiles, _ := ioutil.ReadFile("C:\\Users\\Trainee\\dataEngineeringProject\\catalogs\\" + dbname + "\\" + item.Name() + "\\" + subitem.Name())

					filePath := "catalogs/" + dbname + "/" + item.Name() + "/" + subitem.Name()

					newFile, err := fs.Create(filePath)
					if err != nil {
						return err
					}
					newFile.Write(txtfiles)
					newFile.Close()
					w.Add(filePath)

				}
			}
		} else {

			fmt.Println(item.Name())
		}
	}

	w.Commit("Files updated", &git.CommitOptions{})

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		return err
	}
	return nil
}
