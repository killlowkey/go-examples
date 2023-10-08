package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.TODO()

func Test_SetAndGet(t *testing.T) {
	err := rdb.Set(ctx, "test1", "tes1", 0).Err()
	assert.NoError(t, err)

	res := rdb.Get(ctx, "test1").String()
	assert.Equal(t, "test1", res)
}

func Test_SetEx(t *testing.T) {
	err := rdb.SetEx(ctx, "test2", "tes2", time.Microsecond*500).Err()
	assert.NoError(t, err)

	time.Sleep(time.Second)
	err = rdb.Get(ctx, "test2").Err()
	assert.Error(t, err)
}
