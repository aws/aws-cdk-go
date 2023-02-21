package awsconnect


// Information about the task action.
//
// This field is required if `TriggerEventSource` is one of the following values: `OnZendeskTicketCreate` | `OnZendeskTicketStatusUpdate` | `OnSalesforceCaseCreate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskActionProperty := &TaskActionProperty{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	References: map[string]interface{}{
//   		"referencesKey": &ReferenceProperty{
//   			"type": jsii.String("type"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRule_TaskActionProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The name.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the reference when the `referenceType` is `URL` .
	//
	// Otherwise, null. `URL` is the only accepted type. (Supports variable injection in the `Value` field.)
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

