package previewawsmpamixins


// Contains details for the resource that provides identities to the identity source.
//
// For example, an IAM Identity Center instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   identitySourceParametersProperty := &IdentitySourceParametersProperty{
//   	IamIdentityCenter: &IamIdentityCenterProperty{
//   		ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   		InstanceArn: jsii.String("instanceArn"),
//   		Region: jsii.String("region"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-identitysourceparameters.html
//
type CfnIdentitySourcePropsMixin_IdentitySourceParametersProperty struct {
	// AWS IAM Identity Center credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-identitysourceparameters.html#cfn-mpa-identitysource-identitysourceparameters-iamidentitycenter
	//
	IamIdentityCenter interface{} `field:"optional" json:"iamIdentityCenter" yaml:"iamIdentityCenter"`
}

