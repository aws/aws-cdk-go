package awsecs


// The details on a data volume from another container in the same task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeFrom := &volumeFrom{
//   	readOnly: jsii.Boolean(false),
//   	sourceContainer: jsii.String("sourceContainer"),
//   }
//
type VolumeFrom struct {
	// Specifies whether the container has read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	// The name of another container within the same task definition from which to mount volumes.
	SourceContainer *string `field:"required" json:"sourceContainer" yaml:"sourceContainer"`
}

