package tes

import (
	"fmt"
	"github.com/prometheus/common/model"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	a:= model.Alert{
		Labels:       nil,
		Annotations:  nil,
		StartsAt:     time.Time{},
		EndsAt:       time.Time{},
		GeneratorURL: "",
	}

	fmt.Println(a)
}
