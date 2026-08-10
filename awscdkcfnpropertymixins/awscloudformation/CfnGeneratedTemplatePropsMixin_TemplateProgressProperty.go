package awscloudformation


// A summary of the progress of the template generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   templateProgressProperty := &TemplateProgressProperty{
//   	ResourcesFailed: jsii.Number(123),
//   	ResourcesPending: jsii.Number(123),
//   	ResourcesProcessing: jsii.Number(123),
//   	ResourcesSucceeded: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateprogress.html
//
type CfnGeneratedTemplatePropsMixin_TemplateProgressProperty struct {
	// The number of resources that failed the template generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateprogress.html#cfn-cloudformation-generatedtemplate-templateprogress-resourcesfailed
	//
	ResourcesFailed *float64 `field:"optional" json:"resourcesFailed" yaml:"resourcesFailed"`
	// The number of resources that are still pending the template generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateprogress.html#cfn-cloudformation-generatedtemplate-templateprogress-resourcespending
	//
	ResourcesPending *float64 `field:"optional" json:"resourcesPending" yaml:"resourcesPending"`
	// The number of resources that are in-process for the template generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateprogress.html#cfn-cloudformation-generatedtemplate-templateprogress-resourcesprocessing
	//
	ResourcesProcessing *float64 `field:"optional" json:"resourcesProcessing" yaml:"resourcesProcessing"`
	// The number of resources that succeeded the template generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateprogress.html#cfn-cloudformation-generatedtemplate-templateprogress-resourcessucceeded
	//
	ResourcesSucceeded *float64 `field:"optional" json:"resourcesSucceeded" yaml:"resourcesSucceeded"`
}

