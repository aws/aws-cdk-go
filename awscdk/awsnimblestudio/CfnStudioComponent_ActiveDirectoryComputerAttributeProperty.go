package awsnimblestudio


// An LDAP attribute of an Active Directory computer account, in the form of a name:value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activeDirectoryComputerAttributeProperty := &ActiveDirectoryComputerAttributeProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html
//
type CfnStudioComponent_ActiveDirectoryComputerAttributeProperty struct {
	// The name for the LDAP attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html#cfn-nimblestudio-studiocomponent-activedirectorycomputerattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value for the LDAP attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-activedirectorycomputerattribute.html#cfn-nimblestudio-studiocomponent-activedirectorycomputerattribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

