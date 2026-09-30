package interfacesawsstoragegateway


// A reference to a Volume resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeReference := &VolumeReference{
//   	VolumeArn: jsii.String("volumeArn"),
//   }
//
type VolumeReference struct {
	// The VolumeARN of the Volume resource.
	VolumeArn *string `field:"required" json:"volumeArn" yaml:"volumeArn"`
}

