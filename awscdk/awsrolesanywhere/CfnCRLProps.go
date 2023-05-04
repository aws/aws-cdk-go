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
	// x509 v3 Certificate Revocation List to revoke auth for corresponding certificates presented in CreateSession operations.
	CrlData *string `field:"required" json:"crlData" yaml:"crlData"`
	// The customer specified name of the resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The enabled status of the resource.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of Tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

