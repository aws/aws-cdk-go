package awsquicksight


// The contribution analysis visual display for a line, pie, or bar chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contributionAnalysisDefaultProperty := &ContributionAnalysisDefaultProperty{
//   	ContributorDimensions: []interface{}{
//   		&ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	MeasureFieldId: jsii.String("measureFieldId"),
//   }
//
type CfnDashboard_ContributionAnalysisDefaultProperty struct {
	// The dimensions columns that are used in the contribution analysis, usually a list of `ColumnIdentifiers` .
	ContributorDimensions interface{} `field:"required" json:"contributorDimensions" yaml:"contributorDimensions"`
	// The measure field that is used in the contribution analysis.
	MeasureFieldId *string `field:"required" json:"measureFieldId" yaml:"measureFieldId"`
}

