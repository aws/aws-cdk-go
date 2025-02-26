package awsrum


// Creates or updates a destination to receive extended metrics from CloudWatch RUM.
//
// You can send extended metrics to CloudWatch or to a CloudWatch Evidently experiment.
//
// For more information about extended metrics, see [Extended metrics that you can send to CloudWatch and CloudWatch Evidently](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDestinationProperty := &MetricDestinationProperty{
//   	Destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	DestinationArn: jsii.String("destinationArn"),
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	MetricDefinitions: []interface{}{
//   		&MetricDefinitionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DimensionKeys: map[string]*string{
//   				"dimensionKeysKey": jsii.String("dimensionKeys"),
//   			},
//   			EventPattern: jsii.String("eventPattern"),
//   			Namespace: jsii.String("namespace"),
//   			UnitLabel: jsii.String("unitLabel"),
//   			ValueKey: jsii.String("valueKey"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-metricdestination.html
//
type CfnAppMonitor_MetricDestinationProperty struct {
	// Defines the destination to send the metrics to.
	//
	// Valid values are `CloudWatch` and `Evidently` . If you specify `Evidently` , you must also specify the ARN of the CloudWatch Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-metricdestination.html#cfn-rum-appmonitor-metricdestination-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Use this parameter only if `Destination` is `Evidently` .
	//
	// This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-metricdestination.html#cfn-rum-appmonitor-metricdestination-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// This parameter is required if `Destination` is `Evidently` . If `Destination` is `CloudWatch` , do not use this parameter.
	//
	// This parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-metricdestination.html#cfn-rum-appmonitor-metricdestination-iamrolearn
	//
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// An array of structures which define the metrics that you want to send.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-metricdestination.html#cfn-rum-appmonitor-metricdestination-metricdefinitions
	//
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
}

