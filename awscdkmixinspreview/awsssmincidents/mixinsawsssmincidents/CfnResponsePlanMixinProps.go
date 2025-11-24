package mixinsawsssmincidents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResponsePlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponsePlanMixinProps := &CfnResponsePlanMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			SsmAutomation: &SsmAutomationProperty{
//   				DocumentName: jsii.String("documentName"),
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
//   				RoleArn: jsii.String("roleArn"),
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
//   	IncidentTemplate: &IncidentTemplateProperty{
//   		DedupeString: jsii.String("dedupeString"),
//   		Impact: jsii.Number(123),
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
//   		Title: jsii.String("title"),
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
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html
//
type CfnResponsePlanMixinProps struct {
	// The actions that the response plan starts at the beginning of an incident.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The  chat channel used for collaboration during an incident.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-chatchannel
	//
	ChatChannel interface{} `field:"optional" json:"chatChannel" yaml:"chatChannel"`
	// The human readable name of the response plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-engagements
	//
	Engagements *[]*string `field:"optional" json:"engagements" yaml:"engagements"`
	// Details used to create an incident when using this response plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-incidenttemplate
	//
	IncidentTemplate interface{} `field:"optional" json:"incidentTemplate" yaml:"incidentTemplate"`
	// Information about third-party services integrated into the response plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-integrations
	//
	Integrations interface{} `field:"optional" json:"integrations" yaml:"integrations"`
	// The name of the response plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

