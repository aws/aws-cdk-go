package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   cfnTaskTemplateProps := &CfnTaskTemplateProps{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	Constraints: constraints,
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   	Defaults: []interface{}{
//   		&DefaultFieldValueProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Id: &FieldIdentifierProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Fields: []interface{}{
//   		&FieldProperty{
//   			Id: &FieldIdentifierProperty{
//   				Name: jsii.String("name"),
//   			},
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			SingleSelectOptions: []*string{
//   				jsii.String("singleSelectOptions"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SelfAssignContactFlowArn: jsii.String("selfAssignContactFlowArn"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html
//
type CfnTaskTemplateProps struct {
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Constraints that are applicable to the fields listed.
	//
	// The values can be represented in either JSON or YAML format. For an example of the JSON configuration, see *Examples* at the bottom of this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-constraints
	//
	Constraints interface{} `field:"optional" json:"constraints" yaml:"constraints"`
	// The Amazon Resource Name (ARN) of the flow that runs by default when a task is created by referencing this template.
	//
	// `ContactFlowArn` is not required when there is a field with `fieldType` = `QUICK_CONNECT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-contactflowarn
	//
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The default values for fields when a task is created by referencing this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-defaults
	//
	Defaults interface{} `field:"optional" json:"defaults" yaml:"defaults"`
	// The description of the task template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Fields that are part of the template.
	//
	// A template requires at least one field that has type `Name` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The name of the task template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-selfassigncontactflowarn
	//
	SelfAssignContactFlowArn *string `field:"optional" json:"selfAssignContactFlowArn" yaml:"selfAssignContactFlowArn"`
	// The status of the task template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-tasktemplate.html#cfn-connect-tasktemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

