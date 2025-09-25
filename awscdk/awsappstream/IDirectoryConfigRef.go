package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryConfig.
// Experimental.
type IDirectoryConfigRef interface {
	constructs.IConstruct
	// A reference to a DirectoryConfig resource.
	// Experimental.
	DirectoryConfigRef() *DirectoryConfigReference
}

// The jsii proxy for IDirectoryConfigRef
type jsiiProxy_IDirectoryConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDirectoryConfigRef) DirectoryConfigRef() *DirectoryConfigReference {
	var returns *DirectoryConfigReference
	_jsii_.Get(
		j,
		"directoryConfigRef",
		&returns,
	)
	return returns
}

