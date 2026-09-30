package interfacesawsnetworkflowmonitor


// A reference to a Scope resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scopeReference := &ScopeReference{
//   	ScopeArn: jsii.String("scopeArn"),
//   }
//
type ScopeReference struct {
	// The ScopeArn of the Scope resource.
	ScopeArn *string `field:"required" json:"scopeArn" yaml:"scopeArn"`
}

