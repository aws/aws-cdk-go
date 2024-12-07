package awschatbot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomActionProps := &CfnCustomActionProps{
//   	ActionName: jsii.String("actionName"),
//   	Definition: &CustomActionDefinitionProperty{
//   		CommandText: jsii.String("commandText"),
//   	},
//
//   	// the properties below are optional
//   	AliasName: jsii.String("aliasName"),
//   	Attachments: []interface{}{
//   		&CustomActionAttachmentProperty{
//   			ButtonText: jsii.String("buttonText"),
//   			Criteria: []interface{}{
//   				&CustomActionAttachmentCriteriaProperty{
//   					Operator: jsii.String("operator"),
//   					VariableName: jsii.String("variableName"),
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			NotificationType: jsii.String("notificationType"),
//   			Variables: map[string]*string{
//   				"variablesKey": jsii.String("variables"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html
//
type CfnCustomActionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html#cfn-chatbot-customaction-actionname
	//
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html#cfn-chatbot-customaction-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html#cfn-chatbot-customaction-aliasname
	//
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html#cfn-chatbot-customaction-attachments
	//
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-customaction.html#cfn-chatbot-customaction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

