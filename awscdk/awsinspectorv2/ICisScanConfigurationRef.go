package awsinspectorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CisScanConfiguration.
// Experimental.
type ICisScanConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CisScanConfiguration resource.
	// Experimental.
	CisScanConfigurationRef() *CisScanConfigurationReference
}

// The jsii proxy for ICisScanConfigurationRef
type jsiiProxy_ICisScanConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICisScanConfigurationRef) CisScanConfigurationRef() *CisScanConfigurationReference {
	var returns *CisScanConfigurationReference
	_jsii_.Get(
		j,
		"cisScanConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICisScanConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICisScanConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

