package awsobservabilityadmin


// Parameters for log delivery configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logDeliveryParametersProperty := &LogDeliveryParametersProperty{
//   	LogTypes: []*string{
//   		jsii.String("logTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters.html
//
type CfnOrganizationTelemetryRulePropsMixin_LogDeliveryParametersProperty struct {
	// Types of logs to deliver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters.html#cfn-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-logtypes
	//
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

