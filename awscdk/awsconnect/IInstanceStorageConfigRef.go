package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceStorageConfig.
// Experimental.
type IInstanceStorageConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInstanceStorageConfigRef
type jsiiProxy_IInstanceStorageConfigRef struct {
	internal.Type__constructsIConstruct
}

