package previewawsfrauddetectormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The event type details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventTypeProperty := &EventTypeProperty{
//   	Arn: jsii.String("arn"),
//   	CreatedTime: jsii.String("createdTime"),
//   	Description: jsii.String("description"),
//   	EntityTypes: []interface{}{
//   		&EntityTypeProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	EventVariables: []interface{}{
//   		&EventVariableProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			DataSource: jsii.String("dataSource"),
//   			DataType: jsii.String("dataType"),
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			VariableType: jsii.String("variableType"),
//   		},
//   	},
//   	Inline: jsii.Boolean(false),
//   	Labels: []interface{}{
//   		&LabelProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html
//
type CfnDetectorPropsMixin_EventTypeProperty struct {
	// The entity type ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Timestamp of when the event type was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-createdtime
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The event type description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The event type entity types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-entitytypes
	//
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
	// The event type event variables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-eventvariables
	//
	EventVariables interface{} `field:"optional" json:"eventVariables" yaml:"eventVariables"`
	// Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack.
	//
	// If the value is `true` , CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false` , CloudFormation will validate that the object exists and then use it within the resource without making changes to the object.
	//
	// For example, when creating `AWS::FraudDetector::Detector` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will create/update/delete the Variables as part of stack operations. However, if you set `Inline=false` , CloudFormation will associate the variables to your detector but not execute any changes to the variables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-inline
	//
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// The event type labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Timestamp of when the event type was last updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-lastupdatedtime
	//
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// The event type name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-frauddetector-detector-eventtype.html#cfn-frauddetector-detector-eventtype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

