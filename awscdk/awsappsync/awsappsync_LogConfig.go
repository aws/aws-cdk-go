package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Logging configuration for AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   logConfig := &logConfig{
//   	excludeVerboseContent: jsii.Boolean(false),
//   	fieldLogLevel: awscdk.Aws_appsync.fieldLogLevel_NONE,
//   	role: role,
//   }
//
// Experimental.
type LogConfig struct {
	// exclude verbose content.
	// Experimental.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// log level for fields.
	// Experimental.
	FieldLogLevel FieldLogLevel `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// The role for CloudWatch Logs.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

