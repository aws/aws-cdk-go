package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The interface of a CodeBuild FileSystemLocation.
//
// Implemented by {@link EfsFileSystemLocation}.
type IFileSystemLocation interface {
	// Called by the project when a file system is added so it can perform binding operations on this file system location.
	Bind(scope constructs.Construct, project IProject) *FileSystemConfig
}

// The jsii proxy for IFileSystemLocation
type jsiiProxy_IFileSystemLocation struct {
	_ byte // padding
}

func (i *jsiiProxy_IFileSystemLocation) Bind(scope constructs.Construct, project IProject) *FileSystemConfig {
	if err := i.validateBindParameters(scope, project); err != nil {
		panic(err)
	}
	var returns *FileSystemConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

