package main

import(
	"fmt"
	"bufio"	
	"os"
	"strings"
) 


// Count frequency of each word in a sentence

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the sentence : ")
	sentence, _ := reader.ReadString('\n')

	word_freq := map[string]int{}

	for _, word := range strings.Fields(sentence){
		_, word_exists := word_freq[word]

		if word_exists {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}

	fmt.Println(word_freq)

	// for key, value := range word_freq{
	// 	fmt.Printf("%v : %v\n", key, value)
	// }
	
}