package internal

import (
	"errors"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"strings"
)

var (
	storage = app.LocalStorage
)

var ErrKeyNotExist = errors.New("key doesn't exist in the storage")

const (
	notePrefix       = "note:"
	unassignedString = "e8403a40-402f-11eb-b378-0242ac130002" // use a random uuid as a default value
)

func UseLocalStorage() {
	storage = app.LocalStorage
}

func UseSessionStorage() {
	storage = app.SessionStorage
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
	var ret []string
	for i := 0; i < storage.Len(); i++ {
		k, err := storage.Key(i)
		if err != nil {
			panic("failed to iterate over local storage: " + err.Error())
		}
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
