package awscloudwatch


// Properties for CfnOTelEnrichmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnOTelEnrichmentMixinProps := &CfnOTelEnrichmentMixinProps{
//   	ExcludeFilters: []interface{}{
//   		&OTelEnrichmentMetricSelectorProperty{
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	IncludeFilters: []interface{}{
//   		&OTelEnrichmentMetricSelectorProperty{
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-otelenrichment.html
//
type CfnOTelEnrichmentMixinProps struct {
	// Removes metrics from the include set.
	//
	// Absent or empty means nothing is removed. Evaluated after IncludeFilters, so ExcludeFilters always wins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-otelenrichment.html#cfn-cloudwatch-otelenrichment-excludefilters
	//
	ExcludeFilters interface{} `field:"optional" json:"excludeFilters" yaml:"excludeFilters"`
	// Scopes enrichment to a subset of the account's telemetry.
	//
	// Absent or empty means all namespaces are in scope. Present means only these are. The service enforces a combined cap of 100 selectors across IncludeFilters and ExcludeFilters, and rejects more than one selector for the same namespace within a direction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-otelenrichment.html#cfn-cloudwatch-otelenrichment-includefilters
	//
	IncludeFilters interface{} `field:"optional" json:"includeFilters" yaml:"includeFilters"`
}

