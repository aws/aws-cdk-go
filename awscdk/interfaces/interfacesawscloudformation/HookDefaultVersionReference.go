package interfacesawscloudformation


// A reference to a HookDefaultVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hookDefaultVersionReference := &HookDefaultVersionReference{
//   	HookDefaultVersionArn: jsii.String("hookDefaultVersionArn"),
//   }
//
type HookDefaultVersionReference struct {
	// The Arn of the HookDefaultVersion resource.
	HookDefaultVersionArn *string `field:"required" json:"hookDefaultVersionArn" yaml:"hookDefaultVersionArn"`
}

