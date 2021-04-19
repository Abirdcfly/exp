// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"fmt"
	"time"
)

// Event holds the information about an event that occurred.
// It combines the event metadata with the user supplied labels.
type Event struct {
	Kind    Kind
	ID      uint64    // unique for this process id of the event
	Parent  uint64    // id of the parent event for this event
	At      time.Time // time at which the event is delivered to the exporter.
	Message string
	Labels  []Label
}

// Kind indicates the type of event.
type Kind byte

const (
	// UnknownKind is the default event kind, a real kind should always be chosen.
	UnknownKind = Kind(iota)
	// LogKind is a Labels kind that indicates a log event.
	LogKind
	// StartKind is a Labels kind that indicates a span start event.
	StartKind
	// EndKind is a Labels kind that indicates a span end event.
	EndKind
	// MetricKind is a Labels kind that indicates a metric record event.
	MetricKind
	// AnnotateKind is a Labels kind that reports label values at a point in time.
	AnnotateKind
)

// Find searches the labels of an event to see if one of them has the
// supplied key.
func (ev Event) Find(key string) Label {
	for _, l := range ev.Labels {
		if l.Key() == key {
			return l
		}
	}
	return Label{}
}

// String returns a string representation of the kind for printing.
func (k Kind) String() string {
	switch k {
	case LogKind:
		return "log"
	case StartKind:
		return "start"
	case EndKind:
		return "end"
	case MetricKind:
		return "metric"
	case AnnotateKind:
		return "annotate"
	default:
		return fmt.Sprint(byte(k))
	}
}