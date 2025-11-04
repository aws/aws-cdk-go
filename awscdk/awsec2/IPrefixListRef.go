package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrefixList.
// Experimental.
type IPrefixListRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PrefixList resource.
	// Experimental.
	PrefixListRef() *PrefixListReference
}

// The jsii proxy for IPrefixListRef
type jsiiProxy_IPrefixListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPrefixListRef) PrefixListRef() *PrefixListReference {
	var returns *PrefixListReference
	_jsii_.Get(
		j,
		"prefixListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixListRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

