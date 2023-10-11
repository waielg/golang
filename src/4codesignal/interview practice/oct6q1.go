package main

func main() {

}

func solution(pattern string, source string) int {
    // Handle the special case when source is exactly 1 character
    if len(source) == 1 {
        return 1
    }


    // Initialize our variables, result is int (return) and match is bool
    result := 0

    for i := 0; i <= len(source)-len(pattern); i++ { // Iterate over max index
        match := true // Reset match for each iteration

        for j := 0; j < len(pattern); j++ { // Iterate over pattern string
            // Check if pattern iterations are equal to the source
            if (i+j >= len(source)) || ((pattern[j] == '0' && !Vowel(source[i+j])) ||
               (pattern[j] == '1' && Vowel(source[i+j]))) {
                match = false // If characters don't match the pattern, exit the loop
                break
            }
        }

        if match { // If match isn't false, add to the result
            result++
        }
    }
    return result
}

func Vowel(b byte) bool { // Must be byte to include indexes
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
