package awsses


// Contains information associated with an Amazon CloudWatch event destination to which email sending events are published.
//
// Event destinations, such as Amazon CloudWatch, are associated with configuration sets, which enable you to publish email sending events. For information about using configuration sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html) .
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
	// A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

