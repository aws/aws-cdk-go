package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomMetric.
// Experimental.
type ICustomMetricRef interface {
	constructs.IConstruct
	// A reference to a CustomMetric resource.
	// Experimental.
	CustomMetricRef() *CustomMetricReference
}

// The jsii proxy for ICustomMetricRef
type jsiiProxy_ICustomMetricRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomMetricRef) CustomMetricRef() *CustomMetricReference {
	var returns *CustomMetricReference
	_jsii_.Get(
		j,
		"customMetricRef",
		&returns,
	)
	return returns
}

