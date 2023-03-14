package awsappstream


// An attribute that belongs to an entitlement.
//
// Application entitlements work by matching a supported SAML 2.0 attribute name to a value when a user identity federates to an AppStream 2.0 SAML application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeProperty := &attributeProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnEntitlement_AttributeProperty struct {
	// A supported AWS IAM SAML PrincipalTag attribute that is matched to a value when a user identity federates to an AppStream 2.0 SAML application.
	//
	// The following are supported values:
	//
	// - roles
	// - department
	// - organization
	// - groups
	// - title
	// - costCenter
	// - userType.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A value that is matched to a supported SAML attribute name when a user identity federates to an AppStream 2.0 SAML application.
	Value *string `field:"required" json:"value" yaml:"value"`
}

