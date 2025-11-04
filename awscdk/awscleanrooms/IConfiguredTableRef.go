package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfiguredTable.
// Experimental.
type IConfiguredTableRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConfiguredTable resource.
	// Experimental.
	ConfiguredTableRef() *ConfiguredTableReference
}

// The jsii proxy for IConfiguredTableRef
type jsiiProxy_IConfiguredTableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConfiguredTableRef) ConfiguredTableRef() *ConfiguredTableReference {
	var returns *ConfiguredTableReference
	_jsii_.Get(
		j,
		"configuredTableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguredTableRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguredTableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

