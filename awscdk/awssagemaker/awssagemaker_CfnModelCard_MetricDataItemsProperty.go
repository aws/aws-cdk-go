package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   metricDataItemsProperty := &metricDataItemsProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   	value: value,
//
//   	// the properties below are optional
//   	notes: jsii.String("notes"),
//   	xAxisName: []*string{
//   		jsii.String("xAxisName"),
//   	},
//   	yAxisName: []*string{
//   		jsii.String("yAxisName"),
//   	},
//   }
//
type CfnModelCard_MetricDataItemsProperty struct {
	// `CfnModelCard.MetricDataItemsProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnModelCard.MetricDataItemsProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnModelCard.MetricDataItemsProperty.Value`.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// `CfnModelCard.MetricDataItemsProperty.Notes`.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// `CfnModelCard.MetricDataItemsProperty.XAxisName`.
	XAxisName *[]*string `field:"optional" json:"xAxisName" yaml:"xAxisName"`
	// `CfnModelCard.MetricDataItemsProperty.YAxisName`.
	YAxisName *[]*string `field:"optional" json:"yAxisName" yaml:"yAxisName"`
}

