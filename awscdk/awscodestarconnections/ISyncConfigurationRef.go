package awscodestarconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SyncConfiguration.
// Experimental.
type ISyncConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SyncConfiguration resource.
	// Experimental.
	SyncConfigurationRef() *SyncConfigurationReference
}

// The jsii proxy for ISyncConfigurationRef
type jsiiProxy_ISyncConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISyncConfigurationRef) SyncConfigurationRef() *SyncConfigurationReference {
	var returns *SyncConfigurationReference
	_jsii_.Get(
		j,
		"syncConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISyncConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISyncConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

