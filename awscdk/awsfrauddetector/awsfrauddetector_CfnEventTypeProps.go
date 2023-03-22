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
//   cfnEventTypeProps := &cfnEventTypeProps{
//   	entityTypes: []interface{}{
//   		&entityTypeProperty{
//   			arn: jsii.String("arn"),
//   			createdTime: jsii.String("createdTime"),
//   			description: jsii.String("description"),
//   			inline: jsii.Boolean(false),
//   			lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			name: jsii.String("name"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	eventVariables: []interface{}{
//   		&eventVariableProperty{
//   			arn: jsii.String("arn"),
//   			createdTime: jsii.String("createdTime"),
//   			dataSource: jsii.String("dataSource"),
//   			dataType: jsii.String("dataType"),
//   			defaultValue: jsii.String("defaultValue"),
//   			description: jsii.String("description"),
//   			inline: jsii.Boolean(false),
//   			lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			name: jsii.String("name"),
//   			tags: []*cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			variableType: jsii.String("variableType"),
//   		},
//   	},
//   	labels: []interface{}{
//   		&labelProperty{
//   			arn: jsii.String("arn"),
//   			createdTime: jsii.String("createdTime"),
//   			description: jsii.String("description"),
//   			inline: jsii.Boolean(false),
//   			lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			name: jsii.String("name"),
//   			tags: []*cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

