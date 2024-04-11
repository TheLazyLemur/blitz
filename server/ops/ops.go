package ops

import (
	"fmt"

	"github.com/TheLazyLemur/blitz/server/store"
)

func Set(s store.Storer, k, v string) error {
    fmt.Println("SET", k, v)
    return s.Set(k, v)
}

func Get(s store.Storer, k string) (string, error) {
    fmt.Println("GET", k)
    return s.Get(k)
}
