package awscdkbatchalpha


// Options to configure an EksVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   eksVolumeOptions := &EksVolumeOptions{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	MountPath: jsii.String("mountPath"),
//   	Readonly: jsii.Boolean(false),
//   }
//
// Experimental.
type EksVolumeOptions struct {
	// The name of this volume.
	//
	// The name must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path on the container where the container is mounted.
	// Experimental.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

