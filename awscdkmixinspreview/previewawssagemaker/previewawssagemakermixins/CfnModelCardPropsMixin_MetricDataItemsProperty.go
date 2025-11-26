package previewawssagemakermixins


// Metric data.
//
// The `type` determines the data types that you specify for `value` , `XAxisName` and `YAxisName` . For information about specifying values for metrics, see [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var value interface{}
//
//   metricDataItemsProperty := &MetricDataItemsProperty{
//   	Name: jsii.String("name"),
//   	Notes: jsii.String("notes"),
//   	Type: jsii.String("type"),
//   	Value: value,
//   	XAxisName: []*string{
//   		jsii.String("xAxisName"),
//   	},
//   	YAxisName: []*string{
//   		jsii.String("yAxisName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html
//
type CfnModelCardPropsMixin_MetricDataItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-notes
	//
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-xaxisname
	//
	XAxisName *[]*string `field:"optional" json:"xAxisName" yaml:"xAxisName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-metricdataitems.html#cfn-sagemaker-modelcard-metricdataitems-yaxisname
	//
	YAxisName *[]*string `field:"optional" json:"yAxisName" yaml:"yAxisName"`
}

