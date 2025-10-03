package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FileSystem.
// Experimental.
type IFileSystemRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFileSystemRef
type jsiiProxy_IFileSystemRef struct {
	internal.Type__constructsIConstruct
}

