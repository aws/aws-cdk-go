package awsec2


// Options when constructing UserData for Linux.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxUserDataOptions := &LinuxUserDataOptions{
//   	Shebang: jsii.String("shebang"),
//   }
//
type LinuxUserDataOptions struct {
	// Shebang for the UserData script.
	// Default: "#!/bin/bash"
	//
	Shebang *string `field:"optional" json:"shebang" yaml:"shebang"`
}

