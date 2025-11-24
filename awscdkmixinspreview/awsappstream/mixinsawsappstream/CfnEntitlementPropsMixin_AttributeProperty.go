package mixinsawsappstream


// An attribute that belongs to an entitlement.
//
// Application entitlements work by matching a supported SAML 2.0 attribute name to a value when a user identity federates to an AppStream 2.0 SAML application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeProperty := &AttributeProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-entitlement-attribute.html
//
type CfnEntitlementPropsMixin_AttributeProperty struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-entitlement-attribute.html#cfn-appstream-entitlement-attribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A value that is matched to a supported SAML attribute name when a user identity federates to an AppStream 2.0 SAML application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-entitlement-attribute.html#cfn-appstream-entitlement-attribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

