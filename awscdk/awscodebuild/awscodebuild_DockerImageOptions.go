package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// The options when creating a CodeBuild Docker build image using {@link LinuxBuildImage.fromDockerRegistry} or {@link WindowsBuildImage.fromDockerRegistry}.
//
// Example:
//   environment: &buildEnvironment{
//   	buildImage: codebuild.linuxBuildImage.fromDockerRegistry(jsii.String("my-registry/my-repo"), &dockerImageOptions{
//   		secretsManagerCredentials: secrets,
//   	}),
//   },
//
// Experimental.
type DockerImageOptions struct {
	// The credentials, stored in Secrets Manager, used for accessing the repository holding the image, if the repository is private.
	// Experimental.
	SecretsManagerCredentials awssecretsmanager.ISecret `field:"optional" json:"secretsManagerCredentials" yaml:"secretsManagerCredentials"`
}

