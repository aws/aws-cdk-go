package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Options for a URL-based output destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stringParameter StringParameter
//
//   outputDestinationOptions := &OutputDestinationOptions{
//   	Password: stringParameter,
//   	Username: jsii.String("username"),
//   }
//
// Experimental.
type OutputDestinationOptions struct {
	// An SSM String Parameter holding the password for accessing the downstream system.
	//
	// MediaLive
	// reads it from Parameter Store at channel runtime, so the channel role is granted read access.
	// Default: - no credentials.
	//
	// Experimental.
	Password awsssm.IStringParameter `field:"optional" json:"password" yaml:"password"`
	// The username for accessing the downstream system.
	// Default: - no credentials.
	//
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

