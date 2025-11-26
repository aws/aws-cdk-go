package previewawsfrauddetectormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventTypeMixinProps := &CfnEventTypeMixinProps{
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
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html
//
type CfnEventTypeMixinProps struct {
	// The event type description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The event type entity types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-entitytypes
	//
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
	// The event type event variables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-eventvariables
	//
	EventVariables interface{} `field:"optional" json:"eventVariables" yaml:"eventVariables"`
	// The event type labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The event type name.
	//
	// Pattern : `^[0-9a-z_-]+$`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-eventtype.html#cfn-frauddetector-eventtype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

