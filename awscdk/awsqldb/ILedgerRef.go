package awsqldb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqldb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Ledger.
// Experimental.
type ILedgerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILedgerRef
type jsiiProxy_ILedgerRef struct {
	internal.Type__constructsIConstruct
}

