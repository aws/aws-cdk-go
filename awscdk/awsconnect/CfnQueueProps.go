package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQueue`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQueueProps := &CfnQueueProps{
//   	HoursOfOperationArn: jsii.String("hoursOfOperationArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MaxContacts: jsii.Number(123),
//   	OutboundCallerConfig: &OutboundCallerConfigProperty{
//   		OutboundCallerIdName: jsii.String("outboundCallerIdName"),
//   		OutboundCallerIdNumberArn: jsii.String("outboundCallerIdNumberArn"),
//   		OutboundFlowArn: jsii.String("outboundFlowArn"),
//   	},
//   	QuickConnectArns: []*string{
//   		jsii.String("quickConnectArns"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html
//
type CfnQueueProps struct {
	// The Amazon Resource Name (ARN) of the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-hoursofoperationarn
	//
	HoursOfOperationArn *string `field:"required" json:"hoursOfOperationArn" yaml:"hoursOfOperationArn"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of contacts that can be in the queue before it is considered full.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-maxcontacts
	//
	MaxContacts *float64 `field:"optional" json:"maxContacts" yaml:"maxContacts"`
	// The outbound caller ID name, number, and outbound whisper flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-outboundcallerconfig
	//
	OutboundCallerConfig interface{} `field:"optional" json:"outboundCallerConfig" yaml:"outboundCallerConfig"`
	// The Amazon Resource Names (ARN) of the of the quick connects available to agents who are working the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-quickconnectarns
	//
	QuickConnectArns *[]*string `field:"optional" json:"quickConnectArns" yaml:"quickConnectArns"`
	// The status of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

