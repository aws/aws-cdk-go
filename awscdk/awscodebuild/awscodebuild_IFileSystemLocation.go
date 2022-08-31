package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// The interface of a CodeBuild FileSystemLocation.
//
// Implemented by {@link EfsFileSystemLocation}.
// Experimental.
type IFileSystemLocation interface {
	// Called by the project when a file system is added so it can perform binding operations on this file system location.
	// Experimental.
	Bind(scope awscdk.Construct, project IProject) *FileSystemConfig
}

// The jsii proxy for IFileSystemLocation
type jsiiProxy_IFileSystemLocation struct {
	_ byte // padding
}

func (i *jsiiProxy_IFileSystemLocation) Bind(scope awscdk.Construct, project IProject) *FileSystemConfig {
	var returns *FileSystemConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

