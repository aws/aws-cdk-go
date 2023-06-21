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
//
//   cfnRuleProps := &CfnRuleProps{
//   	Actions: &ActionsProperty{
//   		AssignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// A list of actions to be run when the rule is triggered.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions of the rule.
	Function *string `field:"required" json:"function" yaml:"function"`
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the rule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The publish status of the rule.
	//
	// *Allowed values* : `DRAFT` | `PUBLISHED`.
	PublishStatus *string `field:"required" json:"publishStatus" yaml:"publishStatus"`
	// The event source to trigger the rule.
	TriggerEventSource interface{} `field:"required" json:"triggerEventSource" yaml:"triggerEventSource"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

