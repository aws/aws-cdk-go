package awsaccountaccess


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalProperty := &PrincipalProperty{
//   	IdentityCenter: &IdentityCenterPrincipalProperty{
//   		GroupId: jsii.String("groupId"),
//   		UserId: jsii.String("userId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principal.html
//
type CfnEntitlement_PrincipalProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accountaccess-entitlement-principal.html#cfn-accountaccess-entitlement-principal-identitycenter
	//
	IdentityCenter interface{} `field:"required" json:"identityCenter" yaml:"identityCenter"`
}

