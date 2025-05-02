package awsbatch


// Options for a Kubernetes SecretPath Volume.
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
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
type SecretPathVolumeOptions struct {
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
	// The name of the secret.
	//
	// Must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Specifies whether the secret or the secret's keys must be defined.
	// Default: true.
	//
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

