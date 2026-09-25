// Copyright 2026 Glassbox Users
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"testing"
)

func TestNewWASMCache_NilManagerPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic for nil manager, got none")
		}
		msg, ok := r.(string)
		if !ok {
			t.Fatalf("expected string panic value, got %T: %v", r, r)
		}
		if msg != "cache: NewWASMCache called with nil manager" {
			t.Errorf("unexpected panic message: %q", msg)
		}
	}()
	NewWASMCache(nil, 0, nil)
}
