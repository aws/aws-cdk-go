// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Attributes for importing `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   jobAttributes := &jobAttributes{
//   	jobName: jsii.String("jobName"),
//
//   	// the properties below are optional
//   	role: role,
//   }
//
// Experimental.
type JobAttributes struct {
	// The name of the job.
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The IAM role assumed by Glue to run this job.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

