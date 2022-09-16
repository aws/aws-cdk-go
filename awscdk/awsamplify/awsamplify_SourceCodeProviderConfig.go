package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Configuration for the source code provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//
//   sourceCodeProviderConfig := &sourceCodeProviderConfig{
//   	repository: jsii.String("repository"),
//
//   	// the properties below are optional
//   	accessToken: secretValue,
//   	oauthToken: secretValue,
//   }
//
// Experimental.
type SourceCodeProviderConfig struct {
	// The repository for the application. Must use the `HTTPS` protocol.
	//
	// For example, `https://github.com/aws/aws-cdk`.
	// Experimental.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Personal Access token for 3rd party source control system for an Amplify App, used to create webhook and read-only deploy key.
	//
	// Token is not stored.
	//
	// Either `accessToken` or `oauthToken` must be specified if `repository`
	// is sepcified.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"optional" json:"accessToken" yaml:"accessToken"`
	// OAuth token for 3rd party source control system for an Amplify App, used to create webhook and read-only deploy key.
	//
	// OAuth token is not stored.
	//
	// Either `accessToken` or `oauthToken` must be specified if `repository`
	// is specified.
	// Experimental.
	OauthToken awscdk.SecretValue `field:"optional" json:"oauthToken" yaml:"oauthToken"`
}

