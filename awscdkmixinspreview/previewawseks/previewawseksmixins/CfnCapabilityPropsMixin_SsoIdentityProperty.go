package previewawseksmixins


// An IAM Identity CenterIAM;
//
// Identity Center identity (user or group) that can be assigned permissions in a capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ssoIdentityProperty := &SsoIdentityProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html
//
type CfnCapabilityPropsMixin_SsoIdentityProperty struct {
	// The unique identifier of the IAM Identity CenterIAM;
	//
	// Identity Center user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html#cfn-eks-capability-ssoidentity-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of identity.
	//
	// Valid values are `SSO_USER` or `SSO_GROUP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html#cfn-eks-capability-ssoidentity-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

