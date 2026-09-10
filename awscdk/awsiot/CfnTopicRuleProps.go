package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTopicRule`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html
//
type CfnTopicRuleProps struct {
	// The rule payload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-topicrulepayload
	//
	TopicRulePayload interface{} `field:"required" json:"topicRulePayload" yaml:"topicRulePayload"`
	// The name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Metadata which can be used to manage the topic rule.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-topicrule.html#cfn-iot-topicrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

