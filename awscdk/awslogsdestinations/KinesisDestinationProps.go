package awslogsdestinations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Customize the Kinesis Logs Destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   kinesisDestinationProps := &KinesisDestinationProps{
//   	Role: role,
//   }
//
type KinesisDestinationProps struct {
	// The role to assume to write log events to the destination.
	// Default: - A new Role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

