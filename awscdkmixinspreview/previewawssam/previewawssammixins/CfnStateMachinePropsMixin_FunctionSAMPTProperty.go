package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   functionSAMPTProperty := &FunctionSAMPTProperty{
//   	FunctionName: jsii.String("functionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-functionsampt.html
//
type CfnStateMachinePropsMixin_FunctionSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-statemachine-functionsampt.html#cfn-serverless-statemachine-functionsampt-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
}

