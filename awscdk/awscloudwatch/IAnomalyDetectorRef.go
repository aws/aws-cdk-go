package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyDetector.
// Experimental.
type IAnomalyDetectorRef interface {
	constructs.IConstruct
	// A reference to a AnomalyDetector resource.
	// Experimental.
	AnomalyDetectorRef() *AnomalyDetectorReference
}

// The jsii proxy for IAnomalyDetectorRef
type jsiiProxy_IAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnomalyDetectorRef) AnomalyDetectorRef() *AnomalyDetectorReference {
	var returns *AnomalyDetectorReference
	_jsii_.Get(
		j,
		"anomalyDetectorRef",
		&returns,
	)
	return returns
}

