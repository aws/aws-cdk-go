package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Model aws.sagemaker@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		AlgorithmSpecification: &AlgorithmSpecification{
//   			TrainingImage: []*string{
//   				jsii.String("trainingImage"),
//   			},
//   			TrainingInputMode: []*string{
//   				jsii.String("trainingInputMode"),
//   			},
//   		},
//   		EnableInterContainerTrafficEncryption: []*string{
//   			jsii.String("enableInterContainerTrafficEncryption"),
//   		},
//   		EnableManagedSpotTraining: []*string{
//   			jsii.String("enableManagedSpotTraining"),
//   		},
//   		EnableNetworkIsolation: []*string{
//   			jsii.String("enableNetworkIsolation"),
//   		},
//   		EndpointConfigName: []*string{
//   			jsii.String("endpointConfigName"),
//   		},
//   		EndpointName: []*string{
//   			jsii.String("endpointName"),
//   		},
//   		ExecutionRoleArn: []*string{
//   			jsii.String("executionRoleArn"),
//   		},
//   		HyperParameters: &HyperParameters{
//   			EvalMetric: []*string{
//   				jsii.String("evalMetric"),
//   			},
//   			NumRound: []*string{
//   				jsii.String("numRound"),
//   			},
//   			Objective: []*string{
//   				jsii.String("objective"),
//   			},
//   		},
//   		InputDataConfig: []RequestParametersItem2{
//   			&RequestParametersItem2{
//   				ChannelName: []*string{
//   					jsii.String("channelName"),
//   				},
//   				ContentType: []*string{
//   					jsii.String("contentType"),
//   				},
//   				DataSource: &DataSource1{
//   					S3DataSource: &S3DataSource1{
//   						S3DataDistributionType: []*string{
//   							jsii.String("s3DataDistributionType"),
//   						},
//   						S3DataType: []*string{
//   							jsii.String("s3DataType"),
//   						},
//   						S3Uri: []*string{
//   							jsii.String("s3Uri"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		ModelName: []*string{
//   			jsii.String("modelName"),
//   		},
//   		OutputDataConfig: &OutputDataConfig{
//   			RemoveJobNameFromS3OutputPath: []*string{
//   				jsii.String("removeJobNameFromS3OutputPath"),
//   			},
//   			S3OutputPath: []*string{
//   				jsii.String("s3OutputPath"),
//   			},
//   		},
//   		PrimaryContainer: &PrimaryContainer{
//   			ContainerHostname: []*string{
//   				jsii.String("containerHostname"),
//   			},
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			ModelDataUrl: []*string{
//   				jsii.String("modelDataUrl"),
//   			},
//   		},
//   		ProductionVariants: []RequestParametersItem{
//   			&RequestParametersItem{
//   				InitialInstanceCount: []*string{
//   					jsii.String("initialInstanceCount"),
//   				},
//   				InitialVariantWeight: []*string{
//   					jsii.String("initialVariantWeight"),
//   				},
//   				InstanceType: []*string{
//   					jsii.String("instanceType"),
//   				},
//   				ModelName: []*string{
//   					jsii.String("modelName"),
//   				},
//   				VariantName: []*string{
//   					jsii.String("variantName"),
//   				},
//   			},
//   		},
//   		ResourceConfig: &ResourceConfig{
//   			InstanceCount: []*string{
//   				jsii.String("instanceCount"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			VolumeSizeInGb: []*string{
//   				jsii.String("volumeSizeInGb"),
//   			},
//   		},
//   		RoleArn: []*string{
//   			jsii.String("roleArn"),
//   		},
//   		StoppingCondition: &StoppingCondition{
//   			MaxRuntimeInSeconds: []*string{
//   				jsii.String("maxRuntimeInSeconds"),
//   			},
//   		},
//   		Tags: []interface{}{
//   			tags,
//   		},
//   		TrainingJobName: []*string{
//   			jsii.String("trainingJobName"),
//   		},
//   		TransformInput: &TransformInput{
//   			CompressionType: []*string{
//   				jsii.String("compressionType"),
//   			},
//   			ContentType: []*string{
//   				jsii.String("contentType"),
//   			},
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3DataType: []*string{
//   						jsii.String("s3DataType"),
//   					},
//   					S3Uri: []*string{
//   						jsii.String("s3Uri"),
//   					},
//   				},
//   			},
//   		},
//   		TransformJobName: []*string{
//   			jsii.String("transformJobName"),
//   		},
//   		TransformOutput: &TransformOutput{
//   			S3OutputPath: []*string{
//   				jsii.String("s3OutputPath"),
//   			},
//   		},
//   		TransformResources: &TransformResources{
//   			InstanceCount: []*string{
//   				jsii.String("instanceCount"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		EndpointConfigArn: []*string{
//   			jsii.String("endpointConfigArn"),
//   		},
//   		ModelArn: []*string{
//   			jsii.String("modelArn"),
//   		},
//   		TrainingJobArn: []*string{
//   			jsii.String("trainingJobArn"),
//   		},
//   		TransformJobArn: []*string{
//   			jsii.String("transformJobArn"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		InvokedBy: []*string{
//   			jsii.String("invokedBy"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		SessionContext: &SessionContext{
//   			Attributes: &Attributes{
//   				CreationDate: []*string{
//   					jsii.String("creationDate"),
//   				},
//   				MfaAuthenticated: []*string{
//   					jsii.String("mfaAuthenticated"),
//   				},
//   			},
//   			SessionIssuer: &SessionIssuer{
//   				AccountId: []*string{
//   					jsii.String("accountId"),
//   				},
//   				Arn: []*string{
//   					jsii.String("arn"),
//   				},
//   				PrincipalId: []*string{
//   					jsii.String("principalId"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				UserName: []*string{
//   					jsii.String("userName"),
//   				},
//   			},
//   			WebIdFederationData: []*string{
//   				jsii.String("webIdFederationData"),
//   			},
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *ModelEvents_AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *ModelEvents_AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *ModelEvents_AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}

