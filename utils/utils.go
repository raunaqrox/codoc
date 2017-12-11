package utils

import "io/ioutil"

func ReadFile(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
