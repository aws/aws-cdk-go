package previewawsconnectevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.connect@ContactLensRealtimeRulesMatched event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   contactLensRealtimeRulesMatchedProps := &ContactLensRealtimeRulesMatchedProps{
//   	ActionName: []*string{
//   		jsii.String("actionName"),
//   	},
//   	AgentArn: []*string{
//   		jsii.String("agentArn"),
//   	},
//   	ContactArn: []*string{
//   		jsii.String("contactArn"),
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
//   	InstanceArn: []*string{
//   		jsii.String("instanceArn"),
//   	},
//   	QueueArn: []*string{
//   		jsii.String("queueArn"),
//   	},
//   	RuleName: []*string{
//   		jsii.String("ruleName"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps struct {
	// actionName property.
	//
	// Specify an array of string values to match this event if the actual value of actionName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionName *[]*string `field:"optional" json:"actionName" yaml:"actionName"`
	// agentArn property.
	//
	// Specify an array of string values to match this event if the actual value of agentArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AgentArn *[]*string `field:"optional" json:"agentArn" yaml:"agentArn"`
	// contactArn property.
	//
	// Specify an array of string values to match this event if the actual value of contactArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContactArn *[]*string `field:"optional" json:"contactArn" yaml:"contactArn"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instanceArn property.
	//
	// Specify an array of string values to match this event if the actual value of instanceArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceArn *[]*string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// queueArn property.
	//
	// Specify an array of string values to match this event if the actual value of queueArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueueArn *[]*string `field:"optional" json:"queueArn" yaml:"queueArn"`
	// ruleName property.
	//
	// Specify an array of string values to match this event if the actual value of ruleName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RuleName *[]*string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

