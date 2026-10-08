package awscloudwatch


// Selects metrics within one namespace.
//
// The same shape serves both the include and the exclude direction. Namespaces are compared byte-for-byte and case-sensitively; no wildcards, prefixes or normalization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oTelEnrichmentMetricSelectorProperty := &OTelEnrichmentMetricSelectorProperty{
//   	MetricNames: []*string{
//   		jsii.String("metricNames"),
//   	},
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-otelenrichment-otelenrichmentmetricselector.html
//
type CfnOTelEnrichmentPropsMixin_OTelEnrichmentMetricSelectorProperty struct {
	// Absent or empty means every metric in this namespace, in whichever direction this selector appears.
	//
	// Present means only these metric names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-otelenrichment-otelenrichmentmetricselector.html#cfn-cloudwatch-otelenrichment-otelenrichmentmetricselector-metricnames
	//
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// The CloudWatch namespace this selector applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-otelenrichment-otelenrichmentmetricselector.html#cfn-cloudwatch-otelenrichment-otelenrichmentmetricselector-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

