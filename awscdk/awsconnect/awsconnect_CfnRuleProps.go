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
//   cfnRuleProps := &cfnRuleProps{
//   	actions: &actionsProperty{
//   		assignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
//   		},
//   		eventBridgeActions: []interface{}{
//   			&eventBridgeActionProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		sendNotificationActions: []interface{}{
//   			&sendNotificationActionProperty{
//   				content: jsii.String("content"),
//   				contentType: jsii.String("contentType"),
//   				deliveryMethod: jsii.String("deliveryMethod"),
//   				recipient: &notificationRecipientTypeProperty{
//   					userArns: []*string{
//   						jsii.String("userArns"),
//   					},
//   					userTags: map[string]*string{
//   						"userTagsKey": jsii.String("userTags"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				subject: jsii.String("subject"),
//   			},
//   		},
//   		taskActions: []interface{}{
//   			&taskActionProperty{
//   				contactFlowArn: jsii.String("contactFlowArn"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				description: jsii.String("description"),
//   				references: map[string]interface{}{
//   					"referencesKey": &ReferenceProperty{
//   						"type": jsii.String("type"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	function: jsii.String("function"),
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//   	publishStatus: jsii.String("publishStatus"),
//   	triggerEventSource: &ruleTriggerEventSourceProperty{
//   		eventSourceName: jsii.String("eventSourceName"),
//
//   		// the properties below are optional
//   		integrationAssociationArn: jsii.String("integrationAssociationArn"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

