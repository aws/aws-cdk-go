package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerAlgorithmStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerAlgorithmStateChangeProps := &SageMakerAlgorithmStateChangeProps{
//   	AlgorithmArn: []*string{
//   		jsii.String("algorithmArn"),
//   	},
//   	AlgorithmDescription: []*string{
//   		jsii.String("algorithmDescription"),
//   	},
//   	AlgorithmName: []*string{
//   		jsii.String("algorithmName"),
//   	},
//   	CertifyForMarketplace: []*string{
//   		jsii.String("certifyForMarketplace"),
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
type SageMakerAlgorithmStateChange_SageMakerAlgorithmStateChangeProps struct {
	// AlgorithmArn property.
	//
	// Specify an array of string values to match this event if the actual value of AlgorithmArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmArn *[]*string `field:"optional" json:"algorithmArn" yaml:"algorithmArn"`
	// AlgorithmDescription property.
	//
	// Specify an array of string values to match this event if the actual value of AlgorithmDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmDescription *[]*string `field:"optional" json:"algorithmDescription" yaml:"algorithmDescription"`
	// AlgorithmName property.
	//
	// Specify an array of string values to match this event if the actual value of AlgorithmName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AlgorithmName *[]*string `field:"optional" json:"algorithmName" yaml:"algorithmName"`
	// CertifyForMarketplace property.
	//
	// Specify an array of string values to match this event if the actual value of CertifyForMarketplace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CertifyForMarketplace *[]*string `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerAlgorithmStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

