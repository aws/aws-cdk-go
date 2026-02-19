package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options when binding a DataProcessor to a delivery stream destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   dataProcessorBindOptions := &DataProcessorBindOptions{
//   	Role: role,
//
//   	// the properties below are optional
//   	DynamicPartitioningEnabled: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
type DataProcessorBindOptions struct {
	// The IAM role assumed by Amazon Data Firehose to write to the destination that this DataProcessor will bind to.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Whether the dynamic partitioning is enabled.
	// Default: false.
	//
	DynamicPartitioningEnabled *bool `field:"optional" json:"dynamicPartitioningEnabled" yaml:"dynamicPartitioningEnabled"`
	// S3 bucket prefix.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

