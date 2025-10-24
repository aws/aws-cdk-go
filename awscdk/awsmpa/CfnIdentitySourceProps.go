package awsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIdentitySource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentitySourceProps := &CfnIdentitySourceProps{
//   	IdentitySourceParameters: &IdentitySourceParametersProperty{
//   		IamIdentityCenter: &IamIdentityCenterProperty{
//   			InstanceArn: jsii.String("instanceArn"),
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   		},
//   	},
//
//   	// the properties below are optional
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
type CfnIdentitySourceProps struct {
	// A `IdentitySourceParameters` object.
	//
	// Contains details for the resource that provides identities to the identity source. For example, an IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html#cfn-mpa-identitysource-identitysourceparameters
	//
	IdentitySourceParameters interface{} `field:"required" json:"identitySourceParameters" yaml:"identitySourceParameters"`
	// Tags that you have added to the specified resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mpa-identitysource.html#cfn-mpa-identitysource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

