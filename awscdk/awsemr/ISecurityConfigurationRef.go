package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityConfiguration.
// Experimental.
type ISecurityConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityConfiguration resource.
	// Experimental.
	SecurityConfigurationRef() *SecurityConfigurationReference
}

// The jsii proxy for ISecurityConfigurationRef
type jsiiProxy_ISecurityConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISecurityConfigurationRef) SecurityConfigurationRef() *SecurityConfigurationReference {
	var returns *SecurityConfigurationReference
	_jsii_.Get(
		j,
		"securityConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

