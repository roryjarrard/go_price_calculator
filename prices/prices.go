package prices

import (
	"fmt"

	"github.com/roryjarrard/go-price-calculator/conversion"
	"github.com/roryjarrard/go-price-calculator/filemanager"
)

// TaxIncludedPriceJob represents a job to calculate tax-included prices.
type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// LoadData loads prices from a file and converts them to float64.
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error loading lines from file:", err)
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

	job.TaxIncludedPrices = result

	err := filemanager.WriteJSON(fmt.Sprintf("tax_included_prices_%.2f.json", job.TaxRate), result)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Printf("Processed tax-included prices with tax rate %.2f\n", job.TaxRate)
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob with the given tax rate.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
