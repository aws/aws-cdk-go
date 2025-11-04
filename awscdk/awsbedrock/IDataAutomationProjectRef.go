package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAutomationProject.
// Experimental.
type IDataAutomationProjectRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataAutomationProject resource.
	// Experimental.
	DataAutomationProjectRef() *DataAutomationProjectReference
}

// The jsii proxy for IDataAutomationProjectRef
type jsiiProxy_IDataAutomationProjectRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDataAutomationProjectRef) DataAutomationProjectRef() *DataAutomationProjectReference {
	var returns *DataAutomationProjectReference
	_jsii_.Get(
		j,
		"dataAutomationProjectRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAutomationProjectRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAutomationProjectRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

