package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnDomainProps := &cfnDomainProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	encryptionKey: jsii.String("encryptionKey"),
//   	permissionsPolicyDocument: permissionsPolicyDocument,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainProps struct {
	// A string that specifies the name of the requested domain.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The key used to encrypt the domain.
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The document that defines the resource policy that is set on a domain.
	PermissionsPolicyDocument interface{} `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// A list of tags to be applied to the domain.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

