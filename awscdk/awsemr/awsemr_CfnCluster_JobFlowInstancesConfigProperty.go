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
//   jobFlowInstancesConfigProperty := &jobFlowInstancesConfigProperty{
//   	additionalMasterSecurityGroups: []*string{
//   		jsii.String("additionalMasterSecurityGroups"),
//   	},
//   	additionalSlaveSecurityGroups: []*string{
//   		jsii.String("additionalSlaveSecurityGroups"),
//   	},
//   	coreInstanceFleet: &instanceFleetConfigProperty{
//   		instanceTypeConfigs: []interface{}{
//   			&instanceTypeConfigProperty{
//   				instanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				bidPrice: jsii.String("bidPrice"),
//   				bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   				configurations: []interface{}{
//   					&configurationProperty{
//   						classification: jsii.String("classification"),
//   						configurationProperties: map[string]*string{
//   							"configurationPropertiesKey": jsii.String("configurationProperties"),
//   						},
//   						configurations: []interface{}{
//   							configurationProperty_,
//   						},
//   					},
//   				},
//   				customAmiId: jsii.String("customAmiId"),
//   				ebsConfiguration: &ebsConfigurationProperty{
//   					ebsBlockDeviceConfigs: []interface{}{
//   						&ebsBlockDeviceConfigProperty{
//   							volumeSpecification: &volumeSpecificationProperty{
//   								sizeInGb: jsii.Number(123),
//   								volumeType: jsii.String("volumeType"),
//
//   								// the properties below are optional
//   								iops: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							volumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					ebsOptimized: jsii.Boolean(false),
//   				},
//   				weightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   			onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   				allocationStrategy: jsii.String("allocationStrategy"),
//   			},
//   			spotSpecification: &spotProvisioningSpecificationProperty{
//   				timeoutAction: jsii.String("timeoutAction"),
//   				timeoutDurationMinutes: jsii.Number(123),
//
//   				// the properties below are optional
//   				allocationStrategy: jsii.String("allocationStrategy"),
//   				blockDurationMinutes: jsii.Number(123),
//   			},
//   		},
//   		name: jsii.String("name"),
//   		targetOnDemandCapacity: jsii.Number(123),
//   		targetSpotCapacity: jsii.Number(123),
//   	},
//   	coreInstanceGroup: &instanceGroupConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		autoScalingPolicy: &autoScalingPolicyProperty{
//   			constraints: &scalingConstraintsProperty{
//   				maxCapacity: jsii.Number(123),
//   				minCapacity: jsii.Number(123),
//   			},
//   			rules: []interface{}{
//   				&scalingRuleProperty{
//   					action: &scalingActionProperty{
//   						simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   							scalingAdjustment: jsii.Number(123),
//
//   							// the properties below are optional
//   							adjustmentType: jsii.String("adjustmentType"),
//   							coolDown: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						market: jsii.String("market"),
//   					},
//   					name: jsii.String("name"),
//   					trigger: &scalingTriggerProperty{
//   						cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   							comparisonOperator: jsii.String("comparisonOperator"),
//   							metricName: jsii.String("metricName"),
//   							period: jsii.Number(123),
//   							threshold: jsii.Number(123),
//
//   							// the properties below are optional
//   							dimensions: []interface{}{
//   								&metricDimensionProperty{
//   									key: jsii.String("key"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   							evaluationPeriods: jsii.Number(123),
//   							namespace: jsii.String("namespace"),
//   							statistic: jsii.String("statistic"),
//   							unit: jsii.String("unit"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					description: jsii.String("description"),
//   				},
//   			},
//   		},
//   		bidPrice: jsii.String("bidPrice"),
//   		configurations: []interface{}{
//   			&configurationProperty{
//   				classification: jsii.String("classification"),
//   				configurationProperties: map[string]*string{
//   					"configurationPropertiesKey": jsii.String("configurationProperties"),
//   				},
//   				configurations: []interface{}{
//   					configurationProperty_,
//   				},
//   			},
//   		},
//   		customAmiId: jsii.String("customAmiId"),
//   		ebsConfiguration: &ebsConfigurationProperty{
//   			ebsBlockDeviceConfigs: []interface{}{
//   				&ebsBlockDeviceConfigProperty{
//   					volumeSpecification: &volumeSpecificationProperty{
//   						sizeInGb: jsii.Number(123),
//   						volumeType: jsii.String("volumeType"),
//
//   						// the properties below are optional
//   						iops: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					volumesPerInstance: jsii.Number(123),
//   				},
//   			},
//   			ebsOptimized: jsii.Boolean(false),
//   		},
//   		market: jsii.String("market"),
//   		name: jsii.String("name"),
//   	},
//   	ec2KeyName: jsii.String("ec2KeyName"),
//   	ec2SubnetId: jsii.String("ec2SubnetId"),
//   	ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	emrManagedMasterSecurityGroup: jsii.String("emrManagedMasterSecurityGroup"),
//   	emrManagedSlaveSecurityGroup: jsii.String("emrManagedSlaveSecurityGroup"),
//   	hadoopVersion: jsii.String("hadoopVersion"),
//   	keepJobFlowAliveWhenNoSteps: jsii.Boolean(false),
//   	masterInstanceFleet: &instanceFleetConfigProperty{
//   		instanceTypeConfigs: []interface{}{
//   			&instanceTypeConfigProperty{
//   				instanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				bidPrice: jsii.String("bidPrice"),
//   				bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   				configurations: []interface{}{
//   					&configurationProperty{
//   						classification: jsii.String("classification"),
//   						configurationProperties: map[string]*string{
//   							"configurationPropertiesKey": jsii.String("configurationProperties"),
//   						},
//   						configurations: []interface{}{
//   							configurationProperty_,
//   						},
//   					},
//   				},
//   				customAmiId: jsii.String("customAmiId"),
//   				ebsConfiguration: &ebsConfigurationProperty{
//   					ebsBlockDeviceConfigs: []interface{}{
//   						&ebsBlockDeviceConfigProperty{
//   							volumeSpecification: &volumeSpecificationProperty{
//   								sizeInGb: jsii.Number(123),
//   								volumeType: jsii.String("volumeType"),
//
//   								// the properties below are optional
//   								iops: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							volumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					ebsOptimized: jsii.Boolean(false),
//   				},
//   				weightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   			onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   				allocationStrategy: jsii.String("allocationStrategy"),
//   			},
//   			spotSpecification: &spotProvisioningSpecificationProperty{
//   				timeoutAction: jsii.String("timeoutAction"),
//   				timeoutDurationMinutes: jsii.Number(123),
//
//   				// the properties below are optional
//   				allocationStrategy: jsii.String("allocationStrategy"),
//   				blockDurationMinutes: jsii.Number(123),
//   			},
//   		},
//   		name: jsii.String("name"),
//   		targetOnDemandCapacity: jsii.Number(123),
//   		targetSpotCapacity: jsii.Number(123),
//   	},
//   	masterInstanceGroup: &instanceGroupConfigProperty{
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		autoScalingPolicy: &autoScalingPolicyProperty{
//   			constraints: &scalingConstraintsProperty{
//   				maxCapacity: jsii.Number(123),
//   				minCapacity: jsii.Number(123),
//   			},
//   			rules: []interface{}{
//   				&scalingRuleProperty{
//   					action: &scalingActionProperty{
//   						simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   							scalingAdjustment: jsii.Number(123),
//
//   							// the properties below are optional
//   							adjustmentType: jsii.String("adjustmentType"),
//   							coolDown: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						market: jsii.String("market"),
//   					},
//   					name: jsii.String("name"),
//   					trigger: &scalingTriggerProperty{
//   						cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   							comparisonOperator: jsii.String("comparisonOperator"),
//   							metricName: jsii.String("metricName"),
//   							period: jsii.Number(123),
//   							threshold: jsii.Number(123),
//
//   							// the properties below are optional
//   							dimensions: []interface{}{
//   								&metricDimensionProperty{
//   									key: jsii.String("key"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   							evaluationPeriods: jsii.Number(123),
//   							namespace: jsii.String("namespace"),
//   							statistic: jsii.String("statistic"),
//   							unit: jsii.String("unit"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					description: jsii.String("description"),
//   				},
//   			},
//   		},
//   		bidPrice: jsii.String("bidPrice"),
//   		configurations: []interface{}{
//   			&configurationProperty{
//   				classification: jsii.String("classification"),
//   				configurationProperties: map[string]*string{
//   					"configurationPropertiesKey": jsii.String("configurationProperties"),
//   				},
//   				configurations: []interface{}{
//   					configurationProperty_,
//   				},
//   			},
//   		},
//   		customAmiId: jsii.String("customAmiId"),
//   		ebsConfiguration: &ebsConfigurationProperty{
//   			ebsBlockDeviceConfigs: []interface{}{
//   				&ebsBlockDeviceConfigProperty{
//   					volumeSpecification: &volumeSpecificationProperty{
//   						sizeInGb: jsii.Number(123),
//   						volumeType: jsii.String("volumeType"),
//
//   						// the properties below are optional
//   						iops: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					volumesPerInstance: jsii.Number(123),
//   				},
//   			},
//   			ebsOptimized: jsii.Boolean(false),
//   		},
//   		market: jsii.String("market"),
//   		name: jsii.String("name"),
//   	},
//   	placement: &placementTypeProperty{
//   		availabilityZone: jsii.String("availabilityZone"),
//   	},
//   	serviceAccessSecurityGroup: jsii.String("serviceAccessSecurityGroup"),
//   	taskInstanceFleets: []interface{}{
//   		&instanceFleetConfigProperty{
//   			instanceTypeConfigs: []interface{}{
//   				&instanceTypeConfigProperty{
//   					instanceType: jsii.String("instanceType"),
//
//   					// the properties below are optional
//   					bidPrice: jsii.String("bidPrice"),
//   					bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   					configurations: []interface{}{
//   						&configurationProperty{
//   							classification: jsii.String("classification"),
//   							configurationProperties: map[string]*string{
//   								"configurationPropertiesKey": jsii.String("configurationProperties"),
//   							},
//   							configurations: []interface{}{
//   								configurationProperty_,
//   							},
//   						},
//   					},
//   					customAmiId: jsii.String("customAmiId"),
//   					ebsConfiguration: &ebsConfigurationProperty{
//   						ebsBlockDeviceConfigs: []interface{}{
//   							&ebsBlockDeviceConfigProperty{
//   								volumeSpecification: &volumeSpecificationProperty{
//   									sizeInGb: jsii.Number(123),
//   									volumeType: jsii.String("volumeType"),
//
//   									// the properties below are optional
//   									iops: jsii.Number(123),
//   								},
//
//   								// the properties below are optional
//   								volumesPerInstance: jsii.Number(123),
//   							},
//   						},
//   						ebsOptimized: jsii.Boolean(false),
//   					},
//   					weightedCapacity: jsii.Number(123),
//   				},
//   			},
//   			launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   				onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   					allocationStrategy: jsii.String("allocationStrategy"),
//   				},
//   				spotSpecification: &spotProvisioningSpecificationProperty{
//   					timeoutAction: jsii.String("timeoutAction"),
//   					timeoutDurationMinutes: jsii.Number(123),
//
//   					// the properties below are optional
//   					allocationStrategy: jsii.String("allocationStrategy"),
//   					blockDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			name: jsii.String("name"),
//   			targetOnDemandCapacity: jsii.Number(123),
//   			targetSpotCapacity: jsii.Number(123),
//   		},
//   	},
//   	taskInstanceGroups: []interface{}{
//   		&instanceGroupConfigProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			autoScalingPolicy: &autoScalingPolicyProperty{
//   				constraints: &scalingConstraintsProperty{
//   					maxCapacity: jsii.Number(123),
//   					minCapacity: jsii.Number(123),
//   				},
//   				rules: []interface{}{
//   					&scalingRuleProperty{
//   						action: &scalingActionProperty{
//   							simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   								scalingAdjustment: jsii.Number(123),
//
//   								// the properties below are optional
//   								adjustmentType: jsii.String("adjustmentType"),
//   								coolDown: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							market: jsii.String("market"),
//   						},
//   						name: jsii.String("name"),
//   						trigger: &scalingTriggerProperty{
//   							cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   								comparisonOperator: jsii.String("comparisonOperator"),
//   								metricName: jsii.String("metricName"),
//   								period: jsii.Number(123),
//   								threshold: jsii.Number(123),
//
//   								// the properties below are optional
//   								dimensions: []interface{}{
//   									&metricDimensionProperty{
//   										key: jsii.String("key"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								evaluationPeriods: jsii.Number(123),
//   								namespace: jsii.String("namespace"),
//   								statistic: jsii.String("statistic"),
//   								unit: jsii.String("unit"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						description: jsii.String("description"),
//   					},
//   				},
//   			},
//   			bidPrice: jsii.String("bidPrice"),
//   			configurations: []interface{}{
//   				&configurationProperty{
//   					classification: jsii.String("classification"),
//   					configurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			customAmiId: jsii.String("customAmiId"),
//   			ebsConfiguration: &ebsConfigurationProperty{
//   				ebsBlockDeviceConfigs: []interface{}{
//   					&ebsBlockDeviceConfigProperty{
//   						volumeSpecification: &volumeSpecificationProperty{
//   							sizeInGb: jsii.Number(123),
//   							volumeType: jsii.String("volumeType"),
//
//   							// the properties below are optional
//   							iops: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						volumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   			},
//   			market: jsii.String("market"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	terminationProtected: jsii.Boolean(false),
//   }
//
type CfnCluster_JobFlowInstancesConfigProperty struct {
	// A list of additional Amazon EC2 security group IDs for the master node.
	AdditionalMasterSecurityGroups *[]*string `field:"optional" json:"additionalMasterSecurityGroups" yaml:"additionalMasterSecurityGroups"`
	// A list of additional Amazon EC2 security group IDs for the core and task nodes.
	AdditionalSlaveSecurityGroups *[]*string `field:"optional" json:"additionalSlaveSecurityGroups" yaml:"additionalSlaveSecurityGroups"`
	// Describes the EC2 instances and instance configurations for the core instance fleet when using clusters with the instance fleet configuration.
	CoreInstanceFleet interface{} `field:"optional" json:"coreInstanceFleet" yaml:"coreInstanceFleet"`
	// Describes the EC2 instances and instance configurations for core instance groups when using clusters with the uniform instance group configuration.
	CoreInstanceGroup interface{} `field:"optional" json:"coreInstanceGroup" yaml:"coreInstanceGroup"`
	// The name of the EC2 key pair that can be used to connect to the master node using SSH as the user called "hadoop.".
	Ec2KeyName *string `field:"optional" json:"ec2KeyName" yaml:"ec2KeyName"`
	// Applies to clusters that use the uniform instance group configuration.
	//
	// To launch the cluster in Amazon Virtual Private Cloud (Amazon VPC), set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch. If you do not specify this value and your account supports EC2-Classic, the cluster launches in EC2-Classic.
	Ec2SubnetId *string `field:"optional" json:"ec2SubnetId" yaml:"ec2SubnetId"`
	// Applies to clusters that use the instance fleet configuration.
	//
	// When multiple EC2 subnet IDs are specified, Amazon EMR evaluates them and launches instances in the optimal subnet.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	Ec2SubnetIds *[]*string `field:"optional" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The identifier of the Amazon EC2 security group for the master node.
	//
	// If you specify `EmrManagedMasterSecurityGroup` , you must also specify `EmrManagedSlaveSecurityGroup` .
	EmrManagedMasterSecurityGroup *string `field:"optional" json:"emrManagedMasterSecurityGroup" yaml:"emrManagedMasterSecurityGroup"`
	// The identifier of the Amazon EC2 security group for the core and task nodes.
	//
	// If you specify `EmrManagedSlaveSecurityGroup` , you must also specify `EmrManagedMasterSecurityGroup` .
	EmrManagedSlaveSecurityGroup *string `field:"optional" json:"emrManagedSlaveSecurityGroup" yaml:"emrManagedSlaveSecurityGroup"`
	// Applies only to Amazon EMR release versions earlier than 4.0. The Hadoop version for the cluster. Valid inputs are "0.18" (no longer maintained), "0.20" (no longer maintained), "0.20.205" (no longer maintained), "1.0.3", "2.2.0", or "2.4.0". If you do not set this value, the default of 0.18 is used, unless the `AmiVersion` parameter is set in the RunJobFlow call, in which case the default version of Hadoop for that AMI version is used.
	HadoopVersion *string `field:"optional" json:"hadoopVersion" yaml:"hadoopVersion"`
	// Specifies whether the cluster should remain available after completing all steps.
	//
	// Defaults to `true` . For more information about configuring cluster termination, see [Control Cluster Termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html) in the *EMR Management Guide* .
	KeepJobFlowAliveWhenNoSteps interface{} `field:"optional" json:"keepJobFlowAliveWhenNoSteps" yaml:"keepJobFlowAliveWhenNoSteps"`
	// Describes the EC2 instances and instance configurations for the master instance fleet when using clusters with the instance fleet configuration.
	MasterInstanceFleet interface{} `field:"optional" json:"masterInstanceFleet" yaml:"masterInstanceFleet"`
	// Describes the EC2 instances and instance configurations for the master instance group when using clusters with the uniform instance group configuration.
	MasterInstanceGroup interface{} `field:"optional" json:"masterInstanceGroup" yaml:"masterInstanceGroup"`
	// The Availability Zone in which the cluster runs.
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.
	ServiceAccessSecurityGroup *string `field:"optional" json:"serviceAccessSecurityGroup" yaml:"serviceAccessSecurityGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.TaskInstanceFleets`.
	TaskInstanceFleets interface{} `field:"optional" json:"taskInstanceFleets" yaml:"taskInstanceFleets"`
	// `CfnCluster.JobFlowInstancesConfigProperty.TaskInstanceGroups`.
	TaskInstanceGroups interface{} `field:"optional" json:"taskInstanceGroups" yaml:"taskInstanceGroups"`
	// Specifies whether to lock the cluster to prevent the Amazon EC2 instances from being terminated by API call, user intervention, or in the event of a job-flow error.
	TerminationProtected interface{} `field:"optional" json:"terminationProtected" yaml:"terminationProtected"`
}

