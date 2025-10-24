package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assignContactCategoryActions interface{}
//   var emptyValue interface{}
//   var endAssociatedTasksActions interface{}
//
//   cfnRuleProps := &CfnRuleProps{
//   	Actions: &ActionsProperty{
//   		AssignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
//   		},
//   		CreateCaseActions: []interface{}{
//   			&CreateCaseActionProperty{
//   				Fields: []interface{}{
//   					&FieldProperty{
//   						Id: jsii.String("id"),
//   						Value: &FieldValueProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   				TemplateId: jsii.String("templateId"),
//   			},
//   		},
//   		EndAssociatedTasksActions: []interface{}{
//   			endAssociatedTasksActions,
//   		},
//   		EventBridgeActions: []interface{}{
//   			&EventBridgeActionProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		SendNotificationActions: []interface{}{
//   			&SendNotificationActionProperty{
//   				Content: jsii.String("content"),
//   				ContentType: jsii.String("contentType"),
//   				DeliveryMethod: jsii.String("deliveryMethod"),
//   				Recipient: &NotificationRecipientTypeProperty{
//   					UserArns: []*string{
//   						jsii.String("userArns"),
//   					},
//   					UserTags: map[string]*string{
//   						"userTagsKey": jsii.String("userTags"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Subject: jsii.String("subject"),
//   			},
//   		},
//   		SubmitAutoEvaluationActions: []interface{}{
//   			&SubmitAutoEvaluationActionProperty{
//   				EvaluationFormArn: jsii.String("evaluationFormArn"),
//   			},
//   		},
//   		TaskActions: []interface{}{
//   			&TaskActionProperty{
//   				ContactFlowArn: jsii.String("contactFlowArn"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   				References: map[string]interface{}{
//   					"referencesKey": &ReferenceProperty{
//   						"type": jsii.String("type"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UpdateCaseActions: []interface{}{
//   			&UpdateCaseActionProperty{
//   				Fields: []interface{}{
//   					&FieldProperty{
//   						Id: jsii.String("id"),
//   						Value: &FieldValueProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Function: jsii.String("function"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	PublishStatus: jsii.String("publishStatus"),
//   	TriggerEventSource: &RuleTriggerEventSourceProperty{
//   		EventSourceName: jsii.String("eventSourceName"),
//
//   		// the properties below are optional
//   		IntegrationAssociationArn: jsii.String("integrationAssociationArn"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html
//
type CfnRuleProps struct {
	// A list of actions to be run when the rule is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-function
	//
	Function *string `field:"required" json:"function" yaml:"function"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The publish status of the rule.
	//
	// *Allowed values* : `DRAFT` | `PUBLISHED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-publishstatus
	//
	PublishStatus *string `field:"required" json:"publishStatus" yaml:"publishStatus"`
	// The event source to trigger the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-triggereventsource
	//
	TriggerEventSource interface{} `field:"required" json:"triggerEventSource" yaml:"triggerEventSource"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

