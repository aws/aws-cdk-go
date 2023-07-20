package awsnimblestudio


// The configuration for a AWS Directory Service for Microsoft Active Directory studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryConfigurationProperty := &ActiveDirectoryConfigurationProperty{
//   	ComputerAttributes: []interface{}{
//   		&ActiveDirectoryComputerAttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DirectoryId: jsii.String("directoryId"),
//   	OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectoryconfiguration.html
//
type CfnStudioComponent_ActiveDirectoryConfigurationProperty struct {
	// A collection of custom attributes for an Active Directory computer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectoryconfiguration.html#cfn-nimblestudio-studiocomponent-activedirectoryconfiguration-computerattributes
	//
	ComputerAttributes interface{} `field:"optional" json:"computerAttributes" yaml:"computerAttributes"`
	// The directory ID of the AWS Directory Service for Microsoft Active Directory to access using this studio component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectoryconfiguration.html#cfn-nimblestudio-studiocomponent-activedirectoryconfiguration-directoryid
	//
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The distinguished name (DN) and organizational unit (OU) of an Active Directory computer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectoryconfiguration.html#cfn-nimblestudio-studiocomponent-activedirectoryconfiguration-organizationalunitdistinguishedname
	//
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

