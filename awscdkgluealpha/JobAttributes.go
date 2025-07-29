package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A subset of Job attributes are required for importing an existing job into a CDK project.
//
// This is only used when using fromJobAttributes
// to identify and reference the existing job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   jobAttributes := &JobAttributes{
//   	JobName: jsii.String("jobName"),
//
//   	// the properties below are optional
//   	Role: role,
//   }
//
// Experimental.
type JobAttributes struct {
	// The name of the job.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The IAM role assumed by Glue to run this job.
	// Default: - undefined.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

