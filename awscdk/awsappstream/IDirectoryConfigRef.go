package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectoryConfig.
// Experimental.
type IDirectoryConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDirectoryConfigRef
type jsiiProxy_IDirectoryConfigRef struct {
	internal.Type__constructsIConstruct
}

