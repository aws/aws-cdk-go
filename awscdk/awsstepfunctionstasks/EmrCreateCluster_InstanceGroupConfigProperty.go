package awsstepfunctionstasks


// Configuration defining a new instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceGroupConfigProperty := &InstanceGroupConfigProperty{
//   	InstanceCount: jsii.Number(123),
//   	InstanceRole: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	AutoScalingPolicy: &AutoScalingPolicyProperty{
//   		Constraints: &ScalingConstraintsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   		},
//   		Rules: []scalingRuleProperty{
//   			&scalingRuleProperty{
//   				Action: &ScalingActionProperty{
//   					SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   						ScalingAdjustment: jsii.Number(123),
//
//   						// the properties below are optional
//   						AdjustmentType: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.ScalingAdjustmentType_CHANGE_IN_CAPACITY,
//   						CoolDown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					Market: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.InstanceMarket_ON_DEMAND,
//   				},
//   				Name: jsii.String("name"),
//   				Trigger: &ScalingTriggerProperty{
//   					CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   						ComparisonOperator: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   						MetricName: jsii.String("metricName"),
//   						Period: cdk.Duration_Minutes(jsii.Number(30)),
//
//   						// the properties below are optional
//   						Dimensions: []metricDimensionProperty{
//   							&metricDimensionProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EvaluationPeriods: jsii.Number(123),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmStatistic_SAMPLE_COUNT,
//   						Threshold: jsii.Number(123),
//   						Unit: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmUnit_NONE,
//   					},
//   				},
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	BidPrice: jsii.String("bidPrice"),
//   	Configurations: []*configurationProperty{
//   		&configurationProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   			&ebsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					VolumeSize: size,
//   					VolumeType: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP2,
//
//   					// the properties below are optional
//   					Iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	Market: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.InstanceMarket_ON_DEMAND,
//   	Name: jsii.String("name"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroupConfig.html
//
type EmrCreateCluster_InstanceGroupConfigProperty struct {
	// Target number of instances for the instance group.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The role of the instance group in the cluster.
	InstanceRole EmrCreateCluster_InstanceRoleType `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// The EC2 instance type for all instances in the instance group.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.
	// Default: - None.
	//
	AutoScalingPolicy *EmrCreateCluster_AutoScalingPolicyProperty `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	// Default: - None.
	//
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The list of configurations supplied for an EMR cluster instance group.
	// Default: - None.
	//
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// EBS configurations that will be attached to each EC2 instance in the instance group.
	// Default: - None.
	//
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// Market type of the EC2 instances used to create a cluster node.
	// Default: - EMR selected default.
	//
	Market EmrCreateCluster_InstanceMarket `field:"optional" json:"market" yaml:"market"`
	// Friendly name given to the instance group.
	// Default: - None.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

