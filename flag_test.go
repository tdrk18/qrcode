package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"testing"
)

func TestValidateQuality(t *testing.T) {
	var result qrcode.RecoveryLevel

	result = validateLevel("low")
	if result != qrcode.Low {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", toString(result)))
	}
	result = validateLevel("medium")
	if result != qrcode.Medium {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", toString(result)))
	}
	result = validateLevel("high")
	if result != qrcode.High {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", toString(result)))
	}
	result = validateLevel("highest")
	if result != qrcode.Highest {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", toString(result)))
	}
	result = validateLevel("invalid")
	if result != qrcode.Medium {
		t.Fatal(fmt.Sprintf("failed: validateLevel() returns %s", toString(result)))
	}
}

func toString(level qrcode.RecoveryLevel) string {
	switch level {
	case qrcode.Low:
		return "low"
	case qrcode.Medium:
		return "medium"
	case qrcode.High:
		return "high"
	case qrcode.Highest:
		return "highest"
	default:
		return ""
	}
}
