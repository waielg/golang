# Golang

## Table of Contents 

I. Golang Training

## Golang Training
Mod1
- Created in 2009, version 1 launced in 2011
- concurrency features and network serves (API's, etc.)
- Case: Docker containers, Ebay, Twitch, etc. due to its advantages like powerful library, concurrency and networking features, native binaries with minimal dependencies
- Pros: Compiles to machine code, very fast compiler, Cryptography, JSON, flexible and scalable code, Support for testing "go test" Command
      - Staticaly typed language so tools for this are golint, gofmt, etc.
- Cons: Lack of generics (means have to use diff function for diff data type)
  
Mod2
- Go program basics: Package declation, Import packages, functions, variables, statements/expressions, comands
- style line comments (//) or style block comments (/**/) package comment
- Formatting: avoid lengthy style, simplicity, Ctrl+S, not many () or ;
- Data types: Primitive (integers, floating points)
-       byte = uint8, rune = int32
- modulus remainder: %, Boolean (bool) = true or false]
-       && = and, || = or, ! = not
- String: sequence of bytes, length = len, cannot change a string
- Constants cannot be changed (integers, strings, etc.)
- X += y means x= x+y
- Scoping includes local vs global variables
  
- PRINT ex.
      package main
      import "fmt"
      func main() {
      const pi = 3.13
      fmt.Println("here is", math.pi)
      }
  
- SWITCH ex.
      package main
      import "fmt"
      var workday = 3   //global scope
      func main() {
        switch workday {
         case 1:
                    fmt.Println("Monday")
         case 2:
                    fmt.Println("Tuesday")
         case 3:
                    fmt.Println("Wednesday")
        default:
                    fmt.Println("Take the wknd off")
  }
  
- ARRAYS ex.1
      package main
      import "fmt"
      func main() {
      arr := [4]int{0,1,2,3)
      arr[1] = 10
      fmt.Println("value of arr[1]:", arr[1])
      fmt.Println("Lenth of arr:" len(arr))
}

- ARRAYS ex.2
      package main
      import "fmt"
      func main() {
      var twoDimArr [3][2]int
      for i:=0; i<3; i++ {
            for j:=0; j<2; j++ {
                  twoDimArr[i][j] = i + j
            }
      }
      fmt.Println("two dimensional array values:", twoDimArr
      }

- SLICES ex.
      package main
      import "fmt"
      func main() {
        var sl []int OR
        sl := make([]int,4,3)   //make(T, len, cap) OR
        sl := []int{0,1,2,3}
        sl[1] = 10
        fmt.Println("value of sl[1]:", sl[1]  
        fmt.Println("values:", sl)
        fmt.Println("is zero slice nil?", sl == nil)
  }

- SLICES ex.2
      package main
      import "fmt"
      func main() {
        //append([]T, element1, element2)
        //append([]T, []T...)
        sl := []int{0,1,2,3}
        sl = append(sl, 4,5,6)    //append is like connect 2 slices OR
  }

- MAP ex.
      package main
      import "fmt"
      func main() {
        var prodPrice mpa[string]int
        fmt.Println(prodPrice) // output: map[]

        //declare and initialize wiht make
        temprice := make(map[string]int)
        temprice["convertible widget"] = 150
        prodPrice = temprice
        prodPrice["widget"] = 100
        fmt.Println(prodPrice)  // output: map[convertible widget: 150 widget: 100]
}

- POINTER ex.
   package main
   import "fmt"
      func main() {
              var := 123      //var <name> *T      //T is tye
              var ptr*int = &val
              fmt.Println(ptr)       // output: 0xc000032... aka memory address
              fmt.Println(*ptr)      // output: 123
  }

- POINTER ex.2
  package main
   import "fmt"
      func main() {
             //var <name> *T = new(T)      //T is type and &val is address
             var ptr *int = new(int)
              fmt.Println(ptr, *ptr)       // output: 0  0xc000032...
  }
        
- ELSE IF ex.
     package main
     import "fmt"
      func main() {
        temp := -10
        if temp < 0 {                          // or if temp:= -10; temp < 0 { ...
              fmt.Printlln("Below freezing")
        } else if temp == 0 {
              fmt.Println("At freezinng point")
  `     } else if temp > 0 {                      // or just else { ...      
              fmt.Prinln("Above freezing")
      }
  }

- SWITCH VERSION OF THIS ^^
       package main
        import "fmt"
            func main() {
        temp := -10
        switch {
              case temp < 0;
              fmt.Println("below freeezing")
              case temp == 0;
              fmt.Println("at freezing")
              defalut:
              fmt.Println("above freezing")
        }
        }      

- FOR LOOP ex.
        package main
        import "fmt"
        func main() {
              for i := 0; for i<5; i++ {
                    fmt.Println(i)
        `}
  }

- FOR LOOP BREAK ex.
        package main
        import "fmt"
        func main() {
              for i := 0; for i<5; i++ {
                    if == 2 {
  `                   break            // or Conitnue to isolate 1      // or go to 
                  }
           fmt.Println(i)
        `}
  }

- FOR LOOP OVER SLICE ex.                        // can loop over maps too
          package main
        import "fmt"
        func main() {
              sl := []int{10,20,30,40}
              for i := 0; i<len(sl); i++ {             // or for i, value := range sl
                    fmt.Println(i, sl[i])               // output is 10,20,30,40 for both
  `}
  }

- CONTAINS (string package) ex.
    package main
        import ("fmt" "strings")
        fun main() {
              fmt.PRintln(strings.Contains("Working with string functions", "phunctions")} 
                              //strings.Title (capitalizes)
                                //strings.Trim("__Working with string fucnitons__","_")}

Mod3
AVG FUNCTION
      package main
        import "fmt"
        func avg(x []float64) float64 {
              total:= 0.0      
              for , val:= rang x {
                    total := cal
                    }
                    return total / float64(len(x))
                    }

  
