package internal

import (
	"errors"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"strings"
)

type storageCategory string

const (
	localStorage   storageCategory = "localStorage"
	sessionStorage storageCategory = "sessionStorage"
)

var (
	storage  = app.LocalStorage
	category = localStorage
)

var ErrKeyNotExist = errors.New("key doesn't exist in the storage")

const (
	notePrefix       = "note:"
	unassignedString = "e8403a40-402f-11eb-b378-0242ac130002" // use a random uuid as a default value
)

func UseLocalStorage() {
	storage = app.LocalStorage
	category = localStorage
}

func UseSessionStorage() {
	storage = app.SessionStorage
	category = sessionStorage
}

func StorageGetString(k string) (string, error) {
	v := unassignedString
	err := storage.Get(k, &v)
	if err != nil {
		panic("Cannot use storage: " + err.Error())
	}

	if v == unassignedString {
		return "", ErrKeyNotExist
	} else {
		return v, nil
	}
}

func StorageListNoteId() ([]string, error) {
	length := app.Window().Get(string(category)).Get("length").Int()

	var ret []string
	for i := 0; i < length; i++ {
		k := app.Window().Get(string(category)).Call("key", i).String()
		if isNoteKey(k) {
			ret = append(ret, toNoteIdFromKey(k))
		}
	}

	return ret, nil
}

func StorageGetNote(noteId string) (Note, error) {
	v := Note{
		Title: unassignedString,
		Body:  unassignedString,
	}
	err := storage.Get(toNoteKey(noteId), &v)
	if err != nil {
		panic("Cannot use storage: " + err.Error())
	}

	if v.Title == unassignedString || v.Body == unassignedString {
		return Note{}, ErrKeyNotExist
	} else {
		return v, nil
	}
}

func StorageSetNote(noteId string, note Note) error {
	return storageSet(toNoteKey(noteId), note)
}

func storageSet(k string, v interface{}) error {
	return storage.Set(k, v)
}

func toNoteKey(noteId string) string {
	return notePrefix + noteId
}

func toNoteIdFromKey(k string) string {
	return k[5:]
}

func isNoteKey(k string) bool {
	return strings.Contains(k, notePrefix)
}
