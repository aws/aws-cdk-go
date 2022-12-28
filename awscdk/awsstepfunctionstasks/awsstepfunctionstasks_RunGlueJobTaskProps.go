package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for RunGlueJobTask.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   runGlueJobTaskProps := &runGlueJobTaskProps{
//   	arguments: map[string]*string{
//   		"argumentsKey": jsii.String("arguments"),
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	notifyDelayAfter: duration,
//   	securityConfiguration: jsii.String("securityConfiguration"),
//   	timeout: duration,
//   }
//
// Deprecated: use `GlueStartJobRun`.
type RunGlueJobTaskProps struct {
	// The job arguments specifically for this run.
	//
	// For this job run, they replace the default arguments set in the job definition itself.
	// Deprecated: use `GlueStartJobRun`.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// The service integration pattern indicates different ways to start the Glue job.
	//
	// The valid value for Glue is either FIRE_AND_FORGET or SYNC.
	// Deprecated: use `GlueStartJobRun`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	//
	// Must be at least 1 minute.
	// Deprecated: use `GlueStartJobRun`.
	NotifyDelayAfter awscdk.Duration `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
	// The name of the SecurityConfiguration structure to be used with this job run.
	//
	// This must match the Glue API
	// [single-line string pattern](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html#aws-glue-api-regex-oneLine).
	// Deprecated: use `GlueStartJobRun`.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The job run timeout.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Must be at least 1 minute.
	// Deprecated: use `GlueStartJobRun`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

