package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskActionProperty := &taskActionProperty{
//   	contactFlowArn: jsii.String("contactFlowArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	references: map[string]interface{}{
//   		"referencesKey": &ReferenceProperty{
//   			"type": jsii.String("type"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRule_TaskActionProperty struct {
	// `CfnRule.TaskActionProperty.ContactFlowArn`.
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// `CfnRule.TaskActionProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnRule.TaskActionProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnRule.TaskActionProperty.References`.
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

