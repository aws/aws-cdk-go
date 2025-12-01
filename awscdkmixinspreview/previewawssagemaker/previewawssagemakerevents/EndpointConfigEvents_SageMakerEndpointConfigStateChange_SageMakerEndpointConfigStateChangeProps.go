package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for EndpointConfig aws.sagemaker@SageMakerEndpointConfigStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerEndpointConfigStateChangeProps := &SageMakerEndpointConfigStateChangeProps{
//   	CreationTime: []*string{
//   		jsii.String("creationTime"),
//   	},
//   	EndpointConfigArn: []*string{
//   		jsii.String("endpointConfigArn"),
//   	},
//   	EndpointConfigName: []*string{
//   		jsii.String("endpointConfigName"),
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
//   	ProductionVariants: []SageMakerEndpointConfigStateChangeItem{
//   		&SageMakerEndpointConfigStateChangeItem{
//   			InitialInstanceCount: []*string{
//   				jsii.String("initialInstanceCount"),
//   			},
//   			InitialVariantWeight: []*string{
//   				jsii.String("initialVariantWeight"),
//   			},
//   			InstanceType: []*string{
//   				jsii.String("instanceType"),
//   			},
//   			ModelName: []*string{
//   				jsii.String("modelName"),
//   			},
//   			VariantName: []*string{
//   				jsii.String("variantName"),
//   			},
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
type EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps struct {
	// CreationTime property.
	//
	// Specify an array of string values to match this event if the actual value of CreationTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationTime *[]*string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// EndpointConfigArn property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointConfigArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointConfigArn *[]*string `field:"optional" json:"endpointConfigArn" yaml:"endpointConfigArn"`
	// EndpointConfigName property.
	//
	// Specify an array of string values to match this event if the actual value of EndpointConfigName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the EndpointConfig reference.
	//
	// Experimental.
	EndpointConfigName *[]*string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// ProductionVariants property.
	//
	// Specify an array of string values to match this event if the actual value of ProductionVariants is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProductionVariants *[]*EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeItem `field:"optional" json:"productionVariants" yaml:"productionVariants"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*EndpointConfigEvents_SageMakerEndpointConfigStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

