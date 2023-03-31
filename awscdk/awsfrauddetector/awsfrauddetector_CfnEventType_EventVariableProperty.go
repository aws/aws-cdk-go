package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The variables associated with this event type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventVariableProperty := &eventVariableProperty{
//   	arn: jsii.String("arn"),
//   	createdTime: jsii.String("createdTime"),
//   	dataSource: jsii.String("dataSource"),
//   	dataType: jsii.String("dataType"),
//   	defaultValue: jsii.String("defaultValue"),
//   	description: jsii.String("description"),
//   	inline: jsii.Boolean(false),
//   	lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	variableType: jsii.String("variableType"),
//   }
//
type CfnEventType_EventVariableProperty struct {
	// The event variable ARN.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Timestamp for when event variable was created.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The source of the event variable.
	//
	// Valid values: `EVENT | EXTERNAL_MODEL_SCORE`
	//
	// When defining a variable within a event type, you can only use the `EVENT` value for DataSource when the *Inline* property is set to true. If the *Inline* property is set false, you can use either `EVENT` or `MODEL_SCORE` for DataSource.
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The data type of the event variable.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The default value of the event variable.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The event variable description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack.
	//
	// If the value is `true` , CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false` , CloudFormation will validate that the object exists and then use it within the resource without making changes to the object.
	//
	// For example, when creating `AWS::FraudDetector::EventType` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will create/update/delete the Variables as part of stack operations. However, if you set `Inline=false` , CloudFormation will associate the variables to your event type but not execute any changes to the variables.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// Timestamp for when the event variable was last updated.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// The name of the event variable.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of event variable.
	//
	// For more information, see [Variable types](https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types) .
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

