package repository

import (
    "os"
    "path/filepath"

    "github.com/siddontang/go/ioutil2"
    _ "gopkg.in/src-d/go-billy.v2"
    "gopkg.in/src-d/go-git.v4"
)

type Space struct {
    root string
}

func (self *Space) RepositoryExists(name string) bool {
    return ioutil2.FileExists(filepath.Join(self.root, name))
}

func (self *Space) OpenRepository(name string) (*git.Repository, error) {
    path := filepath.Join(self.root, name+".git")

    if !ioutil2.FileExists(path) {
        return nil, ErrRepositoryNotExist(name)
    }

    //fs := billy.

    return git.Open(nil, nil)
}

func OpenSpace(root string) (*Space, error) {
    if ioutil2.FileExists(root) {
        info, err := os.Stat(root)
        if err != nil {
            return nil, err
        }

        if !info.IsDir() {
            return nil, ErrPathNotDir(root)
        }

        readable, err := can_read(root)
        if err != nil {
            return nil, err
        }

        if !readable {
            return nil, ErrNotReadable(root)
        }
    } else {
        if err := os.Mkdir(root, 0755); err != nil {
            return nil, err
        }
    }

    res := &Space{
        root: root,
    }

    return res, nil
}
