package awspinpointemail


// An object that defines an Amazon CloudWatch destination for email events.
//
// You can use Amazon CloudWatch to monitor and gain insights on your email sending metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchDestinationProperty := &cloudWatchDestinationProperty{
//   	dimensionConfigurations: []interface{}{
//   		&dimensionConfigurationProperty{
//   			defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   			dimensionName: jsii.String("dimensionName"),
//   			dimensionValueSource: jsii.String("dimensionValueSource"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch.
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

