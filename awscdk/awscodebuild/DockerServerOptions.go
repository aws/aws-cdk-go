package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The Docker server configuration CodeBuild use to build your Docker image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   dockerServerOptions := &DockerServerOptions{
//   	ComputeType: awscdk.Aws_codebuild.DockerServerComputeType_SMALL,
//
//   	// the properties below are optional
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type DockerServerOptions struct {
	// The type of compute to use for the docker server.
	//
	// See the `DockerServerComputeType` enum for the possible values.
	ComputeType DockerServerComputeType `field:"required" json:"computeType" yaml:"computeType"`
	// A list of maximum 5 security groups.
	// Default: - no security group.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

