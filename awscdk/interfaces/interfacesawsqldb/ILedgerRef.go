package interfacesawsqldb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsqldb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Ledger.
// Experimental.
type ILedgerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Ledger resource.
	// Experimental.
	LedgerRef() *LedgerReference
}

// The jsii proxy for ILedgerRef
type jsiiProxy_ILedgerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ILedgerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILedgerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

