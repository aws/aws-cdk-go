package interfacesawscloudformation


// A reference to a GuardHook resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardHookReference := &GuardHookReference{
//   	HookArn: jsii.String("hookArn"),
//   }
//
type GuardHookReference struct {
	// The HookArn of the GuardHook resource.
	HookArn *string `field:"required" json:"hookArn" yaml:"hookArn"`
}

