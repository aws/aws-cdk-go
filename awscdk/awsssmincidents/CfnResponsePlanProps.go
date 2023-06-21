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
//   cfnResponsePlanProps := &CfnResponsePlanProps{
//   	IncidentTemplate: &IncidentTemplateProperty{
//   		Impact: jsii.Number(123),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DedupeString: jsii.String("dedupeString"),
//   		IncidentTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NotificationTargets: []interface{}{
//   			&NotificationTargetItemProperty{
//   				SnsTopicArn: jsii.String("snsTopicArn"),
//   			},
//   		},
//   		Summary: jsii.String("summary"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			SsmAutomation: &SsmAutomationProperty{
//   				DocumentName: jsii.String("documentName"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				DocumentVersion: jsii.String("documentVersion"),
//   				DynamicParameters: []interface{}{
//   					&DynamicSsmParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: &DynamicSsmParameterValueProperty{
//   							Variable: jsii.String("variable"),
//   						},
//   					},
//   				},
//   				Parameters: []interface{}{
//   					&SsmParameterProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				TargetAccount: jsii.String("targetAccount"),
//   			},
//   		},
//   	},
//   	ChatChannel: &ChatChannelProperty{
//   		ChatbotSns: []*string{
//   			jsii.String("chatbotSns"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Engagements: []*string{
//   		jsii.String("engagements"),
//   	},
//   	Integrations: []interface{}{
//   		&IntegrationProperty{
//   			PagerDutyConfiguration: &PagerDutyConfigurationProperty{
//   				Name: jsii.String("name"),
//   				PagerDutyIncidentConfiguration: &PagerDutyIncidentConfigurationProperty{
//   					ServiceId: jsii.String("serviceId"),
//   				},
//   				SecretId: jsii.String("secretId"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
	Engagements *[]*string `field:"optional" json:"engagements" yaml:"engagements"`
	// Information about third-party services integrated into the response plan.
	Integrations interface{} `field:"optional" json:"integrations" yaml:"integrations"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

