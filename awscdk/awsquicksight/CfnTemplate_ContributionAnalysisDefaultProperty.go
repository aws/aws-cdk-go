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
type CfnTemplate_ContributionAnalysisDefaultProperty struct {
	// `CfnTemplate.ContributionAnalysisDefaultProperty.ContributorDimensions`.
	ContributorDimensions interface{} `field:"required" json:"contributorDimensions" yaml:"contributorDimensions"`
	// `CfnTemplate.ContributionAnalysisDefaultProperty.MeasureFieldId`.
	MeasureFieldId *string `field:"required" json:"measureFieldId" yaml:"measureFieldId"`
}

