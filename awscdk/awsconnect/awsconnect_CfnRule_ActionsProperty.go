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
//   actionsProperty := &actionsProperty{
//   	assignContactCategoryActions: []interface{}{
//   		assignContactCategoryActions,
//   	},
//   	eventBridgeActions: []interface{}{
//   		&eventBridgeActionProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	sendNotificationActions: []interface{}{
//   		&sendNotificationActionProperty{
//   			content: jsii.String("content"),
//   			contentType: jsii.String("contentType"),
//   			deliveryMethod: jsii.String("deliveryMethod"),
//   			recipient: &notificationRecipientTypeProperty{
//   				userArns: []*string{
//   					jsii.String("userArns"),
//   				},
//   				userTags: map[string]*string{
//   					"userTagsKey": jsii.String("userTags"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			subject: jsii.String("subject"),
//   		},
//   	},
//   	taskActions: []interface{}{
//   		&taskActionProperty{
//   			contactFlowArn: jsii.String("contactFlowArn"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			references: map[string]interface{}{
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

