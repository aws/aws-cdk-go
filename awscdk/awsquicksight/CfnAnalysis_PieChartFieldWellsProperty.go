package awsquicksight


// The field well configuration of a pie chart.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnAnalysis_PieChartFieldWellsProperty struct {
	// The field well configuration of a pie chart.
	PieChartAggregatedFieldWells interface{} `field:"optional" json:"pieChartAggregatedFieldWells" yaml:"pieChartAggregatedFieldWells"`
}

