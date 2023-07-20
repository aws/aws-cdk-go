package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var permissionsPolicyDocument interface{}
//
//   cfnDomainProps := &CfnDomainProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	PermissionsPolicyDocument: permissionsPolicyDocument,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html
//
type CfnDomainProps struct {
	// A string that specifies the name of the requested domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-domain.html#cfn-codeartifact-domain-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
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

