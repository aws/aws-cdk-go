// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Kubernetes HostPath volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   hostPathVolume := batch_alpha.NewHostPathVolume(&HostPathVolumeOptions{
//   	HostPath: jsii.String("hostPath"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	MountPath: jsii.String("mountPath"),
//   	Readonly: jsii.Boolean(false),
//   })
//
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
// Experimental.
type HostPathVolume interface {
	EksVolume
	// The path on the container where the container is mounted.
	// Experimental.
	ContainerPath() *string
	// The name of this volume.
	//
	// The name must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	// Experimental.
	Name() *string
	// The path of the file or directory on the host to mount into containers on the pod.
	//
	// *Note*: HothPath Volumes present many security risks, and should be avoided when possible.
	// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
	//
	// Experimental.
	Path() *string
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Experimental.
	Readonly() *bool
}

// The jsii proxy struct for HostPathVolume
type jsiiProxy_HostPathVolume struct {
	jsiiProxy_EksVolume
}

func (j *jsiiProxy_HostPathVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPathVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPathVolume) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostPathVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewHostPathVolume(options *HostPathVolumeOptions) HostPathVolume {
	_init_.Initialize()

	if err := validateNewHostPathVolumeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostPathVolume{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewHostPathVolume_Override(h HostPathVolume, options *HostPathVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		[]interface{}{options},
		h,
	)
}

// Creates a Kubernetes EmptyDir volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
// Experimental.
func HostPathVolume_EmptyDir(options *EmptyDirVolumeOptions) EmptyDirVolume {
	_init_.Initialize()

	if err := validateHostPathVolume_EmptyDirParameters(options); err != nil {
		panic(err)
	}
	var returns EmptyDirVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		"emptyDir",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Kubernetes HostPath volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
// Experimental.
func HostPathVolume_HostPath(options *HostPathVolumeOptions) HostPathVolume {
	_init_.Initialize()

	if err := validateHostPathVolume_HostPathParameters(options); err != nil {
		panic(err)
	}
	var returns HostPathVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		"hostPath",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// returns `true` if `x` is a HostPathVolume, `false` otherwise.
// Experimental.
func HostPathVolume_IsHostPathVolume(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostPathVolume_IsHostPathVolumeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		"isHostPathVolume",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Creates a Kubernetes Secret volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
// Experimental.
func HostPathVolume_Secret(options *SecretPathVolumeOptions) SecretPathVolume {
	_init_.Initialize()

	if err := validateHostPathVolume_SecretParameters(options); err != nil {
		panic(err)
	}
	var returns SecretPathVolume

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.HostPathVolume",
		"secret",
		[]interface{}{options},
		&returns,
	)

	return returns
}

