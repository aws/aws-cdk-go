package awsiotcoredeviceadvisor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotcoredeviceadvisor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SuiteDefinition.
// Experimental.
type ISuiteDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SuiteDefinition resource.
	// Experimental.
	SuiteDefinitionRef() *SuiteDefinitionReference
}

// The jsii proxy for ISuiteDefinitionRef
type jsiiProxy_ISuiteDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISuiteDefinitionRef) SuiteDefinitionRef() *SuiteDefinitionReference {
	var returns *SuiteDefinitionReference
	_jsii_.Get(
		j,
		"suiteDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuiteDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuiteDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

