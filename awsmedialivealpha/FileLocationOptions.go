package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// Options for a URL-based file location (`FileLocation.url`).
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//   var passwordParam StringParameter
//
//
//   // From an S3 bucket — the channel role is granted read access automatically
//   fromS3 := medialive.FileLocation_FromBucket(bucket, jsii.String("assets/slate.png"))
//
//   // From a URL with optional credentials (SSM parameter read access auto-granted)
//   fromUrl := medialive.FileLocation_Url(jsii.String("https://origin.example.com/font.ttf"), &FileLocationOptions{
//   	Username: jsii.String("ingest-user"),
//   	Password: passwordParam,
//   })
//
// Experimental.
type FileLocationOptions struct {
	// The SSM parameter that holds the password for accessing the upstream system.
	//
	// The channel
	// role is granted read access to the parameter automatically.
	// Default: - no credentials.
	//
	// Experimental.
	Password awsssm.IStringParameter `field:"optional" json:"password" yaml:"password"`
	// The username for accessing the upstream system.
	// Default: - no credentials.
	//
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

