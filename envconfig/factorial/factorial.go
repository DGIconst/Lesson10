package factorial

func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n - 1)
}

func IterativeFactorial(number int) uint64 {  
    var result uint64 = 1  
    if number < 0 {  
     
    } else {  
     for i := 1; i <= number; i++ {  
      result *= uint64(i)  
     
     }  
    }  
    return result  
   }  