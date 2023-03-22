package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnServerCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerCertificateProps := &cfnServerCertificateProps{
//   	certificateBody: jsii.String("certificateBody"),
//   	certificateChain: jsii.String("certificateChain"),
//   	path: jsii.String("path"),
//   	privateKey: jsii.String("privateKey"),
//   	serverCertificateName: jsii.String("serverCertificateName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServerCertificateProps struct {
	// The contents of the public key certificate.
	CertificateBody *string `field:"optional" json:"certificateBody" yaml:"certificateBody"`
	// The contents of the public key certificate chain.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The path for the server certificate.
	//
	// For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide* .
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/). This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! ( `\ u0021` ) through the DEL character ( `\ u007F` ), including most punctuation characters, digits, and upper and lowercased letters.
	//
	// > If you are uploading a server certificate specifically for use with Amazon CloudFront distributions, you must specify a path using the `path` parameter. The path must begin with `/cloudfront` and must include a trailing slash (for example, `/cloudfront/test/` ).
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The contents of the private key in PEM-encoded format.
	//
	// The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//
	// - Any printable ASCII character ranging from the space character ( `\ u0020` ) through the end of the ASCII character range
	// - The printable characters in the Basic Latin and Latin-1 Supplement character set (through `\ u00FF` )
	// - The special characters tab ( `\ u0009` ), line feed ( `\ u000A` ), and carriage return ( `\ u000D` ).
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The name for the server certificate.
	//
	// Do not include the path in this value. The name of the certificate cannot contain any spaces.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	ServerCertificateName *string `field:"optional" json:"serverCertificateName" yaml:"serverCertificateName"`
	// A list of tags that are attached to the server certificate.
	//
	// For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

