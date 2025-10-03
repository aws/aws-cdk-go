package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Folder.
// Experimental.
type IFolderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFolderRef
type jsiiProxy_IFolderRef struct {
	internal.Type__constructsIConstruct
}

