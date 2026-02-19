package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Delivers vended logs to a CfnDeliveryDestination specified by an arn.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryDestinationRef IDeliveryDestinationRef
//
//   destinationLogsDelivery := awscdkmixinspreview.Aws_logs.NewDestinationLogsDelivery(deliveryDestinationRef)
//
// Experimental.
type DestinationLogsDelivery interface {
	ILogsDelivery
	// Binds Delivery Destination to a source resource for the purposes of log delivery and creates a delivery source and a connection between the source and the destination.
	// Experimental.
	Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy struct for DestinationLogsDelivery
type jsiiProxy_DestinationLogsDelivery struct {
	jsiiProxy_ILogsDelivery
}

// Experimental.
func NewDestinationLogsDelivery(destination interfacesawslogs.IDeliveryDestinationRef) DestinationLogsDelivery {
	_init_.Initialize()

	if err := validateNewDestinationLogsDeliveryParameters(destination); err != nil {
		panic(err)
	}
	j := jsiiProxy_DestinationLogsDelivery{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.DestinationLogsDelivery",
		[]interface{}{destination},
		&j,
	)

	return &j
}

// Experimental.
func NewDestinationLogsDelivery_Override(d DestinationLogsDelivery, destination interfacesawslogs.IDeliveryDestinationRef) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.DestinationLogsDelivery",
		[]interface{}{destination},
		d,
	)
}

func (d *jsiiProxy_DestinationLogsDelivery) Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := d.validateBindParameters(scope, logType, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, logType, sourceResourceArn},
		&returns,
	)

	return returns
}

