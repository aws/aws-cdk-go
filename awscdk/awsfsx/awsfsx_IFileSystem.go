package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsfsx/internal"
)

// Interface to implement FSx File Systems.
// Experimental.
type IFileSystem interface {
	awsec2.IConnectable
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId() *string
}

// The jsii proxy for IFileSystem
type jsiiProxy_IFileSystem struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_IFileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

