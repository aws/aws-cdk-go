package awsssmincidents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResponsePlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResponsePlanProps := &cfnResponsePlanProps{
//   	incidentTemplate: &incidentTemplateProperty{
//   		impact: jsii.Number(123),
//   		title: jsii.String("title"),
//
//   		// the properties below are optional
//   		dedupeString: jsii.String("dedupeString"),
//   		incidentTags: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		notificationTargets: []interface{}{
//   			&notificationTargetItemProperty{
//   				snsTopicArn: jsii.String("snsTopicArn"),
//   			},
//   		},
//   		summary: jsii.String("summary"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	actions: []interface{}{
//   		&actionProperty{
//   			ssmAutomation: &ssmAutomationProperty{
//   				documentName: jsii.String("documentName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				documentVersion: jsii.String("documentVersion"),
//   				dynamicParameters: []interface{}{
//   					&dynamicSsmParameterProperty{
//   						key: jsii.String("key"),
//   						value: &dynamicSsmParameterValueProperty{
//   							variable: jsii.String("variable"),
//   						},
//   					},
//   				},
//   				parameters: []interface{}{
//   					&ssmParameterProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				targetAccount: jsii.String("targetAccount"),
//   			},
//   		},
//   	},
//   	chatChannel: &chatChannelProperty{
//   		chatbotSns: []*string{
//   			jsii.String("chatbotSns"),
//   		},
//   	},
//   	displayName: jsii.String("displayName"),
//   	engagements: []*string{
//   		jsii.String("engagements"),
//   	},
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResponsePlanProps struct {
	// Details used to create an incident when using this response plan.
	IncidentTemplate interface{} `field:"required" json:"incidentTemplate" yaml:"incidentTemplate"`
	// The name of the response plan.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The actions that the response plan starts at the beginning of an incident.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The AWS Chatbot chat channel used for collaboration during an incident.
	ChatChannel interface{} `field:"optional" json:"chatChannel" yaml:"chatChannel"`
	// The human readable name of the response plan.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The contacts and escalation plans that the response plan engages during an incident.
	Engagements *[]*string `field:"optional" json:"engagements" yaml:"engagements"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

