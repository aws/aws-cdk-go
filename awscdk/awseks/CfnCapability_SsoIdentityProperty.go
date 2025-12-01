package awseks


// An IAM Identity Center identity (user or group) that can be assigned permissions in a capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssoIdentityProperty := &SsoIdentityProperty{
//   	Id: jsii.String("id"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html
//
type CfnCapability_SsoIdentityProperty struct {
	// The unique identifier of the IAM Identity Center user or group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html#cfn-eks-capability-ssoidentity-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The type of identity.
	//
	// Valid values are SSO_USER or SSO_GROUP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-ssoidentity.html#cfn-eks-capability-ssoidentity-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

