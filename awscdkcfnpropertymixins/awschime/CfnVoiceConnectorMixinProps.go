package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVoiceConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnVoiceConnectorMixinProps := &CfnVoiceConnectorMixinProps{
//   	AwsRegion: jsii.String("awsRegion"),
//   	Name: jsii.String("name"),
//   	NetworkType: jsii.String("networkType"),
//   	RequireEncryption: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html
//
type CfnVoiceConnectorMixinProps struct {
	// The AWS Region in which the Voice Connector is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html#cfn-chime-voiceconnector-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The name of the Voice Connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html#cfn-chime-voiceconnector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of network for the Voice Connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html#cfn-chime-voiceconnector-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Enables or disables encryption for the Voice Connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html#cfn-chime-voiceconnector-requireencryption
	//
	RequireEncryption interface{} `field:"optional" json:"requireEncryption" yaml:"requireEncryption"`
	// The tags assigned to the Voice Connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-voiceconnector.html#cfn-chime-voiceconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

