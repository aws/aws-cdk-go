package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Model aws.sagemaker@SageMakerTransformJobStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerTransformJobStateChangeProps := &SageMakerTransformJobStateChangeProps{
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
//   	ModelName: []*string{
//   		jsii.String("modelName"),
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
//   	TransformEndTime: []*string{
//   		jsii.String("transformEndTime"),
//   	},
//   	TransformInput: &TransformInput{
//   		CompressionType: []*string{
//   			jsii.String("compressionType"),
//   		},
//   		ContentType: []*string{
//   			jsii.String("contentType"),
//   		},
//   		DataSource: &DataSource{
//   			S3DataSource: &S3DataSource{
//   				S3DataType: []*string{
//   					jsii.String("s3DataType"),
//   				},
//   				S3Uri: []*string{
//   					jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   		SplitType: []*string{
//   			jsii.String("splitType"),
//   		},
//   	},
//   	TransformJobArn: []*string{
//   		jsii.String("transformJobArn"),
//   	},
//   	TransformJobName: []*string{
//   		jsii.String("transformJobName"),
//   	},
//   	TransformJobStatus: []*string{
//   		jsii.String("transformJobStatus"),
//   	},
//   	TransformOutput: &TransformOutput{
//   		AssembleWith: []*string{
//   			jsii.String("assembleWith"),
//   		},
//   		S3OutputPath: []*string{
//   			jsii.String("s3OutputPath"),
//   		},
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: []*string{
//   			jsii.String("instanceCount"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   	},
//   	TransformStartTime: []*string{
//   		jsii.String("transformStartTime"),
//   	},
//   }
//
// Experimental.
type ModelEvents_SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps struct {
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
	// ModelName property.
	//
	// Specify an array of string values to match this event if the actual value of ModelName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Model reference.
	//
	// Experimental.
	ModelName *[]*string `field:"optional" json:"modelName" yaml:"modelName"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*ModelEvents_SageMakerTransformJobStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
	// TransformEndTime property.
	//
	// Specify an array of string values to match this event if the actual value of TransformEndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformEndTime *[]*string `field:"optional" json:"transformEndTime" yaml:"transformEndTime"`
	// TransformInput property.
	//
	// Specify an array of string values to match this event if the actual value of TransformInput is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformInput *ModelEvents_SageMakerTransformJobStateChange_TransformInput `field:"optional" json:"transformInput" yaml:"transformInput"`
	// TransformJobArn property.
	//
	// Specify an array of string values to match this event if the actual value of TransformJobArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformJobArn *[]*string `field:"optional" json:"transformJobArn" yaml:"transformJobArn"`
	// TransformJobName property.
	//
	// Specify an array of string values to match this event if the actual value of TransformJobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformJobName *[]*string `field:"optional" json:"transformJobName" yaml:"transformJobName"`
	// TransformJobStatus property.
	//
	// Specify an array of string values to match this event if the actual value of TransformJobStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformJobStatus *[]*string `field:"optional" json:"transformJobStatus" yaml:"transformJobStatus"`
	// TransformOutput property.
	//
	// Specify an array of string values to match this event if the actual value of TransformOutput is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformOutput *ModelEvents_SageMakerTransformJobStateChange_TransformOutput `field:"optional" json:"transformOutput" yaml:"transformOutput"`
	// TransformResources property.
	//
	// Specify an array of string values to match this event if the actual value of TransformResources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformResources *ModelEvents_SageMakerTransformJobStateChange_TransformResources `field:"optional" json:"transformResources" yaml:"transformResources"`
	// TransformStartTime property.
	//
	// Specify an array of string values to match this event if the actual value of TransformStartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformStartTime *[]*string `field:"optional" json:"transformStartTime" yaml:"transformStartTime"`
}

