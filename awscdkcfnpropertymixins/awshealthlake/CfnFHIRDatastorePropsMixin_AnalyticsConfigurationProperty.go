package awshealthlake


// The analytics configuration for the Data Store.
//
// Omitted when the Data Store can't carry analytics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   analyticsConfigurationProperty := &AnalyticsConfigurationProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-analyticsconfiguration.html
//
type CfnFHIRDatastorePropsMixin_AnalyticsConfigurationProperty struct {
	// The analytics status for the Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-analyticsconfiguration.html#cfn-healthlake-fhirdatastore-analyticsconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

