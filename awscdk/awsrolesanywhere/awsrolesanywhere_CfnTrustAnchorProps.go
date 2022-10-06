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
	// `AWS::RolesAnywhere::TrustAnchor.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `AWS::RolesAnywhere::TrustAnchor.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::RolesAnywhere::TrustAnchor.Source`.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// `AWS::RolesAnywhere::TrustAnchor.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

