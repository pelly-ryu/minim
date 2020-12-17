package internal

import (
	"errors"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

var Storage = app.LocalStorage

var ErrKeyNotExist = errors.New("key doesn't exist in the storage")

const (
	unassignedString = "e8403a40-402f-11eb-b378-0242ac130002" // use a random uuid as a default value
)

func StorageGetString(k string) (string, error) {
	v := unassignedString
	err := Storage.Get(k, &v)
	if err != nil {
		panic("Cannot use storage: " + err.Error())
	}

	if v == unassignedString {
		return "", ErrKeyNotExist
	} else {
		return v, nil
	}
}

func StorageGetNote(noteId string) (Note, error) {
	v := Note{
		Title: unassignedString,
		Body:  unassignedString,
	}
	err := Storage.Get(noteId, &v)
	if err != nil {
		panic("Cannot use storage: " + err.Error())
	}

	if v.Title == unassignedString || v.Body == unassignedString {
		return Note{}, ErrKeyNotExist
	} else {
		return v, nil
	}
}

func StorageSet(k string, v interface{}) error {
	return Storage.Set(k, v)
}
