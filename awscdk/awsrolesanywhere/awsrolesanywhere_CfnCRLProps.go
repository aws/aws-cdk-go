package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCRL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCRLProps := &CfnCRLProps{
//   	CrlData: jsii.String("crlData"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustAnchorArn: jsii.String("trustAnchorArn"),
//   }
//
type CfnCRLProps struct {
	// The x509 v3 specified certificate revocation list (CRL).
	CrlData *string `field:"required" json:"crlData" yaml:"crlData"`
	// The name of the certificate revocation list (CRL).
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether the certificate revocation list (CRL) is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of tags to attach to the certificate revocation list (CRL).
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

