package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	// hash[:] => convert a fixed array to slice
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}
	return strings.Join(paths, "/")
}

type PathTransformFunc func(string) string

func DefaultPathTransformFunc(key string) string {
	return key
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (store *Store) WriteStream(folderName string, dataReader io.Reader) error {
	pathName := store.PathTransformFunc(folderName)
	fileName := "SomeFile.txt"
	//Here we are making folder with 0777 permission (Read + Write + Execute)
	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}
	//End

	buf := new(bytes.Buffer)

	io.Copy(buf, dataReader)

	//Creating file in a folder
	file, err := os.Create(pathName + "/" + fileName)
	if err != nil {
		return err
	}

	//Copying content in a file
	n, err := io.Copy(file, dataReader)
	if err != nil {
		return err
	}

	log.Printf("written (%d) bytes to disk", n)

	return nil
}
