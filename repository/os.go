package repository

import "os"

func can_read(path string) (bool, error) {
    f, err := os.Open(path)
    if err == nil {
        f.Close()

        return true, nil
    }

    if os.IsPermission(err) {
        err = nil
    }

    return false, err
}
