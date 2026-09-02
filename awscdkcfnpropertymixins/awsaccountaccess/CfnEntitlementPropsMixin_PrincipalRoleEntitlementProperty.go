package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   principalRoleEntitlementProperty := &PrincipalRoleEntitlementProperty{
//   	Account: jsii.String("account"),
//   	Principal: &PrincipalProperty{
//   		IdentityCenter: &IdentityCenterPrincipalProperty{
//   			GroupId: jsii.String("groupId"),
//   			UserId: jsii.String("userId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html
//
type CfnEntitlementPropsMixin_PrincipalRoleEntitlementProperty struct {
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-principal
	//
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
	// The ARN of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

