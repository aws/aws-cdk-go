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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html
//
type CfnCRLProps struct {
	// x509 v3 Certificate Revocation List to revoke auth for corresponding certificates presented in CreateSession operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html#cfn-rolesanywhere-crl-crldata
	//
	CrlData *string `field:"required" json:"crlData" yaml:"crlData"`
	// The customer specified name of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html#cfn-rolesanywhere-crl-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The enabled status of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html#cfn-rolesanywhere-crl-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of Tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html#cfn-rolesanywhere-crl-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-crl.html#cfn-rolesanywhere-crl-trustanchorarn
	//
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

