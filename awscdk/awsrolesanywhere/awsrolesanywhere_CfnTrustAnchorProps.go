package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTrustAnchor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustAnchorProps := &cfnTrustAnchorProps{
//   	enabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	source: &sourceProperty{
//   		sourceData: &sourceDataProperty{
//   			acmPcaArn: jsii.String("acmPcaArn"),
//   			x509CertificateData: jsii.String("x509CertificateData"),
//   		},
//   		sourceType: jsii.String("sourceType"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTrustAnchorProps struct {
	// Indicates whether the trust anchor is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the trust anchor.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The trust anchor type and its related certificate data.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// A list of tags to attach to the trust anchor.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

