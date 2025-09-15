package awsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TestGridProject.
// Experimental.
type ITestGridProjectRef interface {
	constructs.IConstruct
	// A reference to a TestGridProject resource.
	// Experimental.
	TestGridProjectRef() *TestGridProjectReference
}

// The jsii proxy for ITestGridProjectRef
type jsiiProxy_ITestGridProjectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITestGridProjectRef) TestGridProjectRef() *TestGridProjectReference {
	var returns *TestGridProjectReference
	_jsii_.Get(
		j,
		"testGridProjectRef",
		&returns,
	)
	return returns
}

