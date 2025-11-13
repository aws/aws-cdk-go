package interfacesawscloudformation


// A reference to a LambdaHook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaHookReference := &LambdaHookReference{
//   	HookArn: jsii.String("hookArn"),
//   }
//
type LambdaHookReference struct {
	// The HookArn of the LambdaHook resource.
	HookArn *string `field:"required" json:"hookArn" yaml:"hookArn"`
}

