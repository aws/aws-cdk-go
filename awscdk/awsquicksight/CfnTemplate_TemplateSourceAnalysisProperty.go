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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceanalysis.html
//
type CfnTemplate_TemplateSourceAnalysisProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceanalysis.html#cfn-quicksight-template-templatesourceanalysis-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// A structure containing information about the dataset references used as placeholders in the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourceanalysis.html#cfn-quicksight-template-templatesourceanalysis-datasetreferences
	//
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

