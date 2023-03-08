package awsquicksight


// The source template of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceTemplateProperty := &analysisSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   	dataSetReferences: []interface{}{
//   		&dataSetReferenceProperty{
//   			dataSetArn: jsii.String("dataSetArn"),
//   			dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnAnalysis_AnalysisSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the source template of an analysis.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The dataset references of the source template of an analysis.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

