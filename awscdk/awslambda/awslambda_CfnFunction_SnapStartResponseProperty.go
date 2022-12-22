package awslambda


// The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapStartResponseProperty := &snapStartResponseProperty{
//   	applyOn: jsii.String("applyOn"),
//   	optimizationStatus: jsii.String("optimizationStatus"),
//   }
//
type CfnFunction_SnapStartResponseProperty struct {
	// When set to `PublishedVersions` , Lambda creates a snapshot of the execution environment when you publish a function version.
	ApplyOn *string `field:"optional" json:"applyOn" yaml:"applyOn"`
	// When you provide a [qualified Amazon Resource Name (ARN)](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using) , this response element indicates whether SnapStart is activated for the specified function version.
	OptimizationStatus *string `field:"optional" json:"optimizationStatus" yaml:"optimizationStatus"`
}

