package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Delivers vended logs to a CloudWatch Log Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroupRef ILogGroupRef
//
//   logGroupLogsDelivery := awscdkmixinspreview.Aws_logs.NewLogGroupLogsDelivery(logGroupRef)
//
// Experimental.
type LogGroupLogsDelivery interface {
	ILogsDelivery
	// Binds the log group destination to a delivery source and creates a delivery connection between them.
	// Experimental.
	Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy struct for LogGroupLogsDelivery
type jsiiProxy_LogGroupLogsDelivery struct {
	jsiiProxy_ILogsDelivery
}

// Creates a new log group delivery.
// Experimental.
func NewLogGroupLogsDelivery(logGroup interfacesawslogs.ILogGroupRef) LogGroupLogsDelivery {
	_init_.Initialize()

	if err := validateNewLogGroupLogsDeliveryParameters(logGroup); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogGroupLogsDelivery{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.LogGroupLogsDelivery",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

// Creates a new log group delivery.
// Experimental.
func NewLogGroupLogsDelivery_Override(l LogGroupLogsDelivery, logGroup interfacesawslogs.ILogGroupRef) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.LogGroupLogsDelivery",
		[]interface{}{logGroup},
		l,
	)
}

func (l *jsiiProxy_LogGroupLogsDelivery) Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := l.validateBindParameters(scope, deliverySource, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, deliverySource, sourceResourceArn},
		&returns,
	)

	return returns
}

