package mixinsawscloudwatch


// This object includes parameters that you can use to provide information to CloudWatch to help it build more accurate anomaly detection models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricCharacteristicsProperty := &MetricCharacteristicsProperty{
//   	PeriodicSpikes: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metriccharacteristics.html
//
type CfnAnomalyDetectorPropsMixin_MetricCharacteristicsProperty struct {
	// Set this parameter to true if values for this metric consistently include spikes that should not be considered to be anomalies.
	//
	// With this set to true, CloudWatch will expect to see spikes that occurred consistently during the model training period, and won't flag future similar spikes as anomalies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metriccharacteristics.html#cfn-cloudwatch-anomalydetector-metriccharacteristics-periodicspikes
	//
	PeriodicSpikes interface{} `field:"optional" json:"periodicSpikes" yaml:"periodicSpikes"`
}

