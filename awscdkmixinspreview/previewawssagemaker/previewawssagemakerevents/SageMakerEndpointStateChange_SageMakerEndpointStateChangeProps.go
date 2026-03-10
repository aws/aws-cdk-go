package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerEndpointStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointStateChangeProps := &SageMakerEndpointStateChangeProps{
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	EndpointArn: []*string{
//   		jsii.String("endpointArn"),
//   	},
//   	EndpointConfigName: []*string{
//   		jsii.String("endpointConfigName"),
//   	},
//   	EndpointName: []*string{
//   		jsii.String("endpointName"),
//   	},
//   	EndpointStatus: []*string{
//   		jsii.String("endpointStatus"),
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
//   	LastModifiedTime: []*string{
//   		jsii.String("lastModifiedTime"),
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
type SageMakerEndpointStateChange_SageMakerEndpointStateChangeProps struct {
	// CreationTime property.
	//
	// Specify an array of string values to match this event if the actual value of CreationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// EndpointArn property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointArn *[]*string `field:"optional" json:"endpointArn" yaml:"endpointArn"`
	// EndpointConfigName property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointConfigName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointConfigName *[]*string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// EndpointName property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointName *[]*string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// EndpointStatus property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointStatus *[]*string `field:"optional" json:"endpointStatus" yaml:"endpointStatus"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// LastModifiedTime property.
	//
	// Specify an array of string values to match this event if the actual value of LastModifiedTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedTime *[]*string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerEndpointStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

