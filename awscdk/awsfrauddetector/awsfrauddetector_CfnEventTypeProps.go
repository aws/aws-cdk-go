package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEventType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventTypeProps := &CfnEventTypeProps{
//   	EntityTypes: []interface{}{
//   		&EntityTypeProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []cfnTag{
//   				&cfnTag{
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
//   			Tags: []*cfnTag{
//   				&cfnTag{
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
//   			Tags: []*cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []*cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventTypeProps struct {
	// The event type entity types.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// The event type event variables.
	EventVariables interface{} `field:"required" json:"eventVariables" yaml:"eventVariables"`
	// The event type labels.
	Labels interface{} `field:"required" json:"labels" yaml:"labels"`
	// The event type name.
	//
	// Pattern : `^[0-9a-z_-]+$`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The event type description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

