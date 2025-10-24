package awsemr


// Properties for defining a `CfnInstanceGroupConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ ConfigurationProperty
//
//   cfnInstanceGroupConfigProps := &CfnInstanceGroupConfigProps{
//   	InstanceCount: jsii.Number(123),
//   	InstanceRole: jsii.String("instanceRole"),
//   	InstanceType: jsii.String("instanceType"),
//   	JobFlowId: jsii.String("jobFlowId"),
//
//   	// the properties below are optional
//   	AutoScalingPolicy: &AutoScalingPolicyProperty{
//   		Constraints: &ScalingConstraintsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   		},
//   		Rules: []interface{}{
//   			&ScalingRuleProperty{
//   				Action: &ScalingActionProperty{
//   					SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   						ScalingAdjustment: jsii.Number(123),
//
//   						// the properties below are optional
//   						AdjustmentType: jsii.String("adjustmentType"),
//   						CoolDown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					Market: jsii.String("market"),
//   				},
//   				Name: jsii.String("name"),
//   				Trigger: &ScalingTriggerProperty{
//   					CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   						ComparisonOperator: jsii.String("comparisonOperator"),
//   						MetricName: jsii.String("metricName"),
//   						Period: jsii.Number(123),
//   						Threshold: jsii.Number(123),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EvaluationPeriods: jsii.Number(123),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: jsii.String("statistic"),
//   						Unit: jsii.String("unit"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	BidPrice: jsii.String("bidPrice"),
//   	Configurations: []interface{}{
//   		&ConfigurationProperty{
//   			Classification: jsii.String("classification"),
//   			ConfigurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			Configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	CustomAmiId: jsii.String("customAmiId"),
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []interface{}{
//   			&EbsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					SizeInGb: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//
//   					// the properties below are optional
//   					Iops: jsii.Number(123),
//   					Throughput: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	Market: jsii.String("market"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html
//
type CfnInstanceGroupConfigProps struct {
	// Target number of instances for the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-instancecount
	//
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The role of the instance group in the cluster.
	//
	// *Allowed Values* : TASK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-instancerole
	//
	InstanceRole *string `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// The Amazon EC2 instance type for all instances in the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The ID of an Amazon EMR cluster that you want to associate this instance group with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-jobflowid
	//
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
	// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
	//
	// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-autoscalingpolicy
	//
	AutoScalingPolicy interface{} `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// If specified, indicates that the instance group uses Spot Instances.
	//
	// This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-bidprice
	//
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// > Amazon EMR releases 4.x or later.
	//
	// The list of configurations supplied for an Amazon EMR cluster instance group. You can specify a separate configuration for each instance group (master, core, and task).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the provisioned instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-customamiid
	//
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-ebsconfiguration
	//
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// Market type of the Amazon EC2 instances used to create a cluster node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-market
	//
	Market *string `field:"optional" json:"market" yaml:"market"`
	// Friendly name given to the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-instancegroupconfig.html#cfn-emr-instancegroupconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

