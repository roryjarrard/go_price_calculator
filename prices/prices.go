package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/roryjarrard/go-price-calculator/conversion"
)

// TaxIncludedPriceJob represents a job to calculate tax-included prices.
type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// LoadData loads prices from a file and converts them to float64.
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting prices:", err)
		return
	}

	job.InputPrices = prices
}

// Process calculates tax-included prices and prints the results.
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob with the given tax rate.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
