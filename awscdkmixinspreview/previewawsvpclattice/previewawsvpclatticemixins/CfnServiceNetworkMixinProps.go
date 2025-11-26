package previewawsvpclatticemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServiceNetworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceNetworkMixinProps := &CfnServiceNetworkMixinProps{
//   	AuthType: jsii.String("authType"),
//   	Name: jsii.String("name"),
//   	SharingConfig: &SharingConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetwork.html
//
type CfnServiceNetworkMixinProps struct {
	// The type of IAM policy.
	//
	// - `NONE` : The resource does not use an IAM policy. This is the default.
	// - `AWS_IAM` : The resource uses an IAM policy. When this type is used, auth is enabled and an auth policy is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetwork.html#cfn-vpclattice-servicenetwork-authtype
	//
	// Default: - "NONE".
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The name of the service network.
	//
	// The name must be unique to the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetwork.html#cfn-vpclattice-servicenetwork-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify if the service network should be enabled for sharing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetwork.html#cfn-vpclattice-servicenetwork-sharingconfig
	//
	SharingConfig interface{} `field:"optional" json:"sharingConfig" yaml:"sharingConfig"`
	// The tags for the service network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetwork.html#cfn-vpclattice-servicenetwork-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

