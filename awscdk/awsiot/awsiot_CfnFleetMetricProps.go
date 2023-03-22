package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFleetMetric`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetMetricProps := &cfnFleetMetricProps{
//   	metricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	aggregationField: jsii.String("aggregationField"),
//   	aggregationType: &aggregationTypeProperty{
//   		name: jsii.String("name"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	indexName: jsii.String("indexName"),
//   	period: jsii.Number(123),
//   	queryString: jsii.String("queryString"),
//   	queryVersion: jsii.String("queryVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   }
//
type CfnFleetMetricProps struct {
	// The name of the fleet metric to create.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The field to aggregate.
	AggregationField *string `field:"optional" json:"aggregationField" yaml:"aggregationField"`
	// The type of the aggregation query.
	AggregationType interface{} `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// The fleet metric description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the index to search.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The time in seconds between fleet metric emissions.
	//
	// Range [60(1 min), 86400(1 day)] and must be multiple of 60.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The search query string.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The query version.
	QueryVersion *string `field:"optional" json:"queryVersion" yaml:"queryVersion"`
	// Metadata which can be used to manage the fleet metric.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Used to support unit transformation such as milliseconds to seconds.
	//
	// Must be a unit supported by CW metric. Default to null.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

