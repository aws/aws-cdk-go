package previewawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the delivery of vended logs to a destination.
// Experimental.
type ILogsDelivery interface {
	// Binds the log delivery to a source resource and creates a delivery connection between the source and destination.
	//
	// Returns: The delivery reference.
	// Experimental.
	Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy for ILogsDelivery
type jsiiProxy_ILogsDelivery struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogsDelivery) Bind(scope constructs.IConstruct, logType *string, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := i.validateBindParameters(scope, logType, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, logType, sourceResourceArn},
		&returns,
	)

	return returns
}

