package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   identityCenterPrincipalProperty := &IdentityCenterPrincipalProperty{
//   	GroupId: jsii.String("groupId"),
//   	UserId: jsii.String("userId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-identitycenterprincipal.html
//
type CfnEntitlementPropsMixin_IdentityCenterPrincipalProperty struct {
	// The ID of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-identitycenterprincipal.html#cfn-accountaccess-entitlement-identitycenterprincipal-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The ID of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-identitycenterprincipal.html#cfn-accountaccess-entitlement-identitycenterprincipal-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

