package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageGet(t *testing.T) {
	UseSessionStorage()
	testKey := "a"
	testValue := "123"

	_, err := StorageGetString(testKey)
	assert.NotNil(t, err)
	if !errors.Is(err, ErrKeyNotExist) {
		t.Fatal(err)
	}

	assert.Nil(t, storageSet(testKey, testValue))
	v, err := StorageGetString(testKey)
	assert.Nil(t, err)
	assert.Equal(t, testValue, v)
}
