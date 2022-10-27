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
//   instanceGroupConfigProperty := &instanceGroupConfigProperty{
//   	instanceCount: jsii.Number(123),
//   	instanceRole: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.instanceRoleType_MASTER,
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	autoScalingPolicy: &autoScalingPolicyProperty{
//   		constraints: &scalingConstraintsProperty{
//   			maxCapacity: jsii.Number(123),
//   			minCapacity: jsii.Number(123),
//   		},
//   		rules: []scalingRuleProperty{
//   			&scalingRuleProperty{
//   				action: &scalingActionProperty{
//   					simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   						scalingAdjustment: jsii.Number(123),
//
//   						// the properties below are optional
//   						adjustmentType: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   						coolDown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.instanceMarket_ON_DEMAND,
//   				},
//   				name: jsii.String("name"),
//   				trigger: &scalingTriggerProperty{
//   					cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   						comparisonOperator: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   						metricName: jsii.String("metricName"),
//   						period: cdk.duration.minutes(jsii.Number(30)),
//
//   						// the properties below are optional
//   						dimensions: []metricDimensionProperty{
//   							&metricDimensionProperty{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						evaluationPeriods: jsii.Number(123),
//   						namespace: jsii.String("namespace"),
//   						statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   						threshold: jsii.Number(123),
//   						unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   					},
//   				},
//
//   				// the properties below are optional
//   				description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	bidPrice: jsii.String("bidPrice"),
//   	configurations: []*configurationProperty{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	ebsConfiguration: &ebsConfigurationProperty{
//   		ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   			&ebsBlockDeviceConfigProperty{
//   				volumeSpecification: &volumeSpecificationProperty{
//   					volumeSize: size,
//   					volumeType: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
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
//   	market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.*instanceMarket_ON_DEMAND,
//   	name: jsii.String("name"),
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
	AutoScalingPolicy *EmrCreateCluster_AutoScalingPolicyProperty `field:"optional" json:"autoScalingPolicy" yaml:"autoScalingPolicy"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The list of configurations supplied for an EMR cluster instance group.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// EBS configurations that will be attached to each EC2 instance in the instance group.
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// Market type of the EC2 instances used to create a cluster node.
	Market EmrCreateCluster_InstanceMarket `field:"optional" json:"market" yaml:"market"`
	// Friendly name given to the instance group.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

