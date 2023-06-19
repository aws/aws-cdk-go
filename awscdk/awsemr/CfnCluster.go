package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::EMR::Cluster`.
//
// The `AWS::EMR::Cluster` resource specifies an Amazon EMR cluster. This cluster is a collection of Amazon EC2 instances that run open source big data frameworks and applications to process and analyze vast amounts of data. For more information, see the [Amazon EMR Management Guide](https://docs.aws.amazon.com//emr/latest/ManagementGuide/) .
//
// Amazon EMR now supports launching task instance groups and task instance fleets as part of the `AWS::EMR::Cluster` resource. This can be done by using the `JobFlowInstancesConfig` property type's `TaskInstanceGroups` and `TaskInstanceFleets` subproperties. Using these subproperties reduces delays in provisioning task nodes compared to specifying task nodes with the `AWS::EMR::InstanceGroupConfig` and `AWS::EMR::InstanceFleetConfig` resources. Please refer to the examples at the bottom of this page to learn how to use these subproperties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalInfo interface{}
//   var configurationProperty_ configurationProperty
//
//   cfnCluster := awscdk.Aws_emr.NewCfnCluster(this, jsii.String("MyCfnCluster"), &CfnClusterProps{
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
//   		CoreInstanceGroup: &InstanceGroupConfigProperty{
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
//   		MasterInstanceGroup: &InstanceGroupConfigProperty{
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
//   		Placement: &PlacementTypeProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   		},
//   		ServiceAccessSecurityGroup: jsii.String("serviceAccessSecurityGroup"),
//   		TaskInstanceFleets: []interface{}{
//   			&InstanceFleetConfigProperty{
//   				InstanceTypeConfigs: []interface{}{
//   					&InstanceTypeConfigProperty{
//   						InstanceType: jsii.String("instanceType"),
//
//   						// the properties below are optional
//   						BidPrice: jsii.String("bidPrice"),
//   						BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   						Configurations: []interface{}{
//   							&configurationProperty{
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
//   										SizeInGb: jsii.Number(123),
//   										VolumeType: jsii.String("volumeType"),
//
//   										// the properties below are optional
//   										Iops: jsii.Number(123),
//   									},
//
//   									// the properties below are optional
//   									VolumesPerInstance: jsii.Number(123),
//   								},
//   							},
//   							EbsOptimized: jsii.Boolean(false),
//   						},
//   						WeightedCapacity: jsii.Number(123),
//   					},
//   				},
//   				LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   					OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   					},
//   					SpotSpecification: &SpotProvisioningSpecificationProperty{
//   						TimeoutAction: jsii.String("timeoutAction"),
//   						TimeoutDurationMinutes: jsii.Number(123),
//
//   						// the properties below are optional
//   						AllocationStrategy: jsii.String("allocationStrategy"),
//   						BlockDurationMinutes: jsii.Number(123),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				TargetOnDemandCapacity: jsii.Number(123),
//   				TargetSpotCapacity: jsii.Number(123),
//   			},
//   		},
//   		TaskInstanceGroups: []interface{}{
//   			&InstanceGroupConfigProperty{
//   				InstanceCount: jsii.Number(123),
//   				InstanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				AutoScalingPolicy: &AutoScalingPolicyProperty{
//   					Constraints: &ScalingConstraintsProperty{
//   						MaxCapacity: jsii.Number(123),
//   						MinCapacity: jsii.Number(123),
//   					},
//   					Rules: []interface{}{
//   						&ScalingRuleProperty{
//   							Action: &ScalingActionProperty{
//   								SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   									ScalingAdjustment: jsii.Number(123),
//
//   									// the properties below are optional
//   									AdjustmentType: jsii.String("adjustmentType"),
//   									CoolDown: jsii.Number(123),
//   								},
//
//   								// the properties below are optional
//   								Market: jsii.String("market"),
//   							},
//   							Name: jsii.String("name"),
//   							Trigger: &ScalingTriggerProperty{
//   								CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   									ComparisonOperator: jsii.String("comparisonOperator"),
//   									MetricName: jsii.String("metricName"),
//   									Period: jsii.Number(123),
//   									Threshold: jsii.Number(123),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Key: jsii.String("key"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									EvaluationPeriods: jsii.Number(123),
//   									Namespace: jsii.String("namespace"),
//   									Statistic: jsii.String("statistic"),
//   									Unit: jsii.String("unit"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   						},
//   					},
//   				},
//   				BidPrice: jsii.String("bidPrice"),
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
//   							},
//
//   							// the properties below are optional
//   							VolumesPerInstance: jsii.Number(123),
//   						},
//   					},
//   					EbsOptimized: jsii.Boolean(false),
//   				},
//   				Market: jsii.String("market"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TerminationProtected: jsii.Boolean(false),
//   	},
//   	JobFlowRole: jsii.String("jobFlowRole"),
//   	Name: jsii.String("name"),
//   	ServiceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
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
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   			},
//   		},
//   	},
//   	Configurations: []interface{}{
//   		&configurationProperty{
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
//   	EbsRootVolumeSize: jsii.Number(123),
//   	KerberosAttributes: &KerberosAttributesProperty{
//   		KdcAdminPassword: jsii.String("kdcAdminPassword"),
//   		Realm: jsii.String("realm"),
//
//   		// the properties below are optional
//   		AdDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   		AdDomainJoinUser: jsii.String("adDomainJoinUser"),
//   		CrossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   	},
//   	LogEncryptionKmsKeyId: jsii.String("logEncryptionKmsKeyId"),
//   	LogUri: jsii.String("logUri"),
//   	ManagedScalingPolicy: &ManagedScalingPolicyProperty{
//   		ComputeLimits: &ComputeLimitsProperty{
//   			MaximumCapacityUnits: jsii.Number(123),
//   			MinimumCapacityUnits: jsii.Number(123),
//   			UnitType: jsii.String("unitType"),
//
//   			// the properties below are optional
//   			MaximumCoreCapacityUnits: jsii.Number(123),
//   			MaximumOnDemandCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	OsReleaseLabel: jsii.String("osReleaseLabel"),
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	ScaleDownBehavior: jsii.String("scaleDownBehavior"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	StepConcurrencyLevel: jsii.Number(123),
//   	Steps: []interface{}{
//   		&StepConfigProperty{
//   			HadoopJarStep: &HadoopJarStepConfigProperty{
//   				Jar: jsii.String("jar"),
//
//   				// the properties below are optional
//   				Args: []*string{
//   					jsii.String("args"),
//   				},
//   				MainClass: jsii.String("mainClass"),
//   				StepProperties: []interface{}{
//   					&KeyValueProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ActionOnFailure: jsii.String("actionOnFailure"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VisibleToAllUsers: jsii.Boolean(false),
//   })
//
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A JSON string for selecting additional features.
	AdditionalInfo() interface{}
	SetAdditionalInfo(val interface{})
	// The applications to install on this cluster, for example, Spark, Flink, Oozie, Zeppelin, and so on.
	Applications() interface{}
	SetApplications(val interface{})
	// The public DNS name of the master node (instance), such as `ec2-12-123-123-123.us-west-2.compute.amazonaws.com` .
	AttrMasterPublicDns() *string
	// An IAM role for automatic scaling policies.
	//
	// The default role is `EMR_AutoScaling_DefaultRole` . The IAM role provides permissions that the automatic scaling feature requires to launch and terminate Amazon EC2 instances in an instance group.
	AutoScalingRole() *string
	SetAutoScalingRole(val *string)
	// `AWS::EMR::Cluster.AutoTerminationPolicy`.
	AutoTerminationPolicy() interface{}
	SetAutoTerminationPolicy(val interface{})
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	BootstrapActions() interface{}
	SetBootstrapActions(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Applies only to Amazon EMR releases 4.x and later. The list of configurations that are supplied to the Amazon EMR cluster.
	Configurations() interface{}
	SetConfigurations(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Available only in Amazon EMR releases 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	// The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each Amazon EC2 instance.
	//
	// Available in Amazon EMR releases 4.x and later.
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	// A specification of the number and type of Amazon EC2 instances.
	Instances() interface{}
	SetInstances(val interface{})
	// Also called instance profile and Amazon EC2 role.
	//
	// An IAM role for an Amazon EMR cluster. The Amazon EC2 instances of the cluster assume this role. The default role is `EMR_EC2_DefaultRole` . In order to use the default role, you must have already created it using the AWS CLI or console.
	JobFlowRole() *string
	SetJobFlowRole(val *string)
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	//
	// For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *Amazon EMR Management Guide* .
	KerberosAttributes() interface{}
	SetKerberosAttributes(val interface{})
	// The AWS KMS key used for encrypting log files.
	//
	// This attribute is only available with Amazon EMR 5.30.0 and later, excluding Amazon EMR 6.0.0.
	LogEncryptionKmsKeyId() *string
	SetLogEncryptionKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The path to the Amazon S3 location where logs for this cluster are stored.
	LogUri() *string
	SetLogUri(val *string)
	// Creates or updates a managed scaling policy for an Amazon EMR cluster.
	//
	// The managed scaling policy defines the limits for resources, such as Amazon EC2 instances that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	ManagedScalingPolicy() interface{}
	SetManagedScalingPolicy(val interface{})
	// The name of the cluster.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::EMR::Cluster.OSReleaseLabel`.
	OsReleaseLabel() *string
	SetOsReleaseLabel(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	//
	// Release labels are in the form `emr-x.x.x` , where x.x.x is an Amazon EMR release version such as `emr-5.14.0` . For more information about Amazon EMR release versions and included application versions and features, see [](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/) . The release label applies only to Amazon EMR releases version 4.0 and later. Earlier versions use `AmiVersion` .
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	//
	// `TERMINATE_AT_INSTANCE_HOUR` indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted. This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version. `TERMINATE_AT_TASK_COMPLETION` indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to HDFS corruption. `TERMINATE_AT_TASK_COMPLETION` is available only in Amazon EMR releases 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.
	ScaleDownBehavior() *string
	SetScaleDownBehavior(val *string)
	// The name of the security configuration applied to the cluster.
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	// The IAM role that Amazon EMR assumes in order to access AWS resources on your behalf.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Specifies the number of steps that can be executed concurrently.
	//
	// The default value is `1` . The maximum value is `256` .
	StepConcurrencyLevel() *float64
	SetStepConcurrencyLevel(val *float64)
	// A list of steps to run.
	Steps() interface{}
	SetSteps(val interface{})
	// A list of tags associated with a cluster.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.
	//
	// If this value is set to `true` , all IAM users of that AWS account can view and manage the cluster if they have the proper policy permissions set. If this value is `false` , only the IAM user that created the cluster can view and manage it. This value can be changed using the SetVisibleToAllUsers action.
	//
	// > When you create clusters directly through the EMR console or API, this value is set to `true` by default. However, for `AWS::EMR::Cluster` resources in CloudFormation, the default is `false` .
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AdditionalInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Applications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrMasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMasterPublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AutoScalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AutoTerminationPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) BootstrapActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Configurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Instances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) JobFlowRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KerberosAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ManagedScalingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OsReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osReleaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Steps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Cluster`.
func NewCfnCluster(scope awscdk.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	if err := validateNewCfnClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"monocdk.aws_emr.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope awscdk.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_emr.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster)SetAdditionalInfo(val interface{}) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetApplications(val interface{}) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAutoScalingRole(val *string) {
	_jsii_.Set(
		j,
		"autoScalingRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAutoTerminationPolicy(val interface{}) {
	if err := j.validateSetAutoTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetBootstrapActions(val interface{}) {
	if err := j.validateSetBootstrapActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapActions",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetConfigurations(val interface{}) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEbsRootVolumeSize(val *float64) {
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetInstances(val interface{}) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetJobFlowRole(val *string) {
	if err := j.validateSetJobFlowRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobFlowRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetKerberosAttributes(val interface{}) {
	if err := j.validateSetKerberosAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetLogEncryptionKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetManagedScalingPolicy(val interface{}) {
	if err := j.validateSetManagedScalingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetOsReleaseLabel(val *string) {
	_jsii_.Set(
		j,
		"osReleaseLabel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetReleaseLabel(val *string) {
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetScaleDownBehavior(val *string) {
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetStepConcurrencyLevel(val *float64) {
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSteps(val interface{}) {
	if err := j.validateSetStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetVisibleToAllUsers(val interface{}) {
	if err := j.validateSetVisibleToAllUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleToAllUsers",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_emr.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

