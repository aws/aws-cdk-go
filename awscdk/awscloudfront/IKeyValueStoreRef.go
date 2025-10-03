package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a KeyValueStore.
// Experimental.
type IKeyValueStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKeyValueStoreRef
type jsiiProxy_IKeyValueStoreRef struct {
	internal.Type__constructsIConstruct
}

