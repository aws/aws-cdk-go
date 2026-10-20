package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricConfigurationProperty := &MetricConfigurationProperty{
//   	ExcludedTimeRanges: []interface{}{
//   		&RangeProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	MetricTimeZone: jsii.String("metricTimeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricconfiguration.html
//
type CfnAnomalyDetectorPropsMixin_MetricConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricconfiguration.html#cfn-cloudwatch-anomalydetector-metricconfiguration-excludedtimeranges
	//
	ExcludedTimeRanges interface{} `field:"optional" json:"excludedTimeRanges" yaml:"excludedTimeRanges"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricconfiguration.html#cfn-cloudwatch-anomalydetector-metricconfiguration-metrictimezone
	//
	MetricTimeZone *string `field:"optional" json:"metricTimeZone" yaml:"metricTimeZone"`
}

