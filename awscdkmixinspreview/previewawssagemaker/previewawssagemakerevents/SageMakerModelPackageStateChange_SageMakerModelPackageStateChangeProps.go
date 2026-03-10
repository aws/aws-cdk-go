package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerModelPackageStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelPackageStateChangeProps := &SageMakerModelPackageStateChangeProps{
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
//   	ModelPackageArn: []*string{
//   		jsii.String("modelPackageArn"),
//   	},
//   	ModelPackageDescription: []*string{
//   		jsii.String("modelPackageDescription"),
//   	},
//   	ModelPackageName: []*string{
//   		jsii.String("modelPackageName"),
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
type SageMakerModelPackageStateChange_SageMakerModelPackageStateChangeProps struct {
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
	// ModelPackageArn property.
	//
	// Specify an array of string values to match this event if the actual value of ModelPackageArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelPackageArn *[]*string `field:"optional" json:"modelPackageArn" yaml:"modelPackageArn"`
	// ModelPackageDescription property.
	//
	// Specify an array of string values to match this event if the actual value of ModelPackageDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelPackageDescription *[]*string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// ModelPackageName property.
	//
	// Specify an array of string values to match this event if the actual value of ModelPackageName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ModelPackageName *[]*string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Tags property.
	//
	// Specify an array of string values to match this event if the actual value of Tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*SageMakerModelPackageStateChange_Tags `field:"optional" json:"tags" yaml:"tags"`
}

