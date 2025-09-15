package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RemediationConfiguration.
// Experimental.
type IRemediationConfigurationRef interface {
	constructs.IConstruct
	// A reference to a RemediationConfiguration resource.
	// Experimental.
	RemediationConfigurationRef() *RemediationConfigurationReference
}

// The jsii proxy for IRemediationConfigurationRef
type jsiiProxy_IRemediationConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRemediationConfigurationRef) RemediationConfigurationRef() *RemediationConfigurationReference {
	var returns *RemediationConfigurationReference
	_jsii_.Get(
		j,
		"remediationConfigurationRef",
		&returns,
	)
	return returns
}

