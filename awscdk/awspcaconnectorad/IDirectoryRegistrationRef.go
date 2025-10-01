package awspcaconnectorad

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryRegistration.
// Experimental.
type IDirectoryRegistrationRef interface {
	constructs.IConstruct
	// A reference to a DirectoryRegistration resource.
	// Experimental.
	DirectoryRegistrationRef() *DirectoryRegistrationReference
}

// The jsii proxy for IDirectoryRegistrationRef
type jsiiProxy_IDirectoryRegistrationRef struct {
	internal.Type__constructsIConstruct
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

