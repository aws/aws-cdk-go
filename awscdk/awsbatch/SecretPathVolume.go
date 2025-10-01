package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the configuration of a Kubernetes secret volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretPathVolume := awscdk.Aws_batch.NewSecretPathVolume(&SecretPathVolumeOptions{
//   	Name: jsii.String("name"),
//   	SecretName: jsii.String("secretName"),
//
//   	// the properties below are optional
//   	MountPath: jsii.String("mountPath"),
//   	Optional: jsii.Boolean(false),
//   	Readonly: jsii.Boolean(false),
//   })
//
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
type SecretPathVolume interface {
	EksVolume
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
	// Specifies whether the secret or the secret's keys must be defined.
	// Default: true.
	//
	Optional() *bool
	// If specified, the container has readonly access to the volume.
	//
	// Otherwise, the container has read/write access.
	// Default: false.
	//
	Readonly() *bool
	// The name of the secret.
	//
	// Must be a valid DNS subdomain name.
	// See: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	//
	SecretName() *string
}

// The jsii proxy struct for SecretPathVolume
type jsiiProxy_SecretPathVolume struct {
	jsiiProxy_EksVolume
}

func (j *jsiiProxy_SecretPathVolume) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretPathVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretPathVolume) Optional() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"optional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretPathVolume) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretPathVolume) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}


func NewSecretPathVolume(options *SecretPathVolumeOptions) SecretPathVolume {
	_init_.Initialize()

	if err := validateNewSecretPathVolumeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretPathVolume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewSecretPathVolume_Override(s SecretPathVolume, options *SecretPathVolumeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		[]interface{}{options},
		s,
	)
}

// Creates a Kubernetes EmptyDir volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
//
func SecretPathVolume_EmptyDir(options *EmptyDirVolumeOptions) EmptyDirVolume {
	_init_.Initialize()

	if err := validateSecretPathVolume_EmptyDirParameters(options); err != nil {
		panic(err)
	}
	var returns EmptyDirVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		"emptyDir",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a Kubernetes HostPath volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
//
func SecretPathVolume_HostPath(options *HostPathVolumeOptions) HostPathVolume {
	_init_.Initialize()

	if err := validateSecretPathVolume_HostPathParameters(options); err != nil {
		panic(err)
	}
	var returns HostPathVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		"hostPath",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// returns `true` if `x` is a `SecretPathVolume` and `false` otherwise.
func SecretPathVolume_IsSecretPathVolume(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretPathVolume_IsSecretPathVolumeParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		"isSecretPathVolume",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Creates a Kubernetes Secret volume.
// See: https://kubernetes.io/docs/concepts/storage/volumes/#secret
//
func SecretPathVolume_Secret(options *SecretPathVolumeOptions) SecretPathVolume {
	_init_.Initialize()

	if err := validateSecretPathVolume_SecretParameters(options); err != nil {
		panic(err)
	}
	var returns SecretPathVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.SecretPathVolume",
		"secret",
		[]interface{}{options},
		&returns,
	)

	return returns
}

