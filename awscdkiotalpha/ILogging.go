package awscdkiotalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2/internal"
)

// Represents AWS IoT Logging.
// Experimental.
type ILogging interface {
	awscdk.IResource
	// The log ID.
	// Experimental.
	LogId() *string
}

// The jsii proxy for ILogging
type jsiiProxy_ILogging struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILogging) LogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logId",
		&returns,
	)
	return returns
}

