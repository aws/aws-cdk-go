package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for HLS static key encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue SecretValue
//
//   hlsStaticKeyProps := &HlsStaticKeyProps{
//   	KeyProviderServerUrl: jsii.String("keyProviderServerUrl"),
//   	StaticKeyValue: secretValue,
//   }
//
// Experimental.
type HlsStaticKeyProps struct {
	// The URL of the license server that serves the static key.
	//
	// Required — MediaLive rejects this without a server URL, even though the underlying
	// CloudFormation property is typed as optional.
	// Experimental.
	KeyProviderServerUrl *string `field:"required" json:"keyProviderServerUrl" yaml:"keyProviderServerUrl"`
	// The static key value as a 32-character hexadecimal string.
	// Experimental.
	StaticKeyValue awscdk.SecretValue `field:"required" json:"staticKeyValue" yaml:"staticKeyValue"`
}

