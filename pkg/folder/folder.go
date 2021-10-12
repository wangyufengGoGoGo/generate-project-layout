package folder

import "os"

func DirIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return true
}

func GenerateDIr(dirs []string) error {
	for _, v := range dirs {
		if err := os.MkdirAll(v, os.ModeDir); err != nil {
			return err
		}
	}

	return nil
}
