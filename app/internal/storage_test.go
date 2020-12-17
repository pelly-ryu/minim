package internal

import (
	"errors"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageGet(t *testing.T) {
	Storage = app.SessionStorage
	testKey := "a"
	testValue := "123"

	_, err := StorageGetString(testKey)
	assert.NotNil(t, err)
	if !errors.Is(err, ErrKeyNotExist) {
		t.Fatal(err)
	}

	assert.Nil(t, StorageSet(testKey, testValue))
	v, err := StorageGetString(testKey)
	assert.Nil(t, err)
	assert.Equal(t, testValue, v)
}
