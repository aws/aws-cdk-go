package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDestinationProps := &cfnDestinationProps{
//   	expression: jsii.String("expression"),
//   	expressionType: jsii.String("expressionType"),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDestinationProps struct {
	// The rule name to send messages to.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The type of value in `Expression` .
	ExpressionType *string `field:"required" json:"expressionType" yaml:"expressionType"`
	// The name of the new resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM Role that authorizes the destination.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the new resource.
	//
	// Maximum length is 2048 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

