package interfacesawsivschat


// A reference to a Room resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roomReference := &RoomReference{
//   	RoomArn: jsii.String("roomArn"),
//   }
//
type RoomReference struct {
	// The Arn of the Room resource.
	RoomArn *string `field:"required" json:"roomArn" yaml:"roomArn"`
}

