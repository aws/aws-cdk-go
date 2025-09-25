package awstransfer


// A reference to a Server resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverReference := &ServerReference{
//   	ServerArn: jsii.String("serverArn"),
//   }
//
type ServerReference struct {
	// The Arn of the Server resource.
	ServerArn *string `field:"required" json:"serverArn" yaml:"serverArn"`
}

