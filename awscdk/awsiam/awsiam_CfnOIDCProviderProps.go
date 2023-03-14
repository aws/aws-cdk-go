package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnOIDCProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOIDCProviderProps := &cfnOIDCProviderProps{
//   	thumbprintList: []*string{
//   		jsii.String("thumbprintList"),
//   	},
//
//   	// the properties below are optional
//   	clientIdList: []*string{
//   		jsii.String("clientIdList"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	url: jsii.String("url"),
//   }
//
type CfnOIDCProviderProps struct {
	// A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object.
	//
	// For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
	ThumbprintList *[]*string `field:"required" json:"thumbprintList" yaml:"thumbprintList"`
	// A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object.
	//
	// For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
	ClientIdList *[]*string `field:"optional" json:"clientIdList" yaml:"clientIdList"`
	// A list of tags that are attached to the specified IAM OIDC provider.
	//
	// The returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The URL that the IAM OIDC provider resource object is associated with.
	//
	// For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
	Url *string `field:"optional" json:"url" yaml:"url"`
}

