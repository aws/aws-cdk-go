package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalInfo interface{}
//   var configurationProperty_ configurationProperty
//
//   cfnClusterProps := &cfnClusterProps{
//   	instances: &jobFlowInstancesConfigProperty{
//   		additionalMasterSecurityGroups: []*string{
//   			jsii.String("additionalMasterSecurityGroups"),
//   		},
//   		additionalSlaveSecurityGroups: []*string{
//   			jsii.String("additionalSlaveSecurityGroups"),
//   		},
//   		coreInstanceFleet: &instanceFleetConfigProperty{
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
//   		coreInstanceGroup: &instanceGroupConfigProperty{
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
//   		ec2KeyName: jsii.String("ec2KeyName"),
//   		ec2SubnetId: jsii.String("ec2SubnetId"),
//   		ec2SubnetIds: []*string{
//   			jsii.String("ec2SubnetIds"),
//   		},
//   		emrManagedMasterSecurityGroup: jsii.String("emrManagedMasterSecurityGroup"),
//   		emrManagedSlaveSecurityGroup: jsii.String("emrManagedSlaveSecurityGroup"),
//   		hadoopVersion: jsii.String("hadoopVersion"),
//   		keepJobFlowAliveWhenNoSteps: jsii.Boolean(false),
//   		masterInstanceFleet: &instanceFleetConfigProperty{
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
//   		masterInstanceGroup: &instanceGroupConfigProperty{
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
//   		placement: &placementTypeProperty{
//   			availabilityZone: jsii.String("availabilityZone"),
//   		},
//   		serviceAccessSecurityGroup: jsii.String("serviceAccessSecurityGroup"),
//   		taskInstanceFleets: []interface{}{
//   			&instanceFleetConfigProperty{
//   				instanceTypeConfigs: []interface{}{
//   					&instanceTypeConfigProperty{
//   						instanceType: jsii.String("instanceType"),
//
//   						// the properties below are optional
//   						bidPrice: jsii.String("bidPrice"),
//   						bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   						configurations: []interface{}{
//   							&configurationProperty{
//   								classification: jsii.String("classification"),
//   								configurationProperties: map[string]*string{
//   									"configurationPropertiesKey": jsii.String("configurationProperties"),
//   								},
//   								configurations: []interface{}{
//   									configurationProperty_,
//   								},
//   							},
//   						},
//   						customAmiId: jsii.String("customAmiId"),
//   						ebsConfiguration: &ebsConfigurationProperty{
//   							ebsBlockDeviceConfigs: []interface{}{
//   								&ebsBlockDeviceConfigProperty{
//   									volumeSpecification: &volumeSpecificationProperty{
//   										sizeInGb: jsii.Number(123),
//   										volumeType: jsii.String("volumeType"),
//
//   										// the properties below are optional
//   										iops: jsii.Number(123),
//   									},
//
//   									// the properties below are optional
//   									volumesPerInstance: jsii.Number(123),
//   								},
//   							},
//   							ebsOptimized: jsii.Boolean(false),
//   						},
//   						weightedCapacity: jsii.Number(123),
//   					},
//   				},
//   				launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   					onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   						allocationStrategy: jsii.String("allocationStrategy"),
//   					},
//   					spotSpecification: &spotProvisioningSpecificationProperty{
//   						timeoutAction: jsii.String("timeoutAction"),
//   						timeoutDurationMinutes: jsii.Number(123),
//
//   						// the properties below are optional
//   						allocationStrategy: jsii.String("allocationStrategy"),
//   						blockDurationMinutes: jsii.Number(123),
//   					},
//   				},
//   				name: jsii.String("name"),
//   				targetOnDemandCapacity: jsii.Number(123),
//   				targetSpotCapacity: jsii.Number(123),
//   			},
//   		},
//   		taskInstanceGroups: []interface{}{
//   			&instanceGroupConfigProperty{
//   				instanceCount: jsii.Number(123),
//   				instanceType: jsii.String("instanceType"),
//
//   				// the properties below are optional
//   				autoScalingPolicy: &autoScalingPolicyProperty{
//   					constraints: &scalingConstraintsProperty{
//   						maxCapacity: jsii.Number(123),
//   						minCapacity: jsii.Number(123),
//   					},
//   					rules: []interface{}{
//   						&scalingRuleProperty{
//   							action: &scalingActionProperty{
//   								simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   									scalingAdjustment: jsii.Number(123),
//
//   									// the properties below are optional
//   									adjustmentType: jsii.String("adjustmentType"),
//   									coolDown: jsii.Number(123),
//   								},
//
//   								// the properties below are optional
//   								market: jsii.String("market"),
//   							},
//   							name: jsii.String("name"),
//   							trigger: &scalingTriggerProperty{
//   								cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   									comparisonOperator: jsii.String("comparisonOperator"),
//   									metricName: jsii.String("metricName"),
//   									period: jsii.Number(123),
//   									threshold: jsii.Number(123),
//
//   									// the properties below are optional
//   									dimensions: []interface{}{
//   										&metricDimensionProperty{
//   											key: jsii.String("key"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									evaluationPeriods: jsii.Number(123),
//   									namespace: jsii.String("namespace"),
//   									statistic: jsii.String("statistic"),
//   									unit: jsii.String("unit"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							description: jsii.String("description"),
//   						},
//   					},
//   				},
//   				bidPrice: jsii.String("bidPrice"),
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
//   				market: jsii.String("market"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   		terminationProtected: jsii.Boolean(false),
//   	},
//   	jobFlowRole: jsii.String("jobFlowRole"),
//   	name: jsii.String("name"),
//   	serviceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
//   	additionalInfo: additionalInfo,
//   	applications: []interface{}{
//   		&applicationProperty{
//   			additionalInfo: map[string]*string{
//   				"additionalInfoKey": jsii.String("additionalInfo"),
//   			},
//   			args: []*string{
//   				jsii.String("args"),
//   			},
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	autoScalingRole: jsii.String("autoScalingRole"),
//   	autoTerminationPolicy: &autoTerminationPolicyProperty{
//   		idleTimeout: jsii.Number(123),
//   	},
//   	bootstrapActions: []interface{}{
//   		&bootstrapActionConfigProperty{
//   			name: jsii.String("name"),
//   			scriptBootstrapAction: &scriptBootstrapActionConfigProperty{
//   				path: jsii.String("path"),
//
//   				// the properties below are optional
//   				args: []*string{
//   					jsii.String("args"),
//   				},
//   			},
//   		},
//   	},
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
//   	ebsRootVolumeSize: jsii.Number(123),
//   	kerberosAttributes: &kerberosAttributesProperty{
//   		kdcAdminPassword: jsii.String("kdcAdminPassword"),
//   		realm: jsii.String("realm"),
//
//   		// the properties below are optional
//   		adDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   		adDomainJoinUser: jsii.String("adDomainJoinUser"),
//   		crossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   	},
//   	logEncryptionKmsKeyId: jsii.String("logEncryptionKmsKeyId"),
//   	logUri: jsii.String("logUri"),
//   	managedScalingPolicy: &managedScalingPolicyProperty{
//   		computeLimits: &computeLimitsProperty{
//   			maximumCapacityUnits: jsii.Number(123),
//   			minimumCapacityUnits: jsii.Number(123),
//   			unitType: jsii.String("unitType"),
//
//   			// the properties below are optional
//   			maximumCoreCapacityUnits: jsii.Number(123),
//   			maximumOnDemandCapacityUnits: jsii.Number(123),
//   		},
//   	},
//   	releaseLabel: jsii.String("releaseLabel"),
//   	scaleDownBehavior: jsii.String("scaleDownBehavior"),
//   	securityConfiguration: jsii.String("securityConfiguration"),
//   	stepConcurrencyLevel: jsii.Number(123),
//   	steps: []interface{}{
//   		&stepConfigProperty{
//   			hadoopJarStep: &hadoopJarStepConfigProperty{
//   				jar: jsii.String("jar"),
//
//   				// the properties below are optional
//   				args: []*string{
//   					jsii.String("args"),
//   				},
//   				mainClass: jsii.String("mainClass"),
//   				stepProperties: []interface{}{
//   					&keyValueProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			actionOnFailure: jsii.String("actionOnFailure"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visibleToAllUsers: jsii.Boolean(false),
//   }
//
type CfnClusterProps struct {
	// A specification of the number and type of Amazon EC2 instances.
	Instances interface{} `field:"required" json:"instances" yaml:"instances"`
	// Also called instance profile and EC2 role.
	//
	// An IAM role for an EMR cluster. The EC2 instances of the cluster assume this role. The default role is `EMR_EC2_DefaultRole` . In order to use the default role, you must have already created it using the CLI or console.
	JobFlowRole *string `field:"required" json:"jobFlowRole" yaml:"jobFlowRole"`
	// The name of the cluster.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IAM role that Amazon EMR assumes in order to access AWS resources on your behalf.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// A JSON string for selecting additional features.
	AdditionalInfo interface{} `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// The applications to install on this cluster, for example, Spark, Flink, Oozie, Zeppelin, and so on.
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
	// An IAM role for automatic scaling policies.
	//
	// The default role is `EMR_AutoScaling_DefaultRole` . The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
	AutoScalingRole *string `field:"optional" json:"autoScalingRole" yaml:"autoScalingRole"`
	// `AWS::EMR::Cluster.AutoTerminationPolicy`.
	AutoTerminationPolicy interface{} `field:"optional" json:"autoTerminationPolicy" yaml:"autoTerminationPolicy"`
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	BootstrapActions interface{} `field:"optional" json:"bootstrapActions" yaml:"bootstrapActions"`
	// Applies only to Amazon EMR releases 4.x and later. The list of Configurations supplied to the EMR cluster.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// Available only in Amazon EMR version 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance.
	//
	// Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	//
	// For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *Amazon EMR Management Guide* .
	KerberosAttributes interface{} `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// The AWS KMS key used for encrypting log files.
	//
	// This attribute is only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.
	LogEncryptionKmsKeyId *string `field:"optional" json:"logEncryptionKmsKeyId" yaml:"logEncryptionKmsKeyId"`
	// The path to the Amazon S3 location where logs for this cluster are stored.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// Creates or updates a managed scaling policy for an Amazon EMR cluster.
	//
	// The managed scaling policy defines the limits for resources, such as EC2 instances that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	ManagedScalingPolicy interface{} `field:"optional" json:"managedScalingPolicy" yaml:"managedScalingPolicy"`
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	//
	// Release labels are in the form `emr-x.x.x` , where x.x.x is an Amazon EMR release version such as `emr-5.14.0` . For more information about Amazon EMR release versions and included application versions and features, see [](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/) . The release label applies only to Amazon EMR releases version 4.0 and later. Earlier versions use `AmiVersion` .
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	//
	// `TERMINATE_AT_INSTANCE_HOUR` indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted. This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version. `TERMINATE_AT_TASK_COMPLETION` indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to HDFS corruption. `TERMINATE_AT_TASK_COMPLETION` is available only in Amazon EMR version 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.
	ScaleDownBehavior *string `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// The name of the security configuration applied to the cluster.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Specifies the number of steps that can be executed concurrently.
	//
	// The default value is `1` . The maximum value is `256` .
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// A list of steps to run.
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// A list of tags associated with a cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.
	//
	// If this value is set to `true` , all IAM users of that AWS account can view and manage the cluster if they have the proper policy permissions set. If this value is `false` , only the IAM user that created the cluster can view and manage it. This value can be changed using the SetVisibleToAllUsers action.
	//
	// > When you create clusters directly through the EMR console or API, this value is set to `true` by default. However, for `AWS::EMR::Cluster` resources in CloudFormation, the default is `false` .
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

