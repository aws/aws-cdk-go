package awseventsv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSourceProps := &CfnEventSourceProps{
//   	Configuration: &EventSourceConfigurationProperty{
//   		AwsServiceEventsConfiguration: &AwsServiceEventsConfigurationProperty{
//   			AwsService: jsii.String("awsService"),
//
//   			// the properties below are optional
//   			OnFailureConfiguration: &OnFailureConfigurationProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			Pattern: jsii.String("pattern"),
//   		},
//   		PartnerEventsConfiguration: &PartnerEventsConfigurationProperty{
//   			PartnerEventSourceArn: jsii.String("partnerEventSourceArn"),
//
//   			// the properties below are optional
//   			OnFailureConfiguration: &OnFailureConfigurationProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			PartnerBusKmsKeyIdentifier: jsii.String("partnerBusKmsKeyIdentifier"),
//   			Pattern: jsii.String("pattern"),
//   		},
//   	},
//   	EventBusArn: jsii.String("eventBusArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html
//
type CfnEventSourceProps struct {
	// Event source configuration.
	//
	// Exactly one of the two variants must be set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html#cfn-eventsv2-eventsource-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The ARN of the custom event bus the event source forwards onto.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html#cfn-eventsv2-eventsource-eventbusarn
	//
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The name of the event source.
	//
	// The first character must be alphanumeric; the remaining characters may also include '.', '-', and '_'. Names cannot begin with the reserved aws. prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html#cfn-eventsv2-eventsource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the event source.
	//
	// Control characters and Unicode line separators are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html#cfn-eventsv2-eventsource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags assigned to the event source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-eventsource.html#cfn-eventsv2-eventsource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

