package factorial

import "testing"

func TestInputZero(t *testing.T) {
    result := Factorial(0)

    if result != 1 {
        t.Errorf("expect 1 but result is %d", result)
    }
}

func TestInputTwo(t *testing.T) {
    result := Factorial(2)

    if result != 2 {
        t.Errorf("expect 2 but result is %d", result)
    }
}

func TestInputSix(t *testing.T) {
    result := Factorial(6)

    if result != 720 {
        t.Errorf("expect 720 but result is %d", result)
    }
}

func BenchmarkFactorial(b *testing.B) {
for i := 0; i < b.N; i++ {
    Factorial(20)
}
}

func BenchmarkIterativeFactorial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IterativeFactorial(20)
    }
}