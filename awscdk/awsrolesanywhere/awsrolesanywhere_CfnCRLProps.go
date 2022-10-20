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
	// `AWS::RolesAnywhere::CRL.CrlData`.
	CrlData *string `field:"optional" json:"crlData" yaml:"crlData"`
	// `AWS::RolesAnywhere::CRL.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `AWS::RolesAnywhere::CRL.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::RolesAnywhere::CRL.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::RolesAnywhere::CRL.TrustAnchorArn`.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

