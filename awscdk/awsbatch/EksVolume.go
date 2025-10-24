package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Volume that can be mounted to a container supported by EKS.
//
// Example:
//   jobDefn := batch.NewEksJobDefinition(this, jsii.String("eksf2"), &EksJobDefinitionProps{
//   	Container: batch.NewEksContainerDefinition(this, jsii.String("container"), &EksContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Volumes: []EksVolume{
//   			batch.EksVolume_EmptyDir(&EmptyDirVolumeOptions{
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
type EksVolume interface {
	// The path on the container where the container is mounted.
	// Default: - the container is not mounted.
	//
	ContainerPath() *string
	// The name of this volume.
	//
	// The name must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	Name() *string
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Default: false.
	//
	Readonly() *bool
}

// The jsii proxy struct for EksVolume
type jsiiProxy_EksVolume struct {
	_ byte // padding
}

func (j *jsiiProxy_EksVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


func NewEksVolume_Override(e EksVolume, options *EksVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EksVolume",
		[]interface{}{options},
		e,
	)
}

// Creates a Kubernetes EmptyDir volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
func EksVolume_EmptyDir(options *EmptyDirVolumeOptions) EmptyDirVolume {
	_init_.Initialize()

	if err := validateEksVolume_EmptyDirParameters(options); err != nil {
		panic(err)
	}
	var returns EmptyDirVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EksVolume",
		"emptyDir",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Kubernetes HostPath volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
func EksVolume_HostPath(options *HostPathVolumeOptions) HostPathVolume {
	_init_.Initialize()

	if err := validateEksVolume_HostPathParameters(options); err != nil {
		panic(err)
	}
	var returns HostPathVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EksVolume",
		"hostPath",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Kubernetes Secret volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
func EksVolume_Secret(options *SecretPathVolumeOptions) SecretPathVolume {
	_init_.Initialize()

	if err := validateEksVolume_SecretParameters(options); err != nil {
		panic(err)
	}
	var returns SecretPathVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EksVolume",
		"secret",
		[]interface{}{options},
		&returns,
	)

	return returns
}

