package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTablePaginatedReportOptionsProperty := &PivotTablePaginatedReportOptionsProperty{
//   	OverflowColumnHeaderVisibility: jsii.String("overflowColumnHeaderVisibility"),
//   	VerticalOverflowVisibility: jsii.String("verticalOverflowVisibility"),
//   }
//
type CfnDashboard_PivotTablePaginatedReportOptionsProperty struct {
	// `CfnDashboard.PivotTablePaginatedReportOptionsProperty.OverflowColumnHeaderVisibility`.
	OverflowColumnHeaderVisibility *string `field:"optional" json:"overflowColumnHeaderVisibility" yaml:"overflowColumnHeaderVisibility"`
	// `CfnDashboard.PivotTablePaginatedReportOptionsProperty.VerticalOverflowVisibility`.
	VerticalOverflowVisibility *string `field:"optional" json:"verticalOverflowVisibility" yaml:"verticalOverflowVisibility"`
}

