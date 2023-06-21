package awsquicksight


// The aggregated field wells of a bar chart.
//
// Example:
//
//
type CfnTemplate_BarChartAggregatedFieldWellsProperty struct {
	// The category (y-axis) field well of a bar chart.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color (group/color) field well of a bar chart.
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The small multiples field well of a bar chart.
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a bar chart.
	//
	// Values are aggregated by category.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

