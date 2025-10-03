package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSubnetGroup.
// Experimental.
type IDBSubnetGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBSubnetGroupRef
type jsiiProxy_IDBSubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

