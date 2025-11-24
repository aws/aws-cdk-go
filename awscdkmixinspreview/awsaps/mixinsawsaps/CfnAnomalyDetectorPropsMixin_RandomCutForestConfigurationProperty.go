package mixinsawsaps


// Configuration for the Random Cut Forest algorithm used for anomaly detection in time-series data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   randomCutForestConfigurationProperty := &RandomCutForestConfigurationProperty{
//   	IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   		Amount: jsii.Number(123),
//   		Ratio: jsii.Number(123),
//   	},
//   	IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   		Amount: jsii.Number(123),
//   		Ratio: jsii.Number(123),
//   	},
//   	Query: jsii.String("query"),
//   	SampleSize: jsii.Number(123),
//   	ShingleSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html
//
type CfnAnomalyDetectorPropsMixin_RandomCutForestConfigurationProperty struct {
	// Configuration for ignoring values that are near expected values from above during anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-ignorenearexpectedfromabove
	//
	IgnoreNearExpectedFromAbove interface{} `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// Configuration for ignoring values that are near expected values from below during anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-ignorenearexpectedfrombelow
	//
	IgnoreNearExpectedFromBelow interface{} `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// The Prometheus query used to retrieve the time-series data for anomaly detection.
	//
	// > Random Cut Forest queries must be wrapped by a supported PromQL aggregation operator. For more information, see [Aggregation operators](https://docs.aws.amazon.com/https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators) on the *Prometheus docs* website.
	// >
	// > *Supported PromQL aggregation operators* : `avg` , `count` , `group` , `max` , `min` , `quantile` , `stddev` , `stdvar` , and `sum` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-query
	//
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The number of data points sampled from the input stream for the Random Cut Forest algorithm.
	//
	// The default number is 256 consecutive data points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-samplesize
	//
	// Default: - 256.
	//
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The number of consecutive data points used to create a shingle for the Random Cut Forest algorithm.
	//
	// The default number is 8 consecutive data points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-shinglesize
	//
	// Default: - 8.
	//
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

