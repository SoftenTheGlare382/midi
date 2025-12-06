package midi

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestEventCopy(t *testing.T) {
	event := &Event{}
	event.TimeSignature = &TimeSignature{}
	event.SmpteOffset = &SmpteOffset{}
	if !reflect.DeepEqual(event, event.Copy()) {
		t.Fatal(errors.New("Expected copy to be equal"))
	}
}

func TestEventString(t *testing.T) {
	event := Event{}
	expect := "Ch 0 @ 0 (0) \t0X0"
	str := event.String()
	if str != expect {
		t.Errorf("Expected '%s' got '%s'", expect, str)
	}
}

func TestProgramChangeEncoding(t *testing.T) {
	// Test ProgramChange event encoding conforms to MIDI protocol standards
	// This test verifies the fix for the bug where an extra byte was being written
	event := ProgramChange(0, 1)

	// Encode the event
	data, err := event.Encode()
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	// Verify encoding result follows MIDI standard
	expected := []byte{0x00, 0xC0, 0x01}
	if !bytes.Equal(data, expected) {
		t.Errorf("ProgramChange encoding mismatch. Expected %v, got %v", expected, data)
	}
}
