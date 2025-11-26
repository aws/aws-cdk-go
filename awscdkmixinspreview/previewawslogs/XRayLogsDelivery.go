package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
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
	// Binds the X-Ray destination to a delivery source and creates a delivery connection between them.
	// Experimental.
	Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig
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

func (x *jsiiProxy_XRayLogsDelivery) Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := x.validateBindParameters(scope, deliverySource, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		x,
		"bind",
		[]interface{}{scope, deliverySource, sourceResourceArn},
		&returns,
	)

	return returns
}

