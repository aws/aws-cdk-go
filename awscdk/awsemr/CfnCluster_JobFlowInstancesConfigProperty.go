package awsemr


// `JobFlowInstancesConfig` is a property of the `AWS::EMR::Cluster` resource.
//
// `JobFlowInstancesConfig` defines the instance groups or instance fleets that comprise the cluster. `JobFlowInstancesConfig` must contain either `InstanceFleetConfig` or `InstanceGroupConfig` . They cannot be used together.
//
// You can now define task instance groups or task instance fleets using the `TaskInstanceGroups` and `TaskInstanceFleets` subproperties. Using these subproperties reduces delays in provisioning task nodes compared to specifying task nodes with the `InstanceFleetConfig` and `InstanceGroupConfig` resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   jobFlowInstancesConfigProperty := &JobFlowInstancesConfigProperty{
//   	AdditionalMasterSecurityGroups: []*string{
//   		jsii.String("additionalMasterSecurityGroups"),
//   	},
//   	AdditionalSlaveSecurityGroups: []*string{
//   		jsii.String("additionalSlaveSecurityGroups"),
//   	},
//   	CoreInstanceFleet: &InstanceFleetConfigProperty{
//   		InstanceTypeConfigs: []interface{}{
//   			&InstanceTypeConfigProperty{
//   				InstanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				BidPrice: jsii.String("bidPrice"),
//   				BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   				Configurations: []interface{}{
//   					&configurationProperty{
//   						Classification: jsii.String("classification"),
//   						ConfigurationProperties: map[string]*string{
//   							"configurationPropertiesKey": jsii.String("configurationProperties"),
//   						},
//   						Configurations: []interface{}{
//   							configurationProperty_,
//   						},
//   					},
//   				},
//   				CustomAmiId: jsii.String("customAmiId"),
//   				EbsConfiguration: &EbsConfigurationProperty{
//   					EbsBlockDeviceConfigs: []interface{}{
//   						&EbsBlockDeviceConfigProperty{
//   							VolumeSpecification: &VolumeSpecificationProperty{
//   								SizeInGb: jsii.Number(123),
//   								VolumeType: jsii.String("volumeType"),
//
//   								// the properties below are optional
//   								Iops: jsii.Number(123),
//   								Throughput: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							VolumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					EbsOptimized: jsii.Boolean(false),
//   				},
//   				WeightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   			OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   				AllocationStrategy: jsii.String("allocationStrategy"),
//   			},
//   			SpotSpecification: &SpotProvisioningSpecificationProperty{
//   				TimeoutAction: jsii.String("timeoutAction"),
//   				TimeoutDurationMinutes: jsii.Number(123),
//
//   				// the properties below are optional
//   				AllocationStrategy: jsii.String("allocationStrategy"),
//   				BlockDurationMinutes: jsii.Number(123),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		TargetOnDemandCapacity: jsii.Number(123),
//   		TargetSpotCapacity: jsii.Number(123),
//   	},
//   	CoreInstanceGroup: &InstanceGroupConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		AutoScalingPolicy: &AutoScalingPolicyProperty{
//   			Constraints: &ScalingConstraintsProperty{
//   				MaxCapacity: jsii.Number(123),
//   				MinCapacity: jsii.Number(123),
//   			},
//   			Rules: []interface{}{
//   				&ScalingRuleProperty{
//   					Action: &ScalingActionProperty{
//   						SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   							ScalingAdjustment: jsii.Number(123),
//
//   							// the properties below are optional
//   							AdjustmentType: jsii.String("adjustmentType"),
//   							CoolDown: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						Market: jsii.String("market"),
//   					},
//   					Name: jsii.String("name"),
//   					Trigger: &ScalingTriggerProperty{
//   						CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   							ComparisonOperator: jsii.String("comparisonOperator"),
//   							MetricName: jsii.String("metricName"),
//   							Period: jsii.Number(123),
//   							Threshold: jsii.Number(123),
//
//   							// the properties below are optional
//   							Dimensions: []interface{}{
//   								&MetricDimensionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							EvaluationPeriods: jsii.Number(123),
//   							Namespace: jsii.String("namespace"),
//   							Statistic: jsii.String("statistic"),
//   							Unit: jsii.String("unit"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   				},
//   			},
//   		},
//   		BidPrice: jsii.String("bidPrice"),
//   		Configurations: []interface{}{
//   			&configurationProperty{
//   				Classification: jsii.String("classification"),
//   				ConfigurationProperties: map[string]*string{
//   					"configurationPropertiesKey": jsii.String("configurationProperties"),
//   				},
//   				Configurations: []interface{}{
//   					configurationProperty_,
//   				},
//   			},
//   		},
//   		CustomAmiId: jsii.String("customAmiId"),
//   		EbsConfiguration: &EbsConfigurationProperty{
//   			EbsBlockDeviceConfigs: []interface{}{
//   				&EbsBlockDeviceConfigProperty{
//   					VolumeSpecification: &VolumeSpecificationProperty{
//   						SizeInGb: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//
//   						// the properties below are optional
//   						Iops: jsii.Number(123),
//   						Throughput: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					VolumesPerInstance: jsii.Number(123),
//   				},
//   			},
//   			EbsOptimized: jsii.Boolean(false),
//   		},
//   		Market: jsii.String("market"),
//   		Name: jsii.String("name"),
//   	},
//   	Ec2KeyName: jsii.String("ec2KeyName"),
//   	Ec2SubnetId: jsii.String("ec2SubnetId"),
//   	Ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	EmrManagedMasterSecurityGroup: jsii.String("emrManagedMasterSecurityGroup"),
//   	EmrManagedSlaveSecurityGroup: jsii.String("emrManagedSlaveSecurityGroup"),
//   	HadoopVersion: jsii.String("hadoopVersion"),
//   	KeepJobFlowAliveWhenNoSteps: jsii.Boolean(false),
//   	MasterInstanceFleet: &InstanceFleetConfigProperty{
//   		InstanceTypeConfigs: []interface{}{
//   			&InstanceTypeConfigProperty{
//   				InstanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				BidPrice: jsii.String("bidPrice"),
//   				BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   				Configurations: []interface{}{
//   					&configurationProperty{
//   						Classification: jsii.String("classification"),
//   						ConfigurationProperties: map[string]*string{
//   							"configurationPropertiesKey": jsii.String("configurationProperties"),
//   						},
//   						Configurations: []interface{}{
//   							configurationProperty_,
//   						},
//   					},
//   				},
//   				CustomAmiId: jsii.String("customAmiId"),
//   				EbsConfiguration: &EbsConfigurationProperty{
//   					EbsBlockDeviceConfigs: []interface{}{
//   						&EbsBlockDeviceConfigProperty{
//   							VolumeSpecification: &VolumeSpecificationProperty{
//   								SizeInGb: jsii.Number(123),
//   								VolumeType: jsii.String("volumeType"),
//
//   								// the properties below are optional
//   								Iops: jsii.Number(123),
//   								Throughput: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							VolumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					EbsOptimized: jsii.Boolean(false),
//   				},
//   				WeightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   			OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   				AllocationStrategy: jsii.String("allocationStrategy"),
//   			},
//   			SpotSpecification: &SpotProvisioningSpecificationProperty{
//   				TimeoutAction: jsii.String("timeoutAction"),
//   				TimeoutDurationMinutes: jsii.Number(123),
//
//   				// the properties below are optional
//   				AllocationStrategy: jsii.String("allocationStrategy"),
//   				BlockDurationMinutes: jsii.Number(123),
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		TargetOnDemandCapacity: jsii.Number(123),
//   		TargetSpotCapacity: jsii.Number(123),
//   	},
//   	MasterInstanceGroup: &InstanceGroupConfigProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		AutoScalingPolicy: &AutoScalingPolicyProperty{
//   			Constraints: &ScalingConstraintsProperty{
//   				MaxCapacity: jsii.Number(123),
//   				MinCapacity: jsii.Number(123),
//   			},
//   			Rules: []interface{}{
//   				&ScalingRuleProperty{
//   					Action: &ScalingActionProperty{
//   						SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   							ScalingAdjustment: jsii.Number(123),
//
//   							// the properties below are optional
//   							AdjustmentType: jsii.String("adjustmentType"),
//   							CoolDown: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						Market: jsii.String("market"),
//   					},
//   					Name: jsii.String("name"),
//   					Trigger: &ScalingTriggerProperty{
//   						CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   							ComparisonOperator: jsii.String("comparisonOperator"),
//   							MetricName: jsii.String("metricName"),
//   							Period: jsii.Number(123),
//   							Threshold: jsii.Number(123),
//
//   							// the properties below are optional
//   							Dimensions: []interface{}{
//   								&MetricDimensionProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							EvaluationPeriods: jsii.Number(123),
//   							Namespace: jsii.String("namespace"),
//   							Statistic: jsii.String("statistic"),
//   							Unit: jsii.String("unit"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   				},
//   			},
//   		},
//   		BidPrice: jsii.String("bidPrice"),
//   		Configurations: []interface{}{
//   			&configurationProperty{
//   				Classification: jsii.String("classification"),
//   				ConfigurationProperties: map[string]*string{
//   					"configurationPropertiesKey": jsii.String("configurationProperties"),
//   				},
//   				Configurations: []interface{}{
//   					configurationProperty_,
//   				},
//   			},
//   		},
//   		CustomAmiId: jsii.String("customAmiId"),
//   		EbsConfiguration: &EbsConfigurationProperty{
//   			EbsBlockDeviceConfigs: []interface{}{
//   				&EbsBlockDeviceConfigProperty{
//   					VolumeSpecification: &VolumeSpecificationProperty{
//   						SizeInGb: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//
//   						// the properties below are optional
//   						Iops: jsii.Number(123),
//   						Throughput: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					VolumesPerInstance: jsii.Number(123),
//   				},
//   			},
//   			EbsOptimized: jsii.Boolean(false),
//   		},
//   		Market: jsii.String("market"),
//   		Name: jsii.String("name"),
//   	},
//   	Placement: &PlacementTypeProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   	},
//   	ServiceAccessSecurityGroup: jsii.String("serviceAccessSecurityGroup"),
//   	TaskInstanceFleets: []interface{}{
//   		&InstanceFleetConfigProperty{
//   			InstanceTypeConfigs: []interface{}{
//   				&InstanceTypeConfigProperty{
//   					InstanceType: jsii.String("instanceType"),
//
//   					// the properties below are optional
//   					BidPrice: jsii.String("bidPrice"),
//   					BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   					Configurations: []interface{}{
//   						&configurationProperty{
//   							Classification: jsii.String("classification"),
//   							ConfigurationProperties: map[string]*string{
//   								"configurationPropertiesKey": jsii.String("configurationProperties"),
//   							},
//   							Configurations: []interface{}{
//   								configurationProperty_,
//   							},
//   						},
//   					},
//   					CustomAmiId: jsii.String("customAmiId"),
//   					EbsConfiguration: &EbsConfigurationProperty{
//   						EbsBlockDeviceConfigs: []interface{}{
//   							&EbsBlockDeviceConfigProperty{
//   								VolumeSpecification: &VolumeSpecificationProperty{
//   									SizeInGb: jsii.Number(123),
//   									VolumeType: jsii.String("volumeType"),
//
//   									// the properties below are optional
//   									Iops: jsii.Number(123),
//   									Throughput: jsii.Number(123),
//   								},
//
//   								// the properties below are optional
//   								VolumesPerInstance: jsii.Number(123),
//   							},
//   						},
//   						EbsOptimized: jsii.Boolean(false),
//   					},
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   			LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   				OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   				},
//   				SpotSpecification: &SpotProvisioningSpecificationProperty{
//   					TimeoutAction: jsii.String("timeoutAction"),
//   					TimeoutDurationMinutes: jsii.Number(123),
//
//   					// the properties below are optional
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					BlockDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			TargetOnDemandCapacity: jsii.Number(123),
//   			TargetSpotCapacity: jsii.Number(123),
//   		},
//   	},
//   	TaskInstanceGroups: []interface{}{
//   		&InstanceGroupConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			AutoScalingPolicy: &AutoScalingPolicyProperty{
//   				Constraints: &ScalingConstraintsProperty{
//   					MaxCapacity: jsii.Number(123),
//   					MinCapacity: jsii.Number(123),
//   				},
//   				Rules: []interface{}{
//   					&ScalingRuleProperty{
//   						Action: &ScalingActionProperty{
//   							SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   								ScalingAdjustment: jsii.Number(123),
//
//   								// the properties below are optional
//   								AdjustmentType: jsii.String("adjustmentType"),
//   								CoolDown: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							Market: jsii.String("market"),
//   						},
//   						Name: jsii.String("name"),
//   						Trigger: &ScalingTriggerProperty{
//   							CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   								ComparisonOperator: jsii.String("comparisonOperator"),
//   								MetricName: jsii.String("metricName"),
//   								Period: jsii.Number(123),
//   								Threshold: jsii.Number(123),
//
//   								// the properties below are optional
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								EvaluationPeriods: jsii.Number(123),
//   								Namespace: jsii.String("namespace"),
//   								Statistic: jsii.String("statistic"),
//   								Unit: jsii.String("unit"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						Description: jsii.String("description"),
//   					},
//   				},
//   			},
//   			BidPrice: jsii.String("bidPrice"),
//   			Configurations: []interface{}{
//   				&configurationProperty{
//   					Classification: jsii.String("classification"),
//   					ConfigurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					Configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			CustomAmiId: jsii.String("customAmiId"),
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []interface{}{
//   					&EbsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							SizeInGb: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//
//   							// the properties below are optional
//   							Iops: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			Market: jsii.String("market"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TerminationProtected: jsii.Boolean(false),
//   	UnhealthyNodeReplacement: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html
//
type CfnCluster_JobFlowInstancesConfigProperty struct {
	// A list of additional Amazon EC2 security group IDs for the master node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	//
	AdditionalMasterSecurityGroups *[]*string `field:"optional" json:"additionalMasterSecurityGroups" yaml:"additionalMasterSecurityGroups"`
	// A list of additional Amazon EC2 security group IDs for the core and task nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	//
	AdditionalSlaveSecurityGroups *[]*string `field:"optional" json:"additionalSlaveSecurityGroups" yaml:"additionalSlaveSecurityGroups"`
	// Describes the EC2 instances and instance configurations for the core instance fleet when using clusters with the instance fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancefleet
	//
	CoreInstanceFleet interface{} `field:"optional" json:"coreInstanceFleet" yaml:"coreInstanceFleet"`
	// Describes the EC2 instances and instance configurations for core instance groups when using clusters with the uniform instance group configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	//
	CoreInstanceGroup interface{} `field:"optional" json:"coreInstanceGroup" yaml:"coreInstanceGroup"`
	// The name of the Amazon EC2 key pair that can be used to connect to the master node using SSH as the user called "hadoop.".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2keyname
	//
	Ec2KeyName *string `field:"optional" json:"ec2KeyName" yaml:"ec2KeyName"`
	// Applies to clusters that use the uniform instance group configuration.
	//
	// To launch the cluster in Amazon Virtual Private Cloud (Amazon VPC), set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch. If you do not specify this value and your account supports EC2-Classic, the cluster launches in EC2-Classic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2subnetid
	//
	Ec2SubnetId *string `field:"optional" json:"ec2SubnetId" yaml:"ec2SubnetId"`
	// Applies to clusters that use the instance fleet configuration.
	//
	// When multiple Amazon EC2 subnet IDs are specified, Amazon EMR evaluates them and launches instances in the optimal subnet.
	//
	// > The instance fleet configuration is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2subnetids
	//
	Ec2SubnetIds *[]*string `field:"optional" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The identifier of the Amazon EC2 security group for the master node.
	//
	// If you specify `EmrManagedMasterSecurityGroup` , you must also specify `EmrManagedSlaveSecurityGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedmastersecuritygroup
	//
	EmrManagedMasterSecurityGroup *string `field:"optional" json:"emrManagedMasterSecurityGroup" yaml:"emrManagedMasterSecurityGroup"`
	// The identifier of the Amazon EC2 security group for the core and task nodes.
	//
	// If you specify `EmrManagedSlaveSecurityGroup` , you must also specify `EmrManagedMasterSecurityGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedslavesecuritygroup
	//
	EmrManagedSlaveSecurityGroup *string `field:"optional" json:"emrManagedSlaveSecurityGroup" yaml:"emrManagedSlaveSecurityGroup"`
	// Applies only to Amazon EMR release versions earlier than 4.0. The Hadoop version for the cluster. Valid inputs are "0.18" (no longer maintained), "0.20" (no longer maintained), "0.20.205" (no longer maintained), "1.0.3", "2.2.0", or "2.4.0". If you do not set this value, the default of 0.18 is used, unless the `AmiVersion` parameter is set in the RunJobFlow call, in which case the default version of Hadoop for that AMI version is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-hadoopversion
	//
	HadoopVersion *string `field:"optional" json:"hadoopVersion" yaml:"hadoopVersion"`
	// Specifies whether the cluster should remain available after completing all steps.
	//
	// Defaults to `true` . For more information about configuring cluster termination, see [Control Cluster Termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html) in the *EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-keepjobflowalivewhennosteps
	//
	KeepJobFlowAliveWhenNoSteps interface{} `field:"optional" json:"keepJobFlowAliveWhenNoSteps" yaml:"keepJobFlowAliveWhenNoSteps"`
	// Describes the EC2 instances and instance configurations for the master instance fleet when using clusters with the instance fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-masterinstancefleet
	//
	MasterInstanceFleet interface{} `field:"optional" json:"masterInstanceFleet" yaml:"masterInstanceFleet"`
	// Describes the EC2 instances and instance configurations for the master instance group when using clusters with the uniform instance group configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-masterinstancegroup
	//
	MasterInstanceGroup interface{} `field:"optional" json:"masterInstanceGroup" yaml:"masterInstanceGroup"`
	// The Availability Zone in which the cluster runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-placement
	//
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-serviceaccesssecuritygroup
	//
	ServiceAccessSecurityGroup *string `field:"optional" json:"serviceAccessSecurityGroup" yaml:"serviceAccessSecurityGroup"`
	// Describes the EC2 instances and instance configurations for the task instance fleets when using clusters with the instance fleet configuration.
	//
	// These task instance fleets are added to the cluster as part of the cluster launch. Each task instance fleet must have a unique name specified so that CloudFormation can differentiate between the task instance fleets.
	//
	// > You can currently specify only one task instance fleet for a cluster. After creating the cluster, you can only modify the mutable properties of `InstanceFleetConfig` , which are `TargetOnDemandCapacity` and `TargetSpotCapacity` . Modifying any other property results in cluster replacement. > To allow a maximum of 30 Amazon EC2 instance types per fleet, include `TaskInstanceFleets` when you create your cluster. If you create your cluster without `TaskInstanceFleets` , Amazon EMR uses its default allocation strategy, which allows for a maximum of five Amazon EC2 instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-taskinstancefleets
	//
	TaskInstanceFleets interface{} `field:"optional" json:"taskInstanceFleets" yaml:"taskInstanceFleets"`
	// Describes the EC2 instances and instance configurations for task instance groups when using clusters with the uniform instance group configuration.
	//
	// These task instance groups are added to the cluster as part of the cluster launch. Each task instance group must have a unique name specified so that CloudFormation can differentiate between the task instance groups.
	//
	// > After creating the cluster, you can only modify the mutable properties of `InstanceGroupConfig` , which are `AutoScalingPolicy` and `InstanceCount` . Modifying any other property results in cluster replacement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-taskinstancegroups
	//
	TaskInstanceGroups interface{} `field:"optional" json:"taskInstanceGroups" yaml:"taskInstanceGroups"`
	// Specifies whether to lock the cluster to prevent the Amazon EC2 instances from being terminated by API call, user intervention, or in the event of a job-flow error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-terminationprotected
	//
	TerminationProtected interface{} `field:"optional" json:"terminationProtected" yaml:"terminationProtected"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-unhealthynodereplacement
	//
	UnhealthyNodeReplacement interface{} `field:"optional" json:"unhealthyNodeReplacement" yaml:"unhealthyNodeReplacement"`
}

