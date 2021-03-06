package data

import (
	"reflect"
	"testing"
)

func TestWordSet_Add(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		duplicate bool
	}{
		{"Add1", "dog", false},
		{"Add2", "exact", false},
		{"Add3", "publish", false},
		{"Duplicate1", "exact", true},
		{"Duplicate2", "publish", true},
	}
	set := NewWordSet()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevLen := set.Size()
			set.Add(tt.word)
			currentLen := set.Size()
			if tt.duplicate && prevLen != currentLen {
				t.Error("Set changed side after adding duplicate word")
			}
			if !tt.duplicate && (currentLen-prevLen) != 1 {
				t.Errorf("Set size changed unexpectedly. Expect: %d, Got: %d", prevLen+1, currentLen)
			}
		})
	}
}

func TestWordSet_AddAll(t *testing.T) {
	tests := []struct {
		name  string
		words []string
	}{
		{"AddAll1", []string{"dog"}},
		{"AddAll2", []string{"exact", "cunning"}},
		{"AddAll3", []string{"publish", "president", "transaction"}},
	}
	set := NewWordSet()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevLen := set.Size()
			set.AddAll(tt.words)
			currentLen := set.Size()
			if (currentLen - prevLen) != len(tt.words) {
				t.Errorf("Set size changed unexpectedly. Expect: %d, Got: %d", prevLen+len(tt.words), currentLen)
			}
		})
	}
}

func TestWordSet_Remove(t *testing.T) {
	tests := []struct {
		name string
		word string
	}{
		{"Del1", "dog"},
		{"Del2", "exact"},
		{"Del3", "publish"},
		{"Del4", "president"},
		{"Del5", "transaction"},
		{"Del6", "cunning"},
	}
	set := NewWordSet()
	set.AddAll([]string{"dog", "publish", "president", "transaction", "cunning"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set.Remove(tt.word)
			if set.Contains(tt.word) {
				t.Errorf("Word: %s supposed to be deleted from set, vut it still present", tt.word)
			}
		})
	}
}

func TestWordSet_Size(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"Size1", []string{"dog"}, 1},
		{"Size2", []string{"exact", "cunning"}, 2},
		{"Size3", []string{"publish", "president", "transaction"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewWordSet()
			set.AddAll(tt.words)
			if got := set.Size(); got != tt.want {
				t.Errorf("WordSet.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordSet_Contains(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{"Del1", "dog", true},
		{"Del2", "exact", false},
		{"Del3", "publish", true},
	}
	set := NewWordSet()
	set.AddAll([]string{"dog", "publish", "president", "transaction", "cunning"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := set.Contains(tt.word); got != tt.want {
				t.Errorf("WordSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordSet_Words(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"Size1", []string{"dog"}, []string{"dog"}},
		{"Size2", []string{"exact", "cunning"}, []string{"cunning", "exact"}},
		{"Size3", []string{"publish", "president", "transaction"},
			[]string{"president", "publish", "transaction"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewWordSet()
			set.AddAll(tt.words)
			if got := set.Words(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordSet.Words() = %v, want %v", got, tt.want)
			}
		})
	}
}
