package awsconnect


// A reference to a QuickConnect resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   quickConnectReference := &QuickConnectReference{
//   	QuickConnectArn: jsii.String("quickConnectArn"),
//   }
//
type QuickConnectReference struct {
	// The QuickConnectArn of the QuickConnect resource.
	QuickConnectArn *string `field:"required" json:"quickConnectArn" yaml:"quickConnectArn"`
}

