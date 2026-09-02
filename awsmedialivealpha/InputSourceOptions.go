package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Options for a URL-based input source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stringParameter StringParameter
//
//   inputSourceOptions := &InputSourceOptions{
//   	Password: stringParameter,
//   	Username: jsii.String("username"),
//   }
//
// Experimental.
type InputSourceOptions struct {
	// The SSM parameter that holds the password for accessing the upstream system.
	// Default: - no password.
	//
	// Experimental.
	Password awsssm.IStringParameter `field:"optional" json:"password" yaml:"password"`
	// The username for accessing the upstream system.
	// Default: - no username.
	//
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

