package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for a Kubernetes EmptyDir volume.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDefn := batch.NewEksJobDefinition(this, jsii.String("eksf2"), &EksJobDefinitionProps{
//   	Container: batch.NewEksContainerDefinition(this, jsii.String("container"), &EksContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Volumes: []eksVolume{
//   			batch.*eksVolume_EmptyDir(&EmptyDirVolumeOptions{
//   				Name: jsii.String("myEmptyDirVolume"),
//   				MountPath: jsii.String("/mount/path"),
//   				Medium: batch.EmptyDirMediumType_MEMORY,
//   				Readonly: jsii.Boolean(true),
//   				SizeLimit: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   		},
//   	}),
//   })
//
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
// Experimental.
type EmptyDirVolumeOptions struct {
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
	// The storage type to use for this Volume.
	// Experimental.
	Medium EmptyDirMediumType `field:"optional" json:"medium" yaml:"medium"`
	// The maximum size for this Volume.
	// Experimental.
	SizeLimit awscdk.Size `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

