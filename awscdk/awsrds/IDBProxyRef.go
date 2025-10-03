package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxy.
// Experimental.
type IDBProxyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBProxyRef
type jsiiProxy_IDBProxyRef struct {
	internal.Type__constructsIConstruct
}

