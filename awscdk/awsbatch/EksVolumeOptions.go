package awsbatch


// Options to configure an EksVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksVolumeOptions := &EksVolumeOptions{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	MountPath: jsii.String("mountPath"),
//   	Readonly: jsii.Boolean(false),
//   }
//
type EksVolumeOptions struct {
	// The name of this volume.
	//
	// The name must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path on the container where the volume is mounted.
	// Default: - the volume is not mounted.
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Default: false.
	//
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

