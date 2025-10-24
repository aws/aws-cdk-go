package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options when binding a SchemaConfig to a Destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   schemaConfigurationBindOptions := &SchemaConfigurationBindOptions{
//   	Role: role,
//   }
//
type SchemaConfigurationBindOptions struct {
	// The IAM Role that will be used by the Delivery Stream for access to the Glue data catalog for record format conversion.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

