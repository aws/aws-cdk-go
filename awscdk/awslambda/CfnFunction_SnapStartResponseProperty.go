package awslambda


// The function's SnapStart Response.
//
// When set to PublishedVersions, Lambda creates a snapshot of the execution environment when you publish a function version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapStartResponseProperty := &SnapStartResponseProperty{
//   	ApplyOn: jsii.String("applyOn"),
//   	OptimizationStatus: jsii.String("optimizationStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstartresponse.html
//
type CfnFunction_SnapStartResponseProperty struct {
	// Applying SnapStart setting on function resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstartresponse.html#cfn-lambda-function-snapstartresponse-applyon
	//
	ApplyOn *string `field:"optional" json:"applyOn" yaml:"applyOn"`
	// Indicates whether SnapStart is activated for the specified function version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-snapstartresponse.html#cfn-lambda-function-snapstartresponse-optimizationstatus
	//
	OptimizationStatus *string `field:"optional" json:"optimizationStatus" yaml:"optimizationStatus"`
}

