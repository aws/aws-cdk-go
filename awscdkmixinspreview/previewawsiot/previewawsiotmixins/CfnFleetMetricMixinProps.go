package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFleetMetricPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetMetricMixinProps := &CfnFleetMetricMixinProps{
//   	AggregationField: jsii.String("aggregationField"),
//   	AggregationType: &AggregationTypeProperty{
//   		Name: jsii.String("name"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IndexName: jsii.String("indexName"),
//   	MetricName: jsii.String("metricName"),
//   	Period: jsii.Number(123),
//   	QueryString: jsii.String("queryString"),
//   	QueryVersion: jsii.String("queryVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html
//
type CfnFleetMetricMixinProps struct {
	// The field to aggregate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-aggregationfield
	//
	AggregationField *string `field:"optional" json:"aggregationField" yaml:"aggregationField"`
	// The type of the aggregation query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-aggregationtype
	//
	AggregationType interface{} `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// The fleet metric description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the index to search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The name of the fleet metric to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The time in seconds between fleet metric emissions.
	//
	// Range [60(1 min), 86400(1 day)] and must be multiple of 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The search query string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The query version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-queryversion
	//
	QueryVersion *string `field:"optional" json:"queryVersion" yaml:"queryVersion"`
	// Metadata which can be used to manage the fleet metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Used to support unit transformation such as milliseconds to seconds.
	//
	// Must be a unit supported by CW metric. Default to null.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-fleetmetric.html#cfn-iot-fleetmetric-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

