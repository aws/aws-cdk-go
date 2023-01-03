package awsrum


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionProperty := &metricDefinitionProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dimensionKeys: map[string]*string{
//   		"dimensionKeysKey": jsii.String("dimensionKeys"),
//   	},
//   	eventPattern: jsii.String("eventPattern"),
//   	unitLabel: jsii.String("unitLabel"),
//   	valueKey: jsii.String("valueKey"),
//   }
//
type CfnAppMonitor_MetricDefinitionProperty struct {
	// `CfnAppMonitor.MetricDefinitionProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnAppMonitor.MetricDefinitionProperty.DimensionKeys`.
	DimensionKeys interface{} `field:"optional" json:"dimensionKeys" yaml:"dimensionKeys"`
	// `CfnAppMonitor.MetricDefinitionProperty.EventPattern`.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// `CfnAppMonitor.MetricDefinitionProperty.UnitLabel`.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// `CfnAppMonitor.MetricDefinitionProperty.ValueKey`.
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

