package previewawsemrmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalInfo interface{}
//   var configurationProperty_ ConfigurationProperty
//
//   cfnClusterMixinProps := &CfnClusterMixinProps{
//   	AdditionalInfo: additionalInfo,
//   	Applications: []interface{}{
//   		&ApplicationProperty{
//   			AdditionalInfo: map[string]*string{
//   				"additionalInfoKey": jsii.String("additionalInfo"),
//   			},
//   			Args: []*string{
//   				jsii.String("args"),
//   			},
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	AutoScalingRole: jsii.String("autoScalingRole"),
//   	AutoTerminationPolicy: &AutoTerminationPolicyProperty{
//   		IdleTimeout: jsii.Number(123),
//   	},
//   	BootstrapActions: []interface{}{
//   		&BootstrapActionConfigProperty{
//   			Name: jsii.String("name"),
//   			ScriptBootstrapAction: &ScriptBootstrapActionConfigProperty{
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				Path: jsii.String("path"),
//   			},
//   		},
//   	},
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
//   	EbsRootVolumeIops: jsii.Number(123),
//   	EbsRootVolumeSize: jsii.Number(123),
//   	EbsRootVolumeThroughput: jsii.Number(123),
//   	Instances: &JobFlowInstancesConfigProperty{
//   		AdditionalMasterSecurityGroups: []*string{
//   			jsii.String("additionalMasterSecurityGroups"),
//   		},
//   		AdditionalSlaveSecurityGroups: []*string{
//   			jsii.String("additionalSlaveSecurityGroups"),
//   		},
//   		CoreInstanceFleet: &InstanceFleetConfigProperty{
//   			InstanceTypeConfigs: []interface{}{
//   				&InstanceTypeConfigProperty{
//   					BidPrice: jsii.String("bidPrice"),
//   					BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   					Configurations: []interface{}{
//   						&ConfigurationProperty{
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
//   									Iops: jsii.Number(123),
//   									SizeInGb: jsii.Number(123),
//   									Throughput: jsii.Number(123),
//   									VolumeType: jsii.String("volumeType"),
//   								},
//   								VolumesPerInstance: jsii.Number(123),
//   							},
//   						},
//   						EbsOptimized: jsii.Boolean(false),
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					Priority: jsii.Number(123),
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   			LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   				OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   						CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   						CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   						UsageStrategy: jsii.String("usageStrategy"),
//   					},
//   				},
//   				SpotSpecification: &SpotProvisioningSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					BlockDurationMinutes: jsii.Number(123),
//   					TimeoutAction: jsii.String("timeoutAction"),
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			ResizeSpecifications: &InstanceFleetResizingSpecificationsProperty{
//   				OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   						CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   						CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   						UsageStrategy: jsii.String("usageStrategy"),
//   					},
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   				SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			TargetOnDemandCapacity: jsii.Number(123),
//   			TargetSpotCapacity: jsii.Number(123),
//   		},
//   		CoreInstanceGroup: &InstanceGroupConfigProperty{
//   			AutoScalingPolicy: &AutoScalingPolicyProperty{
//   				Constraints: &ScalingConstraintsProperty{
//   					MaxCapacity: jsii.Number(123),
//   					MinCapacity: jsii.Number(123),
//   				},
//   				Rules: []interface{}{
//   					&ScalingRuleProperty{
//   						Action: &ScalingActionProperty{
//   							Market: jsii.String("market"),
//   							SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   								AdjustmentType: jsii.String("adjustmentType"),
//   								CoolDown: jsii.Number(123),
//   								ScalingAdjustment: jsii.Number(123),
//   							},
//   						},
//   						Description: jsii.String("description"),
//   						Name: jsii.String("name"),
//   						Trigger: &ScalingTriggerProperty{
//   							CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   								ComparisonOperator: jsii.String("comparisonOperator"),
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								EvaluationPeriods: jsii.Number(123),
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   								Period: jsii.Number(123),
//   								Statistic: jsii.String("statistic"),
//   								Threshold: jsii.Number(123),
//   								Unit: jsii.String("unit"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			BidPrice: jsii.String("bidPrice"),
//   			Configurations: []interface{}{
//   				&ConfigurationProperty{
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
//   							Iops: jsii.Number(123),
//   							SizeInGb: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			Market: jsii.String("market"),
//   			Name: jsii.String("name"),
//   		},
//   		Ec2KeyName: jsii.String("ec2KeyName"),
//   		Ec2SubnetId: jsii.String("ec2SubnetId"),
//   		Ec2SubnetIds: []*string{
//   			jsii.String("ec2SubnetIds"),
//   		},
//   		EmrManagedMasterSecurityGroup: jsii.String("emrManagedMasterSecurityGroup"),
//   		EmrManagedSlaveSecurityGroup: jsii.String("emrManagedSlaveSecurityGroup"),
//   		HadoopVersion: jsii.String("hadoopVersion"),
//   		KeepJobFlowAliveWhenNoSteps: jsii.Boolean(false),
//   		MasterInstanceFleet: &InstanceFleetConfigProperty{
//   			InstanceTypeConfigs: []interface{}{
//   				&InstanceTypeConfigProperty{
//   					BidPrice: jsii.String("bidPrice"),
//   					BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   					Configurations: []interface{}{
//   						&ConfigurationProperty{
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
//   									Iops: jsii.Number(123),
//   									SizeInGb: jsii.Number(123),
//   									Throughput: jsii.Number(123),
//   									VolumeType: jsii.String("volumeType"),
//   								},
//   								VolumesPerInstance: jsii.Number(123),
//   							},
//   						},
//   						EbsOptimized: jsii.Boolean(false),
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					Priority: jsii.Number(123),
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   			LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   				OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   						CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   						CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   						UsageStrategy: jsii.String("usageStrategy"),
//   					},
//   				},
//   				SpotSpecification: &SpotProvisioningSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					BlockDurationMinutes: jsii.Number(123),
//   					TimeoutAction: jsii.String("timeoutAction"),
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			ResizeSpecifications: &InstanceFleetResizingSpecificationsProperty{
//   				OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   						CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   						CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   						UsageStrategy: jsii.String("usageStrategy"),
//   					},
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   				SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   					AllocationStrategy: jsii.String("allocationStrategy"),
//   					TimeoutDurationMinutes: jsii.Number(123),
//   				},
//   			},
//   			TargetOnDemandCapacity: jsii.Number(123),
//   			TargetSpotCapacity: jsii.Number(123),
//   		},
//   		MasterInstanceGroup: &InstanceGroupConfigProperty{
//   			AutoScalingPolicy: &AutoScalingPolicyProperty{
//   				Constraints: &ScalingConstraintsProperty{
//   					MaxCapacity: jsii.Number(123),
//   					MinCapacity: jsii.Number(123),
//   				},
//   				Rules: []interface{}{
//   					&ScalingRuleProperty{
//   						Action: &ScalingActionProperty{
//   							Market: jsii.String("market"),
//   							SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   								AdjustmentType: jsii.String("adjustmentType"),
//   								CoolDown: jsii.Number(123),
//   								ScalingAdjustment: jsii.Number(123),
//   							},
//   						},
//   						Description: jsii.String("description"),
//   						Name: jsii.String("name"),
//   						Trigger: &ScalingTriggerProperty{
//   							CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   								ComparisonOperator: jsii.String("comparisonOperator"),
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Key: jsii.String("key"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								EvaluationPeriods: jsii.Number(123),
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   								Period: jsii.Number(123),
//   								Statistic: jsii.String("statistic"),
//   								Threshold: jsii.Number(123),
//   								Unit: jsii.String("unit"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			BidPrice: jsii.String("bidPrice"),
//   			Configurations: []interface{}{
//   				&ConfigurationProperty{
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
//   							Iops: jsii.Number(123),
//   							SizeInGb: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			Market: jsii.String("market"),
//   			Name: jsii.String("name"),
//   		},
//   		Placement: &PlacementTypeProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   		},
//   		ServiceAccessSecurityGroup: jsii.String("serviceAccessSecurityGroup"),
//   		TaskInstanceFleets: []interface{}{
//   			&InstanceFleetConfigProperty{
//   				InstanceTypeConfigs: []interface{}{
//   					&InstanceTypeConfigProperty{
//   						BidPrice: jsii.String("bidPrice"),
//   						BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   						Configurations: []interface{}{
//   							&ConfigurationProperty{
//   								Classification: jsii.String("classification"),
//   								ConfigurationProperties: map[string]*string{
//   									"configurationPropertiesKey": jsii.String("configurationProperties"),
//   								},
//   								Configurations: []interface{}{
//   									configurationProperty_,
//   								},
//   							},
//   						},
//   						CustomAmiId: jsii.String("customAmiId"),
//   						EbsConfiguration: &EbsConfigurationProperty{
//   							EbsBlockDeviceConfigs: []interface{}{
//   								&EbsBlockDeviceConfigProperty{
//   									VolumeSpecification: &VolumeSpecificationProperty{
//   										Iops: jsii.Number(123),
//   										SizeInGb: jsii.Number(123),
//   										Throughput: jsii.Number(123),
//   										VolumeType: jsii.String("volumeType"),
//   									},
//   									VolumesPerInstance: jsii.Number(123),
//   								},
//   							},
//   							EbsOptimized: jsii.Boolean(false),
//   						},
//   						InstanceType: jsii.String("instanceType"),
//   						Priority: jsii.Number(123),
//   						WeightedCapacity: jsii.Number(123),
//   					},
//   				},
//   				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   					OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   						CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   							CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   							CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   							UsageStrategy: jsii.String("usageStrategy"),
//   						},
//   					},
//   					SpotSpecification: &SpotProvisioningSpecificationProperty{
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   						BlockDurationMinutes: jsii.Number(123),
//   						TimeoutAction: jsii.String("timeoutAction"),
//   						TimeoutDurationMinutes: jsii.Number(123),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				ResizeSpecifications: &InstanceFleetResizingSpecificationsProperty{
//   					OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   						CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   							CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   							CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   							UsageStrategy: jsii.String("usageStrategy"),
//   						},
//   						TimeoutDurationMinutes: jsii.Number(123),
//   					},
//   					SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   						TimeoutDurationMinutes: jsii.Number(123),
//   					},
//   				},
//   				TargetOnDemandCapacity: jsii.Number(123),
//   				TargetSpotCapacity: jsii.Number(123),
//   			},
//   		},
//   		TaskInstanceGroups: []interface{}{
//   			&InstanceGroupConfigProperty{
//   				AutoScalingPolicy: &AutoScalingPolicyProperty{
//   					Constraints: &ScalingConstraintsProperty{
//   						MaxCapacity: jsii.Number(123),
//   						MinCapacity: jsii.Number(123),
//   					},
//   					Rules: []interface{}{
//   						&ScalingRuleProperty{
//   							Action: &ScalingActionProperty{
//   								Market: jsii.String("market"),
//   								SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   									AdjustmentType: jsii.String("adjustmentType"),
//   									CoolDown: jsii.Number(123),
//   									ScalingAdjustment: jsii.Number(123),
//   								},
//   							},
//   							Description: jsii.String("description"),
//   							Name: jsii.String("name"),
//   							Trigger: &ScalingTriggerProperty{
//   								CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   									ComparisonOperator: jsii.String("comparisonOperator"),
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									EvaluationPeriods: jsii.Number(123),
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   									Period: jsii.Number(123),
//   									Statistic: jsii.String("statistic"),
//   									Threshold: jsii.Number(123),
//   									Unit: jsii.String("unit"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				BidPrice: jsii.String("bidPrice"),
//   				Configurations: []interface{}{
//   					&ConfigurationProperty{
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
//   								Iops: jsii.Number(123),
//   								SizeInGb: jsii.Number(123),
//   								Throughput: jsii.Number(123),
//   								VolumeType: jsii.String("volumeType"),
//   							},
//   							VolumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					EbsOptimized: jsii.Boolean(false),
//   				},
//   				InstanceCount: jsii.Number(123),
//   				InstanceType: jsii.String("instanceType"),
//   				Market: jsii.String("market"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TerminationProtected: jsii.Boolean(false),
//   		UnhealthyNodeReplacement: jsii.Boolean(false),
//   	},
//   	JobFlowRole: jsii.String("jobFlowRole"),
//   	KerberosAttributes: &KerberosAttributesProperty{
//   		AdDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   		AdDomainJoinUser: jsii.String("adDomainJoinUser"),
//   		CrossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   		KdcAdminPassword: jsii.String("kdcAdminPassword"),
//   		Realm: jsii.String("realm"),
//   	},
//   	LogEncryptionKmsKeyId: jsii.String("logEncryptionKmsKeyId"),
//   	LogUri: jsii.String("logUri"),
//   	ManagedScalingPolicy: &ManagedScalingPolicyProperty{
//   		ComputeLimits: &ComputeLimitsProperty{
//   			MaximumCapacityUnits: jsii.Number(123),
//   			MaximumCoreCapacityUnits: jsii.Number(123),
//   			MaximumOnDemandCapacityUnits: jsii.Number(123),
//   			MinimumCapacityUnits: jsii.Number(123),
//   			UnitType: jsii.String("unitType"),
//   		},
//   		ScalingStrategy: jsii.String("scalingStrategy"),
//   		UtilizationPerformanceIndex: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	OsReleaseLabel: jsii.String("osReleaseLabel"),
//   	PlacementGroupConfigs: []interface{}{
//   		&PlacementGroupConfigProperty{
//   			InstanceRole: jsii.String("instanceRole"),
//   			PlacementStrategy: jsii.String("placementStrategy"),
//   		},
//   	},
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	ScaleDownBehavior: jsii.String("scaleDownBehavior"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	ServiceRole: jsii.String("serviceRole"),
//   	StepConcurrencyLevel: jsii.Number(123),
//   	Steps: []interface{}{
//   		&StepConfigProperty{
//   			ActionOnFailure: jsii.String("actionOnFailure"),
//   			HadoopJarStep: &HadoopJarStepConfigProperty{
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				Jar: jsii.String("jar"),
//   				MainClass: jsii.String("mainClass"),
//   				StepProperties: []interface{}{
//   					&KeyValueProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisibleToAllUsers: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html
//
type CfnClusterMixinProps struct {
	// A JSON string for selecting additional features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-additionalinfo
	//
	AdditionalInfo interface{} `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// The applications to install on this cluster, for example, Spark, Flink, Oozie, Zeppelin, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-applications
	//
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
	// An IAM role for automatic scaling policies.
	//
	// The default role is `EMR_AutoScaling_DefaultRole` . The IAM role provides permissions that the automatic scaling feature requires to launch and terminate Amazon EC2 instances in an instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-autoscalingrole
	//
	AutoScalingRole *string `field:"optional" json:"autoScalingRole" yaml:"autoScalingRole"`
	// An auto-termination policy for an Amazon EMR cluster.
	//
	// An auto-termination policy defines the amount of idle time in seconds after which a cluster automatically terminates. For alternative cluster termination options, see [Control cluster termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-autoterminationpolicy
	//
	AutoTerminationPolicy interface{} `field:"optional" json:"autoTerminationPolicy" yaml:"autoTerminationPolicy"`
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-bootstrapactions
	//
	BootstrapActions interface{} `field:"optional" json:"bootstrapActions" yaml:"bootstrapActions"`
	// Applies only to Amazon EMR releases 4.x and later. The list of configurations that are supplied to the Amazon EMR cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// Available only in Amazon EMR releases 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-customamiid
	//
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The IOPS, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.
	//
	// Available in Amazon EMR releases 6.15.0 and later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-ebsrootvolumeiops
	//
	EbsRootVolumeIops *float64 `field:"optional" json:"ebsRootVolumeIops" yaml:"ebsRootVolumeIops"`
	// The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.
	//
	// Available in Amazon EMR releases 4.x and later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-ebsrootvolumesize
	//
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// The throughput, in MiB/s, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.
	//
	// Available in Amazon EMR releases 6.15.0 and later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-ebsrootvolumethroughput
	//
	EbsRootVolumeThroughput *float64 `field:"optional" json:"ebsRootVolumeThroughput" yaml:"ebsRootVolumeThroughput"`
	// A specification of the number and type of Amazon EC2 instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-instances
	//
	Instances interface{} `field:"optional" json:"instances" yaml:"instances"`
	// Also called instance profile and Amazon EC2 role.
	//
	// An IAM role for an Amazon EMR cluster. The Amazon EC2 instances of the cluster assume this role. The default role is `EMR_EC2_DefaultRole` . In order to use the default role, you must have already created it using the AWS CLI or console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-jobflowrole
	//
	JobFlowRole *string `field:"optional" json:"jobFlowRole" yaml:"jobFlowRole"`
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	//
	// For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *Amazon EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-kerberosattributes
	//
	KerberosAttributes interface{} `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// The AWS KMS key used for encrypting log files.
	//
	// This attribute is only available with Amazon EMR 5.30.0 and later, excluding Amazon EMR 6.0.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-logencryptionkmskeyid
	//
	LogEncryptionKmsKeyId *string `field:"optional" json:"logEncryptionKmsKeyId" yaml:"logEncryptionKmsKeyId"`
	// The path to the Amazon S3 location where logs for this cluster are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// Creates or updates a managed scaling policy for an Amazon EMR cluster.
	//
	// The managed scaling policy defines the limits for resources, such as Amazon EC2 instances that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-managedscalingpolicy
	//
	ManagedScalingPolicy interface{} `field:"optional" json:"managedScalingPolicy" yaml:"managedScalingPolicy"`
	// The name of the cluster.
	//
	// This parameter can't contain the characters <, >, $, |, or ` (backtick).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Linux release specified in a cluster launch RunJobFlow request.
	//
	// If no Amazon Linux release was specified, the default Amazon Linux release is shown in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-osreleaselabel
	//
	OsReleaseLabel *string `field:"optional" json:"osReleaseLabel" yaml:"osReleaseLabel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-placementgroupconfigs
	//
	PlacementGroupConfigs interface{} `field:"optional" json:"placementGroupConfigs" yaml:"placementGroupConfigs"`
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	//
	// Release labels are in the form `emr-x.x.x` , where x.x.x is an Amazon EMR release version such as `emr-5.14.0` . For more information about Amazon EMR release versions and included application versions and features, see [](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/) . The release label applies only to Amazon EMR releases version 4.0 and later. Earlier versions use `AmiVersion` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-releaselabel
	//
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	//
	// `TERMINATE_AT_INSTANCE_HOUR` indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted. This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version. `TERMINATE_AT_TASK_COMPLETION` indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to HDFS corruption. `TERMINATE_AT_TASK_COMPLETION` is available only in Amazon EMR releases 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-scaledownbehavior
	//
	ScaleDownBehavior *string `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// The name of the security configuration applied to the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-securityconfiguration
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The IAM role that Amazon EMR assumes in order to access AWS resources on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-servicerole
	//
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Specifies the number of steps that can be executed concurrently.
	//
	// The default value is `1` . The maximum value is `256` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-stepconcurrencylevel
	//
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// A list of steps to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-steps
	//
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// A list of tags associated with a cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.
	//
	// If this value is set to `true` , all IAM users of that AWS account can view and manage the cluster if they have the proper policy permissions set. If this value is `false` , only the IAM user that created the cluster can view and manage it. This value can be changed using the SetVisibleToAllUsers action.
	//
	// > When you create clusters directly through the EMR console or API, this value is set to `true` by default. However, for `AWS::EMR::Cluster` resources in CloudFormation, the default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-visibletoallusers
	//
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

