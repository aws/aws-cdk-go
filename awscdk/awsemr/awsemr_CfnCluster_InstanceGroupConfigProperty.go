package awsemr


// Use `InstanceGroupConfig` to define instance groups for an EMR cluster.
//
// A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   instanceGroupConfigProperty := &instanceGroupConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	autoScalingPolicy: &autoScalingPolicyProperty{
//   		constraints: &scalingConstraintsProperty{
//   			maxCapacity: jsii.Number(123),
//   			minCapacity: jsii.Number(123),
//   		},
//   		rules: []interface{}{
//   			&scalingRuleProperty{
//   				action: &scalingActionProperty{
//   					simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   						scalingAdjustment: jsii.Number(123),
//
//   						// the properties below are optional
//   						adjustmentType: jsii.String("adjustmentType"),
//   						coolDown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					market: jsii.String("market"),
//   				},
//   				name: jsii.String("name"),
//   				trigger: &scalingTriggerProperty{
//   					cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   						comparisonOperator: jsii.String("comparisonOperator"),
//   						metricName: jsii.String("metricName"),
//   						period: jsii.Number(123),
//   						threshold: jsii.Number(123),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&metricDimensionProperty{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						evaluationPeriods: jsii.Number(123),
//   						namespace: jsii.String("namespace"),
//   						statistic: jsii.String("statistic"),
//   						unit: jsii.String("unit"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	bidPrice: jsii.String("bidPrice"),
//   	configurations: []interface{}{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	customAmiId: jsii.String("customAmiId"),
//   	ebsConfiguration: &ebsConfigurationProperty{
//   		ebsBlockDeviceConfigs: []interface{}{
//   			&ebsBlockDeviceConfigProperty{
//   				volumeSpecification: &volumeSpecificationProperty{
//   					sizeInGb: jsii.Number(123),
//   					volumeType: jsii.String("volumeType"),
//
//   					// the properties below are optional
//   					iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				volumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		ebsOptimized: jsii.Boolean(false),
//   	},
//   	market: jsii.String("market"),
//   	name: jsii.String("name"),
//   }
//
type CfnCluster_InstanceGroupConfigProperty struct {
	// Target number of instances for the instance group.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The EC2 instance type for all instances in the instance group.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// `AutoScalingPolicy` is a subproperty of the [InstanceGroupConfig](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html) property type that specifies the constraints and rules of an automatic scaling policy in Amazon EMR . The automatic scaling policy defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. Only core and task instance groups can use automatic scaling policies. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) .
	AutoScalingPolicy interface{} `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// If specified, indicates that the instance group uses Spot Instances.
	//
	// This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// > Amazon EMR releases 4.x or later.
	//
	// The list of configurations supplied for an EMR cluster instance group. You can specify a separate configuration for each instance group (master, core, and task).
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the provisioned instance group.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// EBS configurations that will be attached to each EC2 instance in the instance group.
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// Market type of the EC2 instances used to create a cluster node.
	Market *string `field:"optional" json:"market" yaml:"market"`
	// Friendly name given to the instance group.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

