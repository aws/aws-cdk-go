package awsecs


// The amount of ephemeral storage to allocate for the task.
//
// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate . For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon ECS Developer Guide;* .
//
// > For tasks using the Fargate launch type, the task requires the following platforms:
// >
// > - Linux platform version `1.4.0` or later.
// > - Windows platform version `1.0.0` or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &EphemeralStorageProperty{
//   	SizeInGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ephemeralstorage.html
//
type CfnTaskDefinition_EphemeralStorageProperty struct {
	// The total amount, in GiB, of ephemeral storage to set for the task.
	//
	// The minimum supported value is `20` GiB and the maximum supported value is `200` GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ephemeralstorage.html#cfn-ecs-taskdefinition-ephemeralstorage-sizeingib
	//
	SizeInGiB *float64 `field:"optional" json:"sizeInGiB" yaml:"sizeInGiB"`
}

