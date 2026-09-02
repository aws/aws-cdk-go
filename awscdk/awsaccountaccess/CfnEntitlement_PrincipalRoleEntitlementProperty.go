package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalRoleEntitlementProperty := &PrincipalRoleEntitlementProperty{
//   	Principal: &PrincipalProperty{
//   		IdentityCenter: &IdentityCenterPrincipalProperty{
//   			GroupId: jsii.String("groupId"),
//   			UserId: jsii.String("userId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Account: jsii.String("account"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html
//
type CfnEntitlement_PrincipalRoleEntitlementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-principal
	//
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
	// The ARN of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principalroleentitlement.html#cfn-accountaccess-entitlement-principalroleentitlement-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
}

