package awscloudformation


// A reference to a HookVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hookVersionReference := &HookVersionReference{
//   	HookVersionArn: jsii.String("hookVersionArn"),
//   }
//
type HookVersionReference struct {
	// The Arn of the HookVersion resource.
	HookVersionArn *string `field:"required" json:"hookVersionArn" yaml:"hookVersionArn"`
}

