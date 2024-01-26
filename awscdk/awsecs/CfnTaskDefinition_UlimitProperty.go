package awsecs


// The `ulimit` settings to pass to the container.
//
// Amazon ECS tasks hosted on AWS Fargate use the default resource limit values set by the operating system with the exception of the `nofile` resource limit parameter which AWS Fargate overrides. The `nofile` resource limit sets a restriction on the number of open files that a container can use. The default `nofile` soft limit is `1024` and the default hard limit is `65535` .
//
// You can specify the `ulimit` settings for a container in a task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimitProperty := &UlimitProperty{
//   	HardLimit: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SoftLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ulimit.html
//
type CfnTaskDefinition_UlimitProperty struct {
	// The hard limit for the `ulimit` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ulimit.html#cfn-ecs-taskdefinition-ulimit-hardlimit
	//
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The `type` of the `ulimit` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ulimit.html#cfn-ecs-taskdefinition-ulimit-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The soft limit for the `ulimit` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-ulimit.html#cfn-ecs-taskdefinition-ulimit-softlimit
	//
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

