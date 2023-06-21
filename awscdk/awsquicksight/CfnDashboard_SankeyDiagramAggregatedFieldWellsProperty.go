package awsquicksight


// The field well configuration of a sankey diagram.
//
// Example:
//
//
type CfnDashboard_SankeyDiagramAggregatedFieldWellsProperty struct {
	// The destination field wells of a sankey diagram.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The source field wells of a sankey diagram.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The weight field wells of a sankey diagram.
	Weight interface{} `field:"optional" json:"weight" yaml:"weight"`
}

