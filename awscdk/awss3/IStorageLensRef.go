package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageLens.
// Experimental.
type IStorageLensRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStorageLensRef
type jsiiProxy_IStorageLensRef struct {
	internal.Type__constructsIConstruct
}

