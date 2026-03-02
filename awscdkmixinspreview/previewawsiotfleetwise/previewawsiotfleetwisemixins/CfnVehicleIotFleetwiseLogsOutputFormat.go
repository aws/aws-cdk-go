package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnVehicleIotFleetwiseLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVehicleIotFleetwiseLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnVehicleIotFleetwiseLogsOutputFormat()
//
// Experimental.
type CfnVehicleIotFleetwiseLogsOutputFormat interface {
}

// The jsii proxy struct for CfnVehicleIotFleetwiseLogsOutputFormat
type jsiiProxy_CfnVehicleIotFleetwiseLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnVehicleIotFleetwiseLogsOutputFormat() CfnVehicleIotFleetwiseLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnVehicleIotFleetwiseLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVehicleIotFleetwiseLogsOutputFormat_Override(c CfnVehicleIotFleetwiseLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

