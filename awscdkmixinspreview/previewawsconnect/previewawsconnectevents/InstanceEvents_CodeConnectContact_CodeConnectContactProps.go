package previewawsconnectevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.connect@CodeConnectContact event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeConnectContactProps := &CodeConnectContactProps{
//   	AgentInfo: &AgentInfo{
//   		AgentArn: []*string{
//   			jsii.String("agentArn"),
//   		},
//   	},
//   	Channel: []*string{
//   		jsii.String("channel"),
//   	},
//   	ContactId: []*string{
//   		jsii.String("contactId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	InitialContactId: []*string{
//   		jsii.String("initialContactId"),
//   	},
//   	InitiationMethod: []*string{
//   		jsii.String("initiationMethod"),
//   	},
//   	InstanceArn: []*string{
//   		jsii.String("instanceArn"),
//   	},
//   	PreviousContactId: []*string{
//   		jsii.String("previousContactId"),
//   	},
//   	QueueInfo: &QueueInfo{
//   		QueueArn: []*string{
//   			jsii.String("queueArn"),
//   		},
//   		QueueType: []*string{
//   			jsii.String("queueType"),
//   		},
//   	},
//   }
//
// Experimental.
type InstanceEvents_CodeConnectContact_CodeConnectContactProps struct {
	// agentInfo property.
	//
	// Specify an array of string values to match this event if the actual value of agentInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentInfo *InstanceEvents_CodeConnectContact_AgentInfo `field:"optional" json:"agentInfo" yaml:"agentInfo"`
	// channel property.
	//
	// Specify an array of string values to match this event if the actual value of channel is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Channel *[]*string `field:"optional" json:"channel" yaml:"channel"`
	// contactId property.
	//
	// Specify an array of string values to match this event if the actual value of contactId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContactId *[]*string `field:"optional" json:"contactId" yaml:"contactId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// initialContactId property.
	//
	// Specify an array of string values to match this event if the actual value of initialContactId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitialContactId *[]*string `field:"optional" json:"initialContactId" yaml:"initialContactId"`
	// initiationMethod property.
	//
	// Specify an array of string values to match this event if the actual value of initiationMethod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InitiationMethod *[]*string `field:"optional" json:"initiationMethod" yaml:"initiationMethod"`
	// instanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of instanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceArn *[]*string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// previousContactId property.
	//
	// Specify an array of string values to match this event if the actual value of previousContactId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousContactId *[]*string `field:"optional" json:"previousContactId" yaml:"previousContactId"`
	// queueInfo property.
	//
	// Specify an array of string values to match this event if the actual value of queueInfo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueueInfo *InstanceEvents_CodeConnectContact_QueueInfo `field:"optional" json:"queueInfo" yaml:"queueInfo"`
}

