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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-cloudwatchdestination.html
//
type CfnConfigurationSetEventDestination_CloudWatchDestinationProperty struct {
	// An array of objects that define the dimensions to use when you send email events to Amazon CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-cloudwatchdestination.html#cfn-pinpointemail-configurationseteventdestination-cloudwatchdestination-dimensionconfigurations
	//
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

