package awssagemaker


// A group of metric data that you use to initialize a metric group object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   metricGroupProperty := &MetricGroupProperty{
//   	MetricData: []interface{}{
//   		&MetricDataItemsProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   			Value: value,
//
//   			// the properties below are optional
//   			Notes: jsii.String("notes"),
//   			XAxisName: []*string{
//   				jsii.String("xAxisName"),
//   			},
//   			YAxisName: []*string{
//   				jsii.String("yAxisName"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricgroup.html
//
type CfnModelCard_MetricGroupProperty struct {
	// A list of metric objects. The `MetricDataItems` list can have one of the following values:.
	//
	// - `bar_chart_metric`
	// - `matrix_metric`
	// - `simple_metric`
	// - `linear_graph_metric`
	//
	// For more information about the metric schema, see the definition section of the [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricgroup.html#cfn-sagemaker-modelcard-metricgroup-metricdata
	//
	MetricData interface{} `field:"required" json:"metricData" yaml:"metricData"`
	// The metric group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricgroup.html#cfn-sagemaker-modelcard-metricgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

