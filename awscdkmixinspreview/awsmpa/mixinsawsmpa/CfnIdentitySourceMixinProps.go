package mixinsawsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdentitySourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdentitySourceMixinProps := &CfnIdentitySourceMixinProps{
//   	IdentitySourceParameters: &IdentitySourceParametersProperty{
//   		IamIdentityCenter: &IamIdentityCenterProperty{
//   			ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   			InstanceArn: jsii.String("instanceArn"),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html
//
type CfnIdentitySourceMixinProps struct {
	// A `IdentitySourceParameters` object.
	//
	// Contains details for the resource that provides identities to the identity source. For example, an IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html#cfn-mpa-identitysource-identitysourceparameters
	//
	IdentitySourceParameters interface{} `field:"optional" json:"identitySourceParameters" yaml:"identitySourceParameters"`
	// Tags that you have added to the specified resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html#cfn-mpa-identitysource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

