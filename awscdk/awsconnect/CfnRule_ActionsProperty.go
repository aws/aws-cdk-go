package awsconnect


// A list of actions to be run when the rule is triggered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assignContactCategoryActions interface{}
//
//   actionsProperty := &ActionsProperty{
//   	AssignContactCategoryActions: []interface{}{
//   		assignContactCategoryActions,
//   	},
//   	EventBridgeActions: []interface{}{
//   		&EventBridgeActionProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	SendNotificationActions: []interface{}{
//   		&SendNotificationActionProperty{
//   			Content: jsii.String("content"),
//   			ContentType: jsii.String("contentType"),
//   			DeliveryMethod: jsii.String("deliveryMethod"),
//   			Recipient: &NotificationRecipientTypeProperty{
//   				UserArns: []*string{
//   					jsii.String("userArns"),
//   				},
//   				UserTags: map[string]*string{
//   					"userTagsKey": jsii.String("userTags"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Subject: jsii.String("subject"),
//   		},
//   	},
//   	TaskActions: []interface{}{
//   		&TaskActionProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			References: map[string]interface{}{
//   				"referencesKey": &ReferenceProperty{
//   					"type": jsii.String("type"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRule_ActionsProperty struct {
	// Information about the contact category action.
	//
	// The syntax can be empty, for example, `{}` .
	AssignContactCategoryActions interface{} `field:"optional" json:"assignContactCategoryActions" yaml:"assignContactCategoryActions"`
	// Information about the EventBridge action.
	EventBridgeActions interface{} `field:"optional" json:"eventBridgeActions" yaml:"eventBridgeActions"`
	// Information about the send notification action.
	SendNotificationActions interface{} `field:"optional" json:"sendNotificationActions" yaml:"sendNotificationActions"`
	// Information about the task action.
	//
	// This field is required if `TriggerEventSource` is one of the following values: `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
	TaskActions interface{} `field:"optional" json:"taskActions" yaml:"taskActions"`
}

