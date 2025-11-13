package interfacesawsconnect


// A reference to a SecurityKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityKeyReference := &SecurityKeyReference{
//   	AssociationId: jsii.String("associationId"),
//   	InstanceId: jsii.String("instanceId"),
//   }
//
type SecurityKeyReference struct {
	// The AssociationId of the SecurityKey resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
	// The InstanceId of the SecurityKey resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

