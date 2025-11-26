package previewawsimagebuildermixins


// Defines the rules by which an image pipeline is automatically disabled when it fails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoDisablePolicyProperty := &AutoDisablePolicyProperty{
//   	FailureCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-autodisablepolicy.html
//
type CfnImagePipelinePropsMixin_AutoDisablePolicyProperty struct {
	// The number of consecutive scheduled image pipeline executions that must fail before Image Builder automatically disables the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagepipeline-autodisablepolicy.html#cfn-imagebuilder-imagepipeline-autodisablepolicy-failurecount
	//
	FailureCount *float64 `field:"optional" json:"failureCount" yaml:"failureCount"`
}

