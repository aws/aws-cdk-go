package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageProfile.
// Experimental.
type IStorageProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStorageProfileRef
type jsiiProxy_IStorageProfileRef struct {
	internal.Type__constructsIConstruct
}

