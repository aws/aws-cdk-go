package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectorProps := &cfnConnectorProps{
//   	accessRole: jsii.String("accessRole"),
//   	as2Config: &as2ConfigProperty{
//   		compression: jsii.String("compression"),
//   		encryptionAlgorithm: jsii.String("encryptionAlgorithm"),
//   		localProfileId: jsii.String("localProfileId"),
//   		mdnResponse: jsii.String("mdnResponse"),
//   		mdnSigningAlgorithm: jsii.String("mdnSigningAlgorithm"),
//   		messageSubject: jsii.String("messageSubject"),
//   		partnerProfileId: jsii.String("partnerProfileId"),
//   		signingAlgorithm: jsii.String("signingAlgorithm"),
//   	},
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	loggingRole: jsii.String("loggingRole"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectorProps struct {
	// `AWS::Transfer::Connector.AccessRole`.
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// `AWS::Transfer::Connector.As2Config`.
	As2Config interface{} `field:"required" json:"as2Config" yaml:"as2Config"`
	// `AWS::Transfer::Connector.Url`.
	Url *string `field:"required" json:"url" yaml:"url"`
	// `AWS::Transfer::Connector.LoggingRole`.
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// `AWS::Transfer::Connector.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

