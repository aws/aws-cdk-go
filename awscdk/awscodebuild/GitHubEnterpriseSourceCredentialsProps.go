package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Creation properties for `GitHubEnterpriseSourceCredentials`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//
//   gitHubEnterpriseSourceCredentialsProps := &GitHubEnterpriseSourceCredentialsProps{
//   	AccessToken: secretValue,
//   }
//
type GitHubEnterpriseSourceCredentialsProps struct {
	// The personal access token to use when contacting the instance of the GitHub Enterprise API.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
}

