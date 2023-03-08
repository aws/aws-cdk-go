package awsecs


// The amount of ephemeral storage to allocate for the task.
//
// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate . For more information, see [Fargate task storage](https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html) in the *Amazon ECS User Guide for AWS Fargate* .
//
// > This parameter is only supported for tasks hosted on Fargate using Linux platform version `1.4.0` or later. This parameter is not supported for Windows containers on Fargate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &ephemeralStorageProperty{
//   	sizeInGiB: jsii.Number(123),
//   }
//
type CfnTaskDefinition_EphemeralStorageProperty struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.
	SizeInGiB *float64 `field:"optional" json:"sizeInGiB" yaml:"sizeInGiB"`
}

