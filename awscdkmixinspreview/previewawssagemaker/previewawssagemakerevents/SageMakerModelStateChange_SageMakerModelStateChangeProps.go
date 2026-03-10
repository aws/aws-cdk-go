package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerModelStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelStateChangeProps := &SageMakerModelStateChangeProps{
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	ExecutionRoleArn: []*string{
//   		jsii.String("executionRoleArn"),
//   	},
//   	ModelArn: []*string{
//   		jsii.String("modelArn"),
//   	},
//   	ModelName: []*string{
//   		jsii.String("modelName"),
//   	},
//   	PrimaryContainer: &PrimaryContainer{
//   		ContainerHostname: []*string{
//   			jsii.String("containerHostname"),
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		ModelDataUrl: []*string{
//   			jsii.String("modelDataUrl"),
//   		},
//   	},
//   	Tags: []Tags{
//   		&Tags{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type SageMakerModelStateChange_SageMakerModelStateChangeProps struct {
	// CreationTime property.
	//
	// Specify an array of string values to match this event if the actual value of CreationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// ExecutionRoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of ExecutionRoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionRoleArn *[]*string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// ModelArn property.
	//
	// Specify an array of string values to match this event if the actual value of ModelArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelArn *[]*string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// ModelName property.
	//
	// Specify an array of string values to match this event if the actual value of ModelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelName *[]*string `field:"optional" json:"modelName" yaml:"modelName"`
	// PrimaryContainer property.
	//
	// Specify an array of string values to match this event if the actual value of PrimaryContainer is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrimaryContainer *SageMakerModelStateChange_PrimaryContainer `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerModelStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

