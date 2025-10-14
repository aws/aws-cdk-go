package awscodestarconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SyncConfiguration.
// Experimental.
type ISyncConfigurationRef interface {
	constructs.IConstruct
	// A reference to a SyncConfiguration resource.
	// Experimental.
	SyncConfigurationRef() *SyncConfigurationReference
}

// The jsii proxy for ISyncConfigurationRef
type jsiiProxy_ISyncConfigurationRef struct {
	internal.Type__constructsIConstruct
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

