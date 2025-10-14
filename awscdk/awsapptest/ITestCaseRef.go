package awsapptest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapptest/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestCase.
// Experimental.
type ITestCaseRef interface {
	constructs.IConstruct
	// A reference to a TestCase resource.
	// Experimental.
	TestCaseRef() *TestCaseReference
}

// The jsii proxy for ITestCaseRef
type jsiiProxy_ITestCaseRef struct {
	internal.Type__constructsIConstruct
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

