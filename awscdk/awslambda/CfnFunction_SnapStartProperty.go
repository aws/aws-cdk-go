package awslambda


// The function's [AWS Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapStartProperty := &SnapStartProperty{
//   	ApplyOn: jsii.String("applyOn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstart.html
//
type CfnFunction_SnapStartProperty struct {
	// Set `ApplyOn` to `PublishedVersions` to create a snapshot of the initialized execution environment when you publish a function version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstart.html#cfn-lambda-function-snapstart-applyon
	//
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}

