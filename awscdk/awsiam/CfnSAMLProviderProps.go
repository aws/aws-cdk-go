package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSAMLProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSAMLProviderProps := &CfnSAMLProviderProps{
//   	AddPrivateKey: jsii.String("addPrivateKey"),
//   	AssertionEncryptionMode: jsii.String("assertionEncryptionMode"),
//   	Name: jsii.String("name"),
//   	PrivateKeyList: []interface{}{
//   		&SAMLPrivateKeyProperty{
//   			KeyId: jsii.String("keyId"),
//   			Timestamp: jsii.String("timestamp"),
//   		},
//   	},
//   	RemovePrivateKey: jsii.String("removePrivateKey"),
//   	SamlMetadataDocument: jsii.String("samlMetadataDocument"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html
//
type CfnSAMLProviderProps struct {
	// Specifies the new private key from your external identity provider.
	//
	// The private key must be a .pem file that uses AES-GCM or AES-CBC encryption algorithm to decrypt SAML assertions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-addprivatekey
	//
	AddPrivateKey *string `field:"optional" json:"addPrivateKey" yaml:"addPrivateKey"`
	// Specifies the encryption setting for the SAML provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-assertionencryptionmode
	//
	AssertionEncryptionMode *string `field:"optional" json:"assertionEncryptionMode" yaml:"assertionEncryptionMode"`
	// The name of the provider to create.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The private key metadata for the SAML provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-privatekeylist
	//
	PrivateKeyList interface{} `field:"optional" json:"privateKeyList" yaml:"privateKeyList"`
	// The Key ID of the private key to remove.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-removeprivatekey
	//
	RemovePrivateKey *string `field:"optional" json:"removePrivateKey" yaml:"removePrivateKey"`
	// An XML document generated by an identity provider (IdP) that supports SAML 2.0. The document includes the issuer's name, expiration information, and keys that can be used to validate the SAML authentication response (assertions) that are received from the IdP. You must generate the metadata document using the identity management software that is used as your organization's IdP.
	//
	// For more information, see [About SAML 2.0-based federation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html) in the *IAM User Guide*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-samlmetadatadocument
	//
	SamlMetadataDocument *string `field:"optional" json:"samlMetadataDocument" yaml:"samlMetadataDocument"`
	// A list of tags that you want to attach to the new IAM SAML provider.
	//
	// Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	//
	// > If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

