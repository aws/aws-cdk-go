package interfacesawsec2


// A reference to a SecurityGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityGroupReference := &SecurityGroupReference{
//   	SecurityGroupId: jsii.String("securityGroupId"),
//   }
//
type SecurityGroupReference struct {
	// The Id of the SecurityGroup resource.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
}

