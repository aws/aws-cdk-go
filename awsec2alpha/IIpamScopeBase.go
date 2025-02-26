package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for IpamScope Class.
// Experimental.
type IIpamScopeBase interface {
	// Function to add a new pool to an IPAM scope.
	// Experimental.
	AddPool(id *string, options *PoolOptions) IIpamPool
	// Reference to the current scope of stack to be passed in order to create a new IPAM pool.
	// Experimental.
	Scope() constructs.Construct
	// Default Scope ids created by the IPAM or a new Resource id.
	// Experimental.
	ScopeId() *string
	// Defines scope type can be either default or custom.
	// Experimental.
	ScopeType() IpamScopeType
}

// The jsii proxy for IIpamScopeBase
type jsiiProxy_IIpamScopeBase struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpamScopeBase) AddPool(id *string, options *PoolOptions) IIpamPool {
	if err := i.validateAddPoolParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpamPool

	_jsii_.Invoke(
		i,
		"addPool",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpamScopeBase) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScopeBase) ScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamScopeBase) ScopeType() IpamScopeType {
	var returns IpamScopeType
	_jsii_.Get(
		j,
		"scopeType",
		&returns,
	)
	return returns
}

