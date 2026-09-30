package interfacesawsmediaconnect


// A reference to a FlowMediaStream resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowMediaStreamReference := &FlowMediaStreamReference{
//   	FlowMediaStreamArn: jsii.String("flowMediaStreamArn"),
//   }
//
type FlowMediaStreamReference struct {
	// The Arn of the FlowMediaStream resource.
	FlowMediaStreamArn *string `field:"required" json:"flowMediaStreamArn" yaml:"flowMediaStreamArn"`
}

