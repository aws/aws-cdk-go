package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCampaignIotFleetwiseLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCampaignIotFleetwiseLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCampaignIotFleetwiseLogsOutputFormat()
//
// Experimental.
type CfnCampaignIotFleetwiseLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCampaignIotFleetwiseLogsOutputFormat
type jsiiProxy_CfnCampaignIotFleetwiseLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCampaignIotFleetwiseLogsOutputFormat() CfnCampaignIotFleetwiseLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCampaignIotFleetwiseLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCampaignIotFleetwiseLogsOutputFormat_Override(c CfnCampaignIotFleetwiseLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

