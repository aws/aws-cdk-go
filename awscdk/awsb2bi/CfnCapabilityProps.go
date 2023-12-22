package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapability`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapabilityProps := &CfnCapabilityProps{
//   	Configuration: &CapabilityConfigurationProperty{
//   		Edi: &EdiConfigurationProperty{
//   			InputLocation: &S3LocationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   			},
//   			OutputLocation: &S3LocationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Key: jsii.String("key"),
//   			},
//   			TransformerId: jsii.String("transformerId"),
//   			Type: &EdiTypeProperty{
//   				X12Details: &X12DetailsProperty{
//   					TransactionSet: jsii.String("transactionSet"),
//   					Version: jsii.String("version"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	InstructionsDocuments: []interface{}{
//   		&S3LocationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html
//
type CfnCapabilityProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-instructionsdocuments
	//
	InstructionsDocuments interface{} `field:"optional" json:"instructionsDocuments" yaml:"instructionsDocuments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

