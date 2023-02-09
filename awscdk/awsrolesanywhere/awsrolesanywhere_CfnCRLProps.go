package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCRL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCRLProps := &cfnCRLProps{
//   	crlData: jsii.String("crlData"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trustAnchorArn: jsii.String("trustAnchorArn"),
//   }
//
type CfnCRLProps struct {
	// The revocation record for a certificate, following the x509 v3 standard.
	CrlData *string `field:"required" json:"crlData" yaml:"crlData"`
	// The name of the certificate revocation list (CRL).
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether the certificate revocation list (CRL) is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of tags to attach to the CRL.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

