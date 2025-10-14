package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedConfigurationVersion.
// Experimental.
type IHostedConfigurationVersionRef interface {
	constructs.IConstruct
	// A reference to a HostedConfigurationVersion resource.
	// Experimental.
	HostedConfigurationVersionRef() *HostedConfigurationVersionReference
}

// The jsii proxy for IHostedConfigurationVersionRef
type jsiiProxy_IHostedConfigurationVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHostedConfigurationVersionRef) HostedConfigurationVersionRef() *HostedConfigurationVersionReference {
	var returns *HostedConfigurationVersionReference
	_jsii_.Get(
		j,
		"hostedConfigurationVersionRef",
		&returns,
	)
	return returns
}

