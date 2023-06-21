package awsquicksight


// The total options for a pivot table visual.
//
// Example:
//
//
type CfnDashboard_PivotTableTotalOptionsProperty struct {
	// The column subtotal options.
	ColumnSubtotalOptions interface{} `field:"optional" json:"columnSubtotalOptions" yaml:"columnSubtotalOptions"`
	// The column total options.
	ColumnTotalOptions interface{} `field:"optional" json:"columnTotalOptions" yaml:"columnTotalOptions"`
	// The row subtotal options.
	RowSubtotalOptions interface{} `field:"optional" json:"rowSubtotalOptions" yaml:"rowSubtotalOptions"`
	// The row total options.
	RowTotalOptions interface{} `field:"optional" json:"rowTotalOptions" yaml:"rowTotalOptions"`
}

