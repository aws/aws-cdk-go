package awsapptest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapptest/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestCase.
// Experimental.
type ITestCaseRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TestCase resource.
	// Experimental.
	TestCaseRef() *TestCaseReference
}

// The jsii proxy for ITestCaseRef
type jsiiProxy_ITestCaseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITestCaseRef) TestCaseRef() *TestCaseReference {
	var returns *TestCaseReference
	_jsii_.Get(
		j,
		"testCaseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITestCaseRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITestCaseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

