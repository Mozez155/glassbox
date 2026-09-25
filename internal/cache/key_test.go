// Copyright 2026 Glassbox Users
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"testing"
)

var validWASMBytes = []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00}

func TestNewCacheKey_EmptyToolVersion(t *testing.T) {
	_, err := NewCacheKey(validWASMBytes, KindCompilation, "", nil, nil)
	if err == nil {
		t.Fatal("expected error for empty toolVersion, got nil")
	}
}

func TestNewCacheKey_EmptyWASMBytes(t *testing.T) {
	_, err := NewCacheKey(nil, KindCompilation, "v1.0.0", nil, nil)
	if err == nil {
		t.Fatal("expected error for empty wasmBytes, got nil")
	}
}

func TestNewCacheKey_EmptyKind(t *testing.T) {
	_, err := NewCacheKey(validWASMBytes, "", "v1.0.0", nil, nil)
	if err == nil {
		t.Fatal("expected error for empty kind, got nil")
	}
}

func TestNewCacheKey_Valid(t *testing.T) {
	key, err := NewCacheKey(validWASMBytes, KindCompilation, "v1.0.0", nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key.ToolVersion != "v1.0.0" {
		t.Errorf("ToolVersion: want v1.0.0, got %s", key.ToolVersion)
	}
	if key.Kind != KindCompilation {
		t.Errorf("Kind: want %s, got %s", KindCompilation, key.Kind)
	}
	if key.ContentHash == "" {
		t.Error("ContentHash must not be empty")
	}
	if key.SchemaVersion != CacheKeyVersion {
		t.Errorf("SchemaVersion: want %d, got %d", CacheKeyVersion, key.SchemaVersion)
	}
}
