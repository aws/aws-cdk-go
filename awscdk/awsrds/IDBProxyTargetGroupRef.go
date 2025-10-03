package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBProxyTargetGroup.
// Experimental.
type IDBProxyTargetGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBProxyTargetGroupRef
type jsiiProxy_IDBProxyTargetGroupRef struct {
	internal.Type__constructsIConstruct
}

