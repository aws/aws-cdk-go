package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainConfiguration.
// Experimental.
type IDomainConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DomainConfiguration resource.
	// Experimental.
	DomainConfigurationRef() *DomainConfigurationReference
}

// The jsii proxy for IDomainConfigurationRef
type jsiiProxy_IDomainConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDomainConfigurationRef) DomainConfigurationRef() *DomainConfigurationReference {
	var returns *DomainConfigurationReference
	_jsii_.Get(
		j,
		"domainConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

