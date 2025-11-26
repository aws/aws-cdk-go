package previewawsconnectmixins


// A list of actions to be run when the rule is triggered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var assignContactCategoryActions interface{}
//   var emptyValue interface{}
//   var endAssociatedTasksActions interface{}
//
//   actionsProperty := &ActionsProperty{
//   	AssignContactCategoryActions: []interface{}{
//   		assignContactCategoryActions,
//   	},
//   	CreateCaseActions: []interface{}{
//   		&CreateCaseActionProperty{
//   			Fields: []interface{}{
//   				&FieldProperty{
//   					Id: jsii.String("id"),
//   					Value: &FieldValueProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   			TemplateId: jsii.String("templateId"),
//   		},
//   	},
//   	EndAssociatedTasksActions: []interface{}{
//   		endAssociatedTasksActions,
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
//   			Subject: jsii.String("subject"),
//   		},
//   	},
//   	SubmitAutoEvaluationActions: []interface{}{
//   		&SubmitAutoEvaluationActionProperty{
//   			EvaluationFormArn: jsii.String("evaluationFormArn"),
//   		},
//   	},
//   	TaskActions: []interface{}{
//   		&TaskActionProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			References: map[string]interface{}{
//   				"referencesKey": &ReferenceProperty{
//   					"type": jsii.String("type"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	UpdateCaseActions: []interface{}{
//   		&UpdateCaseActionProperty{
//   			Fields: []interface{}{
//   				&FieldProperty{
//   					Id: jsii.String("id"),
//   					Value: &FieldValueProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html
//
type CfnRulePropsMixin_ActionsProperty struct {
	// Information about the contact category action.
	//
	// The syntax can be empty, for example, `{}` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-assigncontactcategoryactions
	//
	AssignContactCategoryActions interface{} `field:"optional" json:"assignContactCategoryActions" yaml:"assignContactCategoryActions"`
	// This action will create a case when a rule is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-createcaseactions
	//
	CreateCaseActions interface{} `field:"optional" json:"createCaseActions" yaml:"createCaseActions"`
	// This action will end associated tasks when a rule is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-endassociatedtasksactions
	//
	EndAssociatedTasksActions interface{} `field:"optional" json:"endAssociatedTasksActions" yaml:"endAssociatedTasksActions"`
	// Information about the EventBridge action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-eventbridgeactions
	//
	EventBridgeActions interface{} `field:"optional" json:"eventBridgeActions" yaml:"eventBridgeActions"`
	// Information about the send notification action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-sendnotificationactions
	//
	SendNotificationActions interface{} `field:"optional" json:"sendNotificationActions" yaml:"sendNotificationActions"`
	// This action will submit an auto contact evaluation when a rule is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-submitautoevaluationactions
	//
	SubmitAutoEvaluationActions interface{} `field:"optional" json:"submitAutoEvaluationActions" yaml:"submitAutoEvaluationActions"`
	// Information about the task action.
	//
	// This field is required if `TriggerEventSource` is one of the following values: `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-taskactions
	//
	TaskActions interface{} `field:"optional" json:"taskActions" yaml:"taskActions"`
	// This action will update a case when a rule is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-actions.html#cfn-connect-rule-actions-updatecaseactions
	//
	UpdateCaseActions interface{} `field:"optional" json:"updateCaseActions" yaml:"updateCaseActions"`
}

