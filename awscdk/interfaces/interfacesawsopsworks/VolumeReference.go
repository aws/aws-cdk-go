package interfacesawsopsworks


// A reference to a Volume resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeReference := &VolumeReference{
//   	VolumeId: jsii.String("volumeId"),
//   }
//
type VolumeReference struct {
	// The Id of the Volume resource.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

