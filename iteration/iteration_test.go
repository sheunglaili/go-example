package iteration

import "testing"

func TestRepeat(t *testing.T){
	t.Run("repeat 6 times",func (t *testing.T){
	   got := Repeat("a",6)
	   expect := "aaaaaa"
	   if(got != expect){
		   t.Errorf("expected %q but got %q",expect,got)
	   }
	})

	t.Run("repeat 7 times",func (t *testing.T){
		got := Repeat("b",7)
		expect := "bbbbbbb"
		if(got != expect){
			t.Errorf("expected %q but got %q",expect,got)
		}
	 })
}

func BenchmarkRepeat(b *testing.B){
	for i := 0 ; i < b.N ; i++ {
		Repeat("a",5)
	}
}