package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitlementProperty := &EntitlementProperty{
//   	PrincipalRole: &PrincipalRoleEntitlementProperty{
//   		Principal: &PrincipalProperty{
//   			IdentityCenter: &IdentityCenterPrincipalProperty{
//   				GroupId: jsii.String("groupId"),
//   				UserId: jsii.String("userId"),
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		Account: jsii.String("account"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-entitlement.html
//
type CfnEntitlement_EntitlementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-entitlement.html#cfn-accountaccess-entitlement-entitlement-principalrole
	//
	PrincipalRole interface{} `field:"required" json:"principalRole" yaml:"principalRole"`
}

