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
//   cloudWatchDestinationProperty := &CloudWatchDestinationProperty{
//   	DimensionConfigurations: []interface{}{
//   		&DimensionConfigurationProperty{
//   			DefaultDimensionValue: jsii.String("defaultDimensionValue"),
//   			DimensionName: jsii.String("dimensionName"),
//   			DimensionValueSource: jsii.String("dimensionValueSource"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-cloudwatchdestination.html
//
type CfnConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-cloudwatchdestination.html#cfn-ses-configurationseteventdestination-cloudwatchdestination-dimensionconfigurations
	//
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

