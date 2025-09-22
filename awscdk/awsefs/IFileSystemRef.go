package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FileSystem.
// Experimental.
type IFileSystemRef interface {
	constructs.IConstruct
	// A reference to a FileSystem resource.
	// Experimental.
	FileSystemRef() *FileSystemReference
}

// The jsii proxy for IFileSystemRef
type jsiiProxy_IFileSystemRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFileSystemRef) FileSystemRef() *FileSystemReference {
	var returns *FileSystemReference
	_jsii_.Get(
		j,
		"fileSystemRef",
		&returns,
	)
	return returns
}

