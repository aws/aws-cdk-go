package previewawsconnectmixins


// Information about the task action.
//
// This field is required if `TriggerEventSource` is one of the following values: `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskActionProperty := &TaskActionProperty{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	References: map[string]interface{}{
//   		"referencesKey": &ReferenceProperty{
//   			"type": jsii.String("type"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-taskaction.html
//
type CfnRulePropsMixin_TaskActionProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-taskaction.html#cfn-connect-rule-taskaction-contactflowarn
	//
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The description.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-taskaction.html#cfn-connect-rule-taskaction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-taskaction.html#cfn-connect-rule-taskaction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the reference when the `referenceType` is `URL` .
	//
	// Otherwise, null. `URL` is the only accepted type. (Supports variable injection in the `Value` field.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-taskaction.html#cfn-connect-rule-taskaction-references
	//
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

