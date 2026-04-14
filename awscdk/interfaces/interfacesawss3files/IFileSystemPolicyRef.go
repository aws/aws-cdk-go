package interfacesawss3files

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3files/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FileSystemPolicy.
// Experimental.
type IFileSystemPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FileSystemPolicy resource.
	// Experimental.
	FileSystemPolicyRef() *FileSystemPolicyReference
}

// The jsii proxy for IFileSystemPolicyRef
type jsiiProxy_IFileSystemPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFileSystemPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFileSystemPolicyRef) FileSystemPolicyRef() *FileSystemPolicyReference {
	var returns *FileSystemPolicyReference
	_jsii_.Get(
		j,
		"fileSystemPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileSystemPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFileSystemPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

