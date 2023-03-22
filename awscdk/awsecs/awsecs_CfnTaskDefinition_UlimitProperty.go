package awsecs


// The `ulimit` settings to pass to the container.
//
// Amazon ECS tasks hosted on AWS Fargate use the default resource limit values set by the operating system with the exception of the `nofile` resource limit parameter which AWS Fargate overrides. The `nofile` resource limit sets a restriction on the number of open files that a container can use. The default `nofile` soft limit is `1024` and the default hard limit is `4096` .
//
// You can specify the `ulimit` settings for a container in a task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimitProperty := &ulimitProperty{
//   	hardLimit: jsii.Number(123),
//   	name: jsii.String("name"),
//   	softLimit: jsii.Number(123),
//   }
//
type CfnTaskDefinition_UlimitProperty struct {
	// The hard limit for the `ulimit` type.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The `type` of the `ulimit` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The soft limit for the `ulimit` type.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

