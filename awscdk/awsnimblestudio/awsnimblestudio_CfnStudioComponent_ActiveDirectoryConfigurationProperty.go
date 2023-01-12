package awsnimblestudio


// The configuration for a Microsoft Active Directory (Microsoft AD) studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryConfigurationProperty := &activeDirectoryConfigurationProperty{
//   	computerAttributes: []interface{}{
//   		&activeDirectoryComputerAttributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	directoryId: jsii.String("directoryId"),
//   	organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   }
//
type CfnStudioComponent_ActiveDirectoryConfigurationProperty struct {
	// A collection of custom attributes for an Active Directory computer.
	ComputerAttributes interface{} `field:"optional" json:"computerAttributes" yaml:"computerAttributes"`
	// The directory ID of the Directory Service for Microsoft Active Directory to access using this studio component.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The distinguished name (DN) and organizational unit (OU) of an Active Directory computer.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

