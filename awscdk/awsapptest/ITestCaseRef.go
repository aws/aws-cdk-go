package awsapptest

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapptest/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestCase.
// Experimental.
type ITestCaseRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITestCaseRef
type jsiiProxy_ITestCaseRef struct {
	internal.Type__constructsIConstruct
}

