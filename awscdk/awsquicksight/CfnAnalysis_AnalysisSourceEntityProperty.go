package awsquicksight


// The source entity of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceEntityProperty := &AnalysisSourceEntityProperty{
//   	SourceTemplate: &AnalysisSourceTemplateProperty{
//   		Arn: jsii.String("arn"),
//   		DataSetReferences: []interface{}{
//   			&DataSetReferenceProperty{
//   				DataSetArn: jsii.String("dataSetArn"),
//   				DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_AnalysisSourceEntityProperty struct {
	// The source template for the source entity of the analysis.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

