package awsce


// Properties for defining a `CfnAnomalyMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyMonitorProps := &CfnAnomalyMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//   	MonitorType: jsii.String("monitorType"),
//
//   	// the properties below are optional
//   	MonitorDimension: jsii.String("monitorDimension"),
//   	MonitorSpecification: jsii.String("monitorSpecification"),
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html
//
type CfnAnomalyMonitorProps struct {
	// The name of the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorname
	//
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The type of the monitor.
	//
	// Set this to `DIMENSIONAL` for an AWS managed monitor. AWS managed monitors automatically track up to the top 5,000 values by cost within a dimension of your choosing. Each dimension value is evaluated independently. If you start incurring cost in a new value of your chosen dimension, it will automatically be analyzed by an AWS managed monitor.
	//
	// Set this to `CUSTOM` for a customer managed monitor. Customer managed monitors let you select specific dimension values that get monitored in aggregate.
	//
	// For more information about monitor types, see [Monitor types](https://docs.aws.amazon.com/cost-management/latest/userguide/getting-started-ad.html#monitor-type-def) in the *Billing and Cost Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitortype
	//
	MonitorType *string `field:"required" json:"monitorType" yaml:"monitorType"`
	// For customer managed monitors, do not specify this field.
	//
	// For AWS managed monitors, this field controls which cost dimension is automatically analyzed by the monitor. For `TAG` and `COST_CATEGORY` dimensions, you must also specify MonitorSpecification to configure the specific tag or cost category key to analyze.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitordimension
	//
	MonitorDimension *string `field:"optional" json:"monitorDimension" yaml:"monitorDimension"`
	// The array of `MonitorSpecification` in JSON array format.
	//
	// For instance, you can use `MonitorSpecification` to specify a tag, Cost Category, or linked account for your custom anomaly monitor. For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#aws-resource-ce-anomalymonitor--examples) section of this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorspecification
	//
	MonitorSpecification *string `field:"optional" json:"monitorSpecification" yaml:"monitorSpecification"`
	// Tags to assign to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-resourcetags
	//
	ResourceTags *[]*CfnAnomalyMonitor_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

