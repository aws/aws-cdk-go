package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageLensGroup.
// Experimental.
type IStorageLensGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStorageLensGroupRef
type jsiiProxy_IStorageLensGroupRef struct {
	internal.Type__constructsIConstruct
}

