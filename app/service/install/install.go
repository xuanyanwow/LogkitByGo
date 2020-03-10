package install

import "os"

func IsInstall() bool {
	if IsExistLock() {
		return true
	}
	return false
}

func IsExistLock() bool {
	_, err := os.Stat("./public/install.lock")
	return err == nil || os.IsExist(err)
}
