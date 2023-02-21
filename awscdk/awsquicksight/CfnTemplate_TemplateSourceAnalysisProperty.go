package awsquicksight


// The source analysis of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceAnalysisProperty := &TemplateSourceAnalysisProperty{
//   	Arn: jsii.String("arn"),
//   	DataSetReferences: []interface{}{
//   		&DataSetReferenceProperty{
//   			DataSetArn: jsii.String("dataSetArn"),
//   			DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnTemplate_TemplateSourceAnalysisProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// A structure containing information about the dataset references used as placeholders in the template.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

