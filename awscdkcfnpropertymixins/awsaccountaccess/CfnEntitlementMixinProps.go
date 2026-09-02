package awsaccountaccess


// Properties for CfnEntitlementPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEntitlementMixinProps := &CfnEntitlementMixinProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	Entitlement: &EntitlementProperty{
//   		PrincipalRole: &PrincipalRoleEntitlementProperty{
//   			Account: jsii.String("account"),
//   			Principal: &PrincipalProperty{
//   				IdentityCenter: &IdentityCenterPrincipalProperty{
//   					GroupId: jsii.String("groupId"),
//   					UserId: jsii.String("userId"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html
//
type CfnEntitlementMixinProps struct {
	// The ARN of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html#cfn-accountaccess-entitlement-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html#cfn-accountaccess-entitlement-entitlement
	//
	Entitlement interface{} `field:"optional" json:"entitlement" yaml:"entitlement"`
}

