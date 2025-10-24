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
//
//   			// the properties below are optional
//   			CapabilityDirection: jsii.String("capabilityDirection"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html
//
type CfnCapabilityProps struct {
	// Specifies a structure that contains the details for a capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The display name of the capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Returns the type of the capability.
	//
	// Currently, only `edi` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability.
	//
	// Each item contains the name of the bucket and the key, to identify the document's location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-instructionsdocuments
	//
	InstructionsDocuments interface{} `field:"optional" json:"instructionsDocuments" yaml:"instructionsDocuments"`
	// Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.
	//
	// You can attach this metadata to resources (capabilities, partnerships, and so on) for any purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

