package awsiotactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Configuration properties of an action for CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   cloudWatchLogsActionProps := &CloudWatchLogsActionProps{
//   	Role: role,
//   }
//
// Experimental.
type CloudWatchLogsActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

