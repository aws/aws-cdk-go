package mixinsawsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapabilityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityMixinProps := &CfnCapabilityMixinProps{
//   	Configuration: &CapabilityConfigurationProperty{
//   		Edi: &EdiConfigurationProperty{
//   			CapabilityDirection: jsii.String("capabilityDirection"),
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
//   	InstructionsDocuments: []interface{}{
//   		&S3LocationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html
//
type CfnCapabilityMixinProps struct {
	// Specifies a structure that contains the details for a capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability.
	//
	// Each item contains the name of the bucket and the key, to identify the document's location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-instructionsdocuments
	//
	InstructionsDocuments interface{} `field:"optional" json:"instructionsDocuments" yaml:"instructionsDocuments"`
	// The display name of the capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.
	//
	// You can attach this metadata to resources (capabilities, partnerships, and so on) for any purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Returns the type of the capability.
	//
	// Currently, only `edi` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-capability.html#cfn-b2bi-capability-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

