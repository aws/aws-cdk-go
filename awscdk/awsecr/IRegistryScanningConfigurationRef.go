package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryScanningConfiguration.
// Experimental.
type IRegistryScanningConfigurationRef interface {
	constructs.IConstruct
	// A reference to a RegistryScanningConfiguration resource.
	// Experimental.
	RegistryScanningConfigurationRef() *RegistryScanningConfigurationReference
}

// The jsii proxy for IRegistryScanningConfigurationRef
type jsiiProxy_IRegistryScanningConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRegistryScanningConfigurationRef) RegistryScanningConfigurationRef() *RegistryScanningConfigurationReference {
	var returns *RegistryScanningConfigurationReference
	_jsii_.Get(
		j,
		"registryScanningConfigurationRef",
		&returns,
	)
	return returns
}

