package awsnimblestudio


// An LDAP attribute of an Active Directory computer account, in the form of a name:value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryComputerAttributeProperty := &activeDirectoryComputerAttributeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnStudioComponent_ActiveDirectoryComputerAttributeProperty struct {
	// The name for the LDAP attribute.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value for the LDAP attribute.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

