package awsquicksight


// The field well configuration of a line chart.
//
// Example:
//
//
type CfnTemplate_LineChartAggregatedFieldWellsProperty struct {
	// The category field wells of a line chart.
	//
	// Values are grouped by category fields.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The color field wells of a line chart.
	//
	// Values are grouped by category fields.
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The small multiples field well of a line chart.
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a line chart.
	//
	// Values are aggregated based on categories.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

