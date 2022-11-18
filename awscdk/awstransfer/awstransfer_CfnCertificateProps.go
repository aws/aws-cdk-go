package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificate: jsii.String("certificate"),
//   	usage: jsii.String("usage"),
//
//   	// the properties below are optional
//   	activeDate: jsii.String("activeDate"),
//   	certificateChain: jsii.String("certificateChain"),
//   	description: jsii.String("description"),
//   	inactiveDate: jsii.String("inactiveDate"),
//   	privateKey: jsii.String("privateKey"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCertificateProps struct {
	// `AWS::Transfer::Certificate.Certificate`.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// `AWS::Transfer::Certificate.Usage`.
	Usage *string `field:"required" json:"usage" yaml:"usage"`
	// `AWS::Transfer::Certificate.ActiveDate`.
	ActiveDate *string `field:"optional" json:"activeDate" yaml:"activeDate"`
	// `AWS::Transfer::Certificate.CertificateChain`.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// `AWS::Transfer::Certificate.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Transfer::Certificate.InactiveDate`.
	InactiveDate *string `field:"optional" json:"inactiveDate" yaml:"inactiveDate"`
	// `AWS::Transfer::Certificate.PrivateKey`.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// `AWS::Transfer::Certificate.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

