package awspcaconnectorad

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryRegistration.
// Experimental.
type IDirectoryRegistrationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DirectoryRegistration resource.
	// Experimental.
	DirectoryRegistrationRef() *DirectoryRegistrationReference
}

// The jsii proxy for IDirectoryRegistrationRef
type jsiiProxy_IDirectoryRegistrationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDirectoryRegistrationRef) DirectoryRegistrationRef() *DirectoryRegistrationReference {
	var returns *DirectoryRegistrationReference
	_jsii_.Get(
		j,
		"directoryRegistrationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectoryRegistrationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectoryRegistrationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

