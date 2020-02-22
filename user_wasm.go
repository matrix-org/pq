// Package pq is a pure Go Postgres driver for the database/sql package.

// +build js wasm

package pq


func userCurrent() (string, error) {
	return "none", nil
}
