package awsquicksight


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
type CfnAnalysis_ContributionAnalysisDefaultProperty struct {
	// `CfnAnalysis.ContributionAnalysisDefaultProperty.ContributorDimensions`.
	ContributorDimensions interface{} `field:"required" json:"contributorDimensions" yaml:"contributorDimensions"`
	// `CfnAnalysis.ContributionAnalysisDefaultProperty.MeasureFieldId`.
	MeasureFieldId *string `field:"required" json:"measureFieldId" yaml:"measureFieldId"`
}

