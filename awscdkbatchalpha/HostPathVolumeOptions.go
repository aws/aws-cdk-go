// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Options for a kubernetes HostPath volume.
//
// Example:
//   var jobDefn eksJobDefinition
//
//   jobDefn.Container.AddVolume(batch.EksVolume_EmptyDir(&EmptyDirVolumeOptions{
//   	Name: jsii.String("emptyDir"),
//   	MountPath: jsii.String("/Volumes/emptyDir"),
//   }))
//   jobDefn.Container.AddVolume(batch.EksVolume_HostPath(&HostPathVolumeOptions{
//   	Name: jsii.String("hostPath"),
//   	HostPath: jsii.String("/sys"),
//   	MountPath: jsii.String("/Volumes/hostPath"),
//   }))
//   jobDefn.Container.AddVolume(batch.EksVolume_Secret(&SecretPathVolumeOptions{
//   	Name: jsii.String("secret"),
//   	Optional: jsii.Boolean(true),
//   	MountPath: jsii.String("/Volumes/secret"),
//   	SecretName: jsii.String("mySecret"),
//   }))
//
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
// Experimental.
type HostPathVolumeOptions struct {
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
	// The path of the file or directory on the host to mount into containers on the pod.
	//
	// *Note*: HothPath Volumes present many security risks, and should be avoided when possible.
	// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
	//
	// Experimental.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
}

