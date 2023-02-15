package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrustAnchor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrustAnchorProps := &cfnTrustAnchorProps{
//   	name: jsii.String("name"),
//   	source: &sourceProperty{
//   		sourceData: &sourceDataProperty{
//   			acmPcaArn: jsii.String("acmPcaArn"),
//   			x509CertificateData: jsii.String("x509CertificateData"),
//   		},
//   		sourceType: jsii.String("sourceType"),
//   	},
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTrustAnchorProps struct {
	// The name of the trust anchor.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The trust anchor type and its related certificate data.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Indicates whether the trust anchor is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of tags to attach to the trust anchor.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

