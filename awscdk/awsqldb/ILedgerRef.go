package awsqldb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsqldb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Ledger.
// Experimental.
type ILedgerRef interface {
	constructs.IConstruct
	// A reference to a Ledger resource.
	// Experimental.
	LedgerRef() *LedgerReference
}

// The jsii proxy for ILedgerRef
type jsiiProxy_ILedgerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILedgerRef) LedgerRef() *LedgerReference {
	var returns *LedgerReference
	_jsii_.Get(
		j,
		"ledgerRef",
		&returns,
	)
	return returns
}

