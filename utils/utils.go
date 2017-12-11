package utils

import "io/ioutil"

func ReadFile(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func WriteFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
