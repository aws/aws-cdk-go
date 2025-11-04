package awsinspectorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeSecurityScanConfiguration.
// Experimental.
type ICodeSecurityScanConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CodeSecurityScanConfiguration resource.
	// Experimental.
	CodeSecurityScanConfigurationRef() *CodeSecurityScanConfigurationReference
}

// The jsii proxy for ICodeSecurityScanConfigurationRef
type jsiiProxy_ICodeSecurityScanConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICodeSecurityScanConfigurationRef) CodeSecurityScanConfigurationRef() *CodeSecurityScanConfigurationReference {
	var returns *CodeSecurityScanConfigurationReference
	_jsii_.Get(
		j,
		"codeSecurityScanConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSecurityScanConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSecurityScanConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

