package prices

import (
	"fmt"

	"github.com/roryjarrard/go-price-calculator/conversion"
	"github.com/roryjarrard/go-price-calculator/iomanager"
)

// TaxIncludedPriceJob represents a job to calculate tax-included prices.
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// LoadData loads prices from a file and converts them to float64.
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
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

	job.IOManager.WriteResult(job)
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob with the given tax rate.
func NewTaxIncludedPriceJob(m iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   m,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
