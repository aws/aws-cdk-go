package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTaskTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var constraints interface{}
//
//   cfnTaskTemplateProps := &cfnTaskTemplateProps{
//   	instanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	clientToken: jsii.String("clientToken"),
//   	constraints: constraints,
//   	contactFlowArn: jsii.String("contactFlowArn"),
//   	defaults: []interface{}{
//   		&defaultFieldValueProperty{
//   			defaultValue: jsii.String("defaultValue"),
//   			id: &fieldIdentifierProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	fields: []interface{}{
//   		&fieldProperty{
//   			id: &fieldIdentifierProperty{
//   				name: jsii.String("name"),
//   			},
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			singleSelectOptions: []*string{
//   				jsii.String("singleSelectOptions"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	status: jsii.String("status"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTaskTemplateProps struct {
	// `AWS::Connect::TaskTemplate.InstanceArn`.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// `AWS::Connect::TaskTemplate.ClientToken`.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// `AWS::Connect::TaskTemplate.Constraints`.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// `AWS::Connect::TaskTemplate.ContactFlowArn`.
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// `AWS::Connect::TaskTemplate.Defaults`.
	Defaults interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// `AWS::Connect::TaskTemplate.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Connect::TaskTemplate.Fields`.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// `AWS::Connect::TaskTemplate.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Connect::TaskTemplate.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `AWS::Connect::TaskTemplate.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

