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
//   cfnCRLProps := &cfnCRLProps{
//   	crlData: jsii.String("crlData"),
//   	enabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
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
	// x509 v3 Certificate Revocation List to revoke auth for corresponding certificates presented in CreateSession operations.
	CrlData *string `field:"optional" json:"crlData" yaml:"crlData"`
	// The enabled status of the resource.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The customer specified name of the resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of Tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

