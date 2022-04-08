package git

import (
	"fmt"
	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	http "github.com/go-git/go-git/v5/plumbing/transport/http"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"io/ioutil"
	"os"
)

var storer *memory.Storage
var fs billy.Filesystem

const gitFoldername = "catalogs"

func CommitAndPush(remoteName string, username string, password string, repo string, cPath string, removedFile string, dbname string) error {
	storer = memory.NewStorage()
	fs = memfs.New()
	auth := &http.BasicAuth{
		Username: username,
		Password: password,
	}
	repository := repo
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
	items, _ := ioutil.ReadDir(cPath + string(os.PathSeparator) + dbname)
	for _, item := range items {
		if len(removedFile) > 0 {
			remFile := gitFoldername + string(os.PathSeparator) + dbname + string(os.PathSeparator) + item.Name() + string(os.PathSeparator) + removedFile
			w.Remove(remFile)
		}
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(cPath + string(os.PathSeparator) + dbname + string(os.PathSeparator) + item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {

					txtfiles, _ := ioutil.ReadFile(cPath + string(os.PathSeparator) + dbname + string(os.PathSeparator) + item.Name() + string(os.PathSeparator) + subitem.Name())

					filePath := gitFoldername + string(os.PathSeparator) + dbname + string(os.PathSeparator) + item.Name() + string(os.PathSeparator) + subitem.Name()

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
		RemoteName: remoteName,
		Auth:       auth,
	})
	if err != nil {
		return err
	}
	return nil
}
