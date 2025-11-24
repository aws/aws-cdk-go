package mixinsawsce


// Properties for CfnAnomalyMonitorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyMonitorMixinProps := &CfnAnomalyMonitorMixinProps{
//   	MonitorDimension: jsii.String("monitorDimension"),
//   	MonitorName: jsii.String("monitorName"),
//   	MonitorSpecification: jsii.String("monitorSpecification"),
//   	MonitorType: jsii.String("monitorType"),
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
type CfnAnomalyMonitorMixinProps struct {
	// The dimensions to evaluate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitordimension
	//
	MonitorDimension *string `field:"optional" json:"monitorDimension" yaml:"monitorDimension"`
	// The name of the monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorname
	//
	MonitorName *string `field:"optional" json:"monitorName" yaml:"monitorName"`
	// The array of `MonitorSpecification` in JSON array format.
	//
	// For instance, you can use `MonitorSpecification` to specify a tag, Cost Category, or linked account for your custom anomaly monitor. For further information, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#aws-resource-ce-anomalymonitor--examples) section of this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitorspecification
	//
	MonitorSpecification *string `field:"optional" json:"monitorSpecification" yaml:"monitorSpecification"`
	// The possible type values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-monitortype
	//
	MonitorType *string `field:"optional" json:"monitorType" yaml:"monitorType"`
	// Tags to assign to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html#cfn-ce-anomalymonitor-resourcetags
	//
	ResourceTags *[]*CfnAnomalyMonitorPropsMixin_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

