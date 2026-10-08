package awsconnect


// Information about the metric configuration for an evaluation form question.
//
// Use this to associate a business outcome metric with a question.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluationFormMetricConfigurationProperty := &EvaluationFormMetricConfigurationProperty{
//   	MetricName: jsii.String("metricName"),
//   	MetricType: jsii.String("metricType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormMetricConfigurationProperty struct {
	// The name of the metric.
	//
	// Valid values are:
	//   +  ``SALE_SUCCESS`` – Sale success.
	//   +  ``CSAT`` – Customer satisfaction.
	//   +  ``CHURN_PROPENSITY`` – Churn propensity.
	//   +  ``SELF_SERVICE_SUCCESS`` – Self-service success.
	//   +  ``PARTIAL_SELF_SERVICE_SUCCESS`` – Partial self-service success.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html#cfn-connect-evaluationform-evaluationformmetricconfiguration-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The type of metric.
	//
	// Currently, only ``BUSINESS_OUTCOME`` is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmetricconfiguration.html#cfn-connect-evaluationform-evaluationformmetricconfiguration-metrictype
	//
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
}

