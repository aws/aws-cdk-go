package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// The options when creating a CodeBuild Docker build image using `LinuxBuildImage.fromDockerRegistry` or `WindowsBuildImage.fromDockerRegistry`.
//
// Example:
//   Environment: &BuildEnvironment{
//   	BuildImage: codebuild.LinuxBuildImage_FromDockerRegistry(jsii.String("my-registry/my-repo"), &DockerImageOptions{
//   		SecretsManagerCredentials: secrets,
//   	}),
//   },
//
type DockerImageOptions struct {
	// The credentials, stored in Secrets Manager, used for accessing the repository holding the image, if the repository is private.
	// Default: no credentials will be used (we assume the repository is public).
	//
	SecretsManagerCredentials awssecretsmanager.ISecret `field:"optional" json:"secretsManagerCredentials" yaml:"secretsManagerCredentials"`
}

