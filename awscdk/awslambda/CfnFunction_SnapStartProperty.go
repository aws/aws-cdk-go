package awslambda


// The function's SnapStart setting.
//
// When set to PublishedVersions, Lambda creates a snapshot of the execution environment when you publish a function version.
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
	// Applying SnapStart setting on function resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstart.html#cfn-lambda-function-snapstart-applyon
	//
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}

