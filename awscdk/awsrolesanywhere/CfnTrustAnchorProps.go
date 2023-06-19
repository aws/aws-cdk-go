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
//   cfnTrustAnchorProps := &CfnTrustAnchorProps{
//   	Name: jsii.String("name"),
//   	Source: &SourceProperty{
//   		SourceData: &SourceDataProperty{
//   			AcmPcaArn: jsii.String("acmPcaArn"),
//   			X509CertificateData: jsii.String("x509CertificateData"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//   	},
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// The tags to attach to the trust anchor.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

