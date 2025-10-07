package fs

import (
	"testing"

	"github.com/suzuki3jp/mn/internal/entity"
)

func TestParseRecord_Success(t *testing.T) {
	rec := []string{"1.5", "2.5", "A"}
	p, err := ParseRecord(2, rec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.X != 1.5 || p.Y != 2.5 || !p.IsA {
		t.Fatalf("unexpected point: %+v", p)
	}
}

func TestParseRecord_InvalidCols(t *testing.T) {
	rec := []string{"1.5", "2.5"}
	_, err := ParseRecord(3, rec)
	if err == nil {
		t.Fatalf("expected error for insufficient cols")
	}
	if _, ok := err.(*InvalidRowError); !ok {
		t.Fatalf("expected InvalidRowError, got %T", err)
	}
}

func TestParseRecord_InvalidNumber(t *testing.T) {
	rec := []string{"notnum", "2.5", "A"}
	_, err := ParseRecord(4, rec)
	if err == nil {
		t.Fatalf("expected error for invalid x")
	}
	if ire, ok := err.(*InvalidRowError); ok {
		if ire.Line != 4 {
			t.Fatalf("expected line 4 in error, got %d", ire.Line)
		}
	} else {
		t.Fatalf("expected InvalidRowError, got %T", err)
	}
}

func TestParseRecord_InvalidSeries(t *testing.T) {
	rec := []string{"1", "2", "C"}
	_, err := ParseRecord(5, rec)
	if err == nil {
		t.Fatalf("expected error for invalid series")
	}
	if ire, ok := err.(*InvalidRowError); ok {
		if ire.Line != 5 {
			t.Fatalf("expected line 5 in error, got %d", ire.Line)
		}
	} else {
		t.Fatalf("expected InvalidRowError, got %T", err)
	}
}

func TestParseRecords_AllGood(t *testing.T) {
	records := [][]string{
		{"x", "y", "series"},
		{"1", "2", "A"},
		{"3", "4", "B"},
	}
	pts, err := ParseRecords(records)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pts) != 2 {
		t.Fatalf("expected 2 points, got %d", len(pts))
	}
	if pts[0] != (entity.Point{X: 1, Y: 2, IsA: true}) {
		t.Fatalf("unexpected first point: %+v", pts[0])
	}
}
