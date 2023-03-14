package awsce


// Properties for defining a `CfnAnomalyMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyMonitorProps := &cfnAnomalyMonitorProps{
//   	monitorName: jsii.String("monitorName"),
//   	monitorType: jsii.String("monitorType"),
//
//   	// the properties below are optional
//   	monitorDimension: jsii.String("monitorDimension"),
//   	monitorSpecification: jsii.String("monitorSpecification"),
//   	resourceTags: []interface{}{
//   		&resourceTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAnomalyMonitorProps struct {
	// The name of the monitor.
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The possible type values.
	MonitorType *string `field:"required" json:"monitorType" yaml:"monitorType"`
	// The dimensions to evaluate.
	MonitorDimension *string `field:"optional" json:"monitorDimension" yaml:"monitorDimension"`
	// The array of `MonitorSpecification` in JSON array format.
	//
	// For instance, you can use `MonitorSpecification` to specify a tag, Cost Category, or linked account for your custom anomaly monitor. For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#aws-resource-ce-anomalymonitor--examples) section of this page.
	MonitorSpecification *string `field:"optional" json:"monitorSpecification" yaml:"monitorSpecification"`
	// `AWS::CE::AnomalyMonitor.ResourceTags`.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

