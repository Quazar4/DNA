package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateGCContent(sequence string) float64 {
	totalLength := len(sequence)
	gcCount := strings.Count(sequence, "G") + strings.Count(sequence, "C")
	return float64(gcCount) / float64(totalLength) * 100
}

func reverseComplement(sequence string) string {
	complement := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	reversedSequence := reverseString(sequence)
	return complement.Replace(reversedSequence)
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter DNA sequence: ")
	sequence, _ := reader.ReadString('\n')
	sequence = strings.ToUpper(strings.TrimSpace(sequence))

	gcContent := calculateGCContent(sequence)
	fmt.Printf("GC content: %.2f%%\n", gcContent)

	reverseComp := reverseComplement(sequence)
	fmt.Println("Reverse complement:", reverseComp)
}
