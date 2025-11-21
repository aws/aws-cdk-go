package awsivschat

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRoom`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRoomProps := &CfnRoomProps{
//   	LoggingConfigurationIdentifiers: []interface{}{
//   		jsii.String("loggingConfigurationIdentifiers"),
//   	},
//   	MaximumMessageLength: jsii.Number(123),
//   	MaximumMessageRatePerSecond: jsii.Number(123),
//   	MessageReviewHandler: &MessageReviewHandlerProperty{
//   		FallbackResult: jsii.String("fallbackResult"),
//   		Uri: jsii.String("uri"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html
//
type CfnRoomProps struct {
	// List of logging-configuration identifiers attached to the room.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-loggingconfigurationidentifiers
	//
	LoggingConfigurationIdentifiers *[]interface{} `field:"optional" json:"loggingConfigurationIdentifiers" yaml:"loggingConfigurationIdentifiers"`
	// Maximum number of characters in a single message.
	//
	// Messages are expected to be UTF-8 encoded and this limit applies specifically to rune/code-point count, not number of bytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-maximummessagelength
	//
	// Default: - 500.
	//
	MaximumMessageLength *float64 `field:"optional" json:"maximumMessageLength" yaml:"maximumMessageLength"`
	// Maximum number of messages per second that can be sent to the room (by all clients).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-maximummessageratepersecond
	//
	// Default: - 10.
	//
	MaximumMessageRatePerSecond *float64 `field:"optional" json:"maximumMessageRatePerSecond" yaml:"maximumMessageRatePerSecond"`
	// Configuration information for optional review of messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-messagereviewhandler
	//
	MessageReviewHandler interface{} `field:"optional" json:"messageReviewHandler" yaml:"messageReviewHandler"`
	// Room name.
	//
	// The value does not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivschat-room-tag.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-room.html#cfn-ivschat-room-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

