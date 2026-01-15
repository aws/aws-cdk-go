package awsmpa


// Contains details for the resource that provides identities to the identity source.
//
// For example, an IAM Identity Center instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identitySourceParametersProperty := &IdentitySourceParametersProperty{
//   	IamIdentityCenter: &IamIdentityCenterProperty{
//   		InstanceArn: jsii.String("instanceArn"),
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-identitysourceparameters.html
//
type CfnIdentitySource_IdentitySourceParametersProperty struct {
	// SSOlong credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-identitysourceparameters.html#cfn-mpa-identitysource-identitysourceparameters-iamidentitycenter
	//
	IamIdentityCenter interface{} `field:"required" json:"iamIdentityCenter" yaml:"iamIdentityCenter"`
}

