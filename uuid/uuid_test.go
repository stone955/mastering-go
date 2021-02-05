package uuid

import (
	"testing"
)

func TestGenGoogleUUID(t *testing.T) {
	ID := GenGoogleUUID()
	t.Logf("Google UUID = %v\n", ID)
}

func BenchmarkGenGoogleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenGoogleUUID()
	}
}
