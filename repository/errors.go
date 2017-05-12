package repository

import (
    "fmt"
)

type ErrPathNotDir string

func (self ErrNErrPathNotDirotExist) Error() string {²
    return fmt.Sprintf("Path [%s] is not a directory", string(self))
}

type ErrNotReadable string

func (self ErrNotReadable) Error() string {
    return fmt.Sprintf("Path [%s] cannot be read", string(self))
}

type ErrRepositoryNotExist string

func (self ErrRepositoryNotExist) Error() string {
    return fmt.Sprintf("Repository [%s] does not exist", string(self))
}

type ErrRepositoryAlreadyExists string

func (self ErrRepositoryAlreadyExists) Error() string {
    return fmt.Sprintf("Repository [%s] already exists", string(self))
}
