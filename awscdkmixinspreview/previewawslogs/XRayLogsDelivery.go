package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Delivers vended logs to AWS X-Ray.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   xRayLogsDelivery := awscdkmixinspreview.Aws_logs.NewXRayLogsDelivery()
//
// Experimental.
type XRayLogsDelivery interface {
	ILogsDelivery
	// Binds X-Ray Destination to a source resource for the purposes of log delivery and creates a delivery source, a delivery destination, and a connection between them.
	// Experimental.
	Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy struct for XRayLogsDelivery
type jsiiProxy_XRayLogsDelivery struct {
	jsiiProxy_ILogsDelivery
}

// Creates a new X-Ray delivery.
// Experimental.
func NewXRayLogsDelivery() XRayLogsDelivery {
	_init_.Initialize()

	j := jsiiProxy_XRayLogsDelivery{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.XRayLogsDelivery",
		nil, // no parameters
		&j,
	)

	return &j
}

// Creates a new X-Ray delivery.
// Experimental.
func NewXRayLogsDelivery_Override(x XRayLogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.XRayLogsDelivery",
		nil, // no parameters
		x,
	)
}

func (x *jsiiProxy_XRayLogsDelivery) Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := x.validateBindParameters(scope, logType, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		x,
		"bind",
		[]interface{}{scope, logType, sourceResourceArn},
		&returns,
	)

	return returns
}

