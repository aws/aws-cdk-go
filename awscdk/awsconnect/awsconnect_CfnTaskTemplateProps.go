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
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Constraints that are applicable to the fields listed.
	//
	// The values can be represented in either JSON or YAML format. For an example of the JSON configuration, see *Examples* at the bottom of this page.
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// The Amazon Resource Name (ARN) of the flow that runs by default when a task is created by referencing this template.
	//
	// `ContactFlowArn` is not required when there is a field with `fieldType` = `QUICK_CONNECT` .
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The default values for fields when a task is created by referencing this template.
	Defaults interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// The description of the task template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Fields that are part of the template.
	//
	// A template requires at least one field that has type `Name` .
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the task template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the task template.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

