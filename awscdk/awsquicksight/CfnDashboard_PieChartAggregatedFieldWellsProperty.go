package awsquicksight


// The field well configuration of a pie chart.
//
// Example:
//
//
type CfnDashboard_PieChartAggregatedFieldWellsProperty struct {
	// The category (group/color) field wells of a pie chart.
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The small multiples field well of a pie chart.
	SmallMultiples interface{} `field:"optional" json:"smallMultiples" yaml:"smallMultiples"`
	// The value field wells of a pie chart.
	//
	// Values are aggregated based on categories.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

