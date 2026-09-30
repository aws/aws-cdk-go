package awsobservabilityadmin


// Configuration parameters for Amazon MSK cluster monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mskMonitoringParametersProperty := &MskMonitoringParametersProperty{
//   	EnhancedMonitoring: jsii.String("enhancedMonitoring"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-mskmonitoringparameters.html
//
type CfnOrganizationTelemetryRule_MskMonitoringParametersProperty struct {
	// The level of enhanced monitoring for the MSK cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-mskmonitoringparameters.html#cfn-observabilityadmin-organizationtelemetryrule-mskmonitoringparameters-enhancedmonitoring
	//
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
}

