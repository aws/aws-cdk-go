package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// Properties to define an ECS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsFargateLaunchTargetOptions := &EcsFargateLaunchTargetOptions{
//   	PlatformVersion: awscdk.Aws_ecs.FargatePlatformVersion_LATEST,
//   }
//
// Experimental.
type EcsFargateLaunchTargetOptions struct {
	// Refers to a specific runtime environment for Fargate task infrastructure.
	//
	// Fargate platform version is a combination of the kernel and container runtime versions.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"required" json:"platformVersion" yaml:"platformVersion"`
}

