package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for defining access for a Docker Credential composed of ECR repos.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   ecrDockerCredentialOptions := &ecrDockerCredentialOptions{
//   	assumeRole: role,
//   	usages: []dockerCredentialUsage{
//   		awscdk.Pipelines.*dockerCredentialUsage_SYNTH,
//   	},
//   }
//
type EcrDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	Usages *[]DockerCredentialUsage `field:"optional" json:"usages" yaml:"usages"`
}

