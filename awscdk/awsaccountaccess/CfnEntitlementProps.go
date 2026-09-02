package awsaccountaccess


// Properties for defining a `CfnEntitlement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntitlementProps := &CfnEntitlementProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	Entitlement: &EntitlementProperty{
//   		PrincipalRole: &PrincipalRoleEntitlementProperty{
//   			Principal: &PrincipalProperty{
//   				IdentityCenter: &IdentityCenterPrincipalProperty{
//   					GroupId: jsii.String("groupId"),
//   					UserId: jsii.String("userId"),
//   				},
//   			},
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			Account: jsii.String("account"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html
//
type CfnEntitlementProps struct {
	// The ARN of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html#cfn-accountaccess-entitlement-applicationarn
	//
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-entitlement.html#cfn-accountaccess-entitlement-entitlement
	//
	Entitlement interface{} `field:"required" json:"entitlement" yaml:"entitlement"`
}

