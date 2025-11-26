package previewawscodeartifactmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var permissionsPolicyDocument interface{}
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	PermissionsPolicyDocument: permissionsPolicyDocument,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html
//
type CfnDomainMixinProps struct {
	// A string that specifies the name of the requested domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html#cfn-codeartifact-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The key used to encrypt the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html#cfn-codeartifact-domain-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The document that defines the resource policy that is set on a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html#cfn-codeartifact-domain-permissionspolicydocument
	//
	PermissionsPolicyDocument interface{} `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// A list of tags to be applied to the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html#cfn-codeartifact-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

