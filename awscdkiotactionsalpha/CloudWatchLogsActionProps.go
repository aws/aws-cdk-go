package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
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

