// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props common to all ComputeEnvironments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   computeEnvironmentProps := &ComputeEnvironmentProps{
//   	ComputeEnvironmentName: jsii.String("computeEnvironmentName"),
//   	Enabled: jsii.Boolean(false),
//   	ServiceRole: role,
//   }
//
// Experimental.
type ComputeEnvironmentProps struct {
	// The name of the ComputeEnvironment.
	// Experimental.
	ComputeEnvironmentName *string `field:"optional" json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// Whether or not this ComputeEnvironment can accept jobs from a Queue.
	//
	// Enabled ComputeEnvironments can accept jobs from a Queue and
	// can scale instances up or down.
	// Disabled ComputeEnvironments cannot accept jobs from a Queue or
	// scale instances up or down.
	//
	// If you change a ComputeEnvironment from enabled to disabled while it is executing jobs,
	// Jobs in the `STARTED` or `RUNNING` states will not
	// be interrupted. As jobs complete, the ComputeEnvironment will scale instances down to `minvCpus`.
	//
	// To ensure you aren't billed for unused capacity, set `minvCpus` to `0`.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The role Batch uses to perform actions on your behalf in your account, such as provision instances to run your jobs.
	// Experimental.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
}

