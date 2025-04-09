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
//   var role role
//
//   dataProcessorBindOptions := &DataProcessorBindOptions{
//   	Role: role,
//   }
//
type DataProcessorBindOptions struct {
	// The IAM role assumed by Amazon Data Firehose to write to the destination that this DataProcessor will bind to.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

