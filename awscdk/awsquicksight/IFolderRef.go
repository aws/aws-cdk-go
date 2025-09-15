package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Folder.
// Experimental.
type IFolderRef interface {
	constructs.IConstruct
	// A reference to a Folder resource.
	// Experimental.
	FolderRef() *FolderReference
}

// The jsii proxy for IFolderRef
type jsiiProxy_IFolderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFolderRef) FolderRef() *FolderReference {
	var returns *FolderReference
	_jsii_.Get(
		j,
		"folderRef",
		&returns,
	)
	return returns
}

