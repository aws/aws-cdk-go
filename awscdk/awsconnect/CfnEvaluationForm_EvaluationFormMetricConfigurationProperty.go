package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormMetricConfigurationProperty := &EvaluationFormMetricConfigurationProperty{
//   	MetricName: jsii.String("metricName"),
//   	MetricType: jsii.String("metricType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html
//
type CfnEvaluationForm_EvaluationFormMetricConfigurationProperty struct {
	// The name of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html#cfn-connect-evaluationform-evaluationformmetricconfiguration-metricname
	//
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The type of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html#cfn-connect-evaluationform-evaluationformmetricconfiguration-metrictype
	//
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
}

