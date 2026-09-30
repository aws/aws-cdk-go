package awsbatch


// The EKS access entry configuration for the compute environment.
//
// Controls whether AWS Batch manages the EKS access entry for the compute environment's service role, or inherits it from the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eksAccessEntryProperty := &EksAccessEntryProperty{
//   	DesiredState: jsii.String("desiredState"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-eksaccessentry.html
//
type CfnComputeEnvironmentPropsMixin_EksAccessEntryProperty struct {
	// The desired state of the EKS access entry managed by AWS Batch.
	//
	// When omitted, AWS Batch applies INHERIT_FROM_CLUSTER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-eksaccessentry.html#cfn-batch-computeenvironment-eksaccessentry-desiredstate
	//
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The read-only status of the EKS access entry, returned by DescribeComputeEnvironments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-eksaccessentry.html#cfn-batch-computeenvironment-eksaccessentry-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

