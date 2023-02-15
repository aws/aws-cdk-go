package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   cfnCluster := awscdk.Aws_emr.NewCfnCluster(this, jsii.String("MyCfnCluster"), &cfnClusterProps{
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
//   	osReleaseLabel: jsii.String("osReleaseLabel"),
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
	// The default role is `EMR_AutoScaling_DefaultRole` . The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
	AutoScalingRole() *string
	SetAutoScalingRole(val *string)
	// `AWS::EMR::Cluster.AutoTerminationPolicy`.
	AutoTerminationPolicy() interface{}
	SetAutoTerminationPolicy(val interface{})
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	BootstrapActions() interface{}
	SetBootstrapActions(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Applies only to Amazon EMR releases 4.x and later. The list of Configurations supplied to the EMR cluster.
	Configurations() interface{}
	SetConfigurations(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Available only in Amazon EMR version 5.7.0 and later. The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	// The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance.
	//
	// Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	// A specification of the number and type of Amazon EC2 instances.
	Instances() interface{}
	SetInstances(val interface{})
	// Also called instance profile and EC2 role.
	//
	// An IAM role for an EMR cluster. The EC2 instances of the cluster assume this role. The default role is `EMR_EC2_DefaultRole` . In order to use the default role, you must have already created it using the CLI or console.
	JobFlowRole() *string
	SetJobFlowRole(val *string)
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	//
	// For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *Amazon EMR Management Guide* .
	KerberosAttributes() interface{}
	SetKerberosAttributes(val interface{})
	// The AWS KMS key used for encrypting log files.
	//
	// This attribute is only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.
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
	LogicalId() *string
	// The path to the Amazon S3 location where logs for this cluster are stored.
	LogUri() *string
	SetLogUri(val *string)
	// Creates or updates a managed scaling policy for an Amazon EMR cluster.
	//
	// The managed scaling policy defines the limits for resources, such as EC2 instances that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	ManagedScalingPolicy() interface{}
	SetManagedScalingPolicy(val interface{})
	// The name of the cluster.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The Amazon Linux release specified in a cluster launch RunJobFlow request.
	//
	// If no Amazon Linux release was specified, the default Amazon Linux release is shown in the response.
	OsReleaseLabel() *string
	SetOsReleaseLabel(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	//
	// Release labels are in the form `emr-x.x.x` , where x.x.x is an Amazon EMR release version such as `emr-5.14.0` . For more information about Amazon EMR release versions and included application versions and features, see [](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/) . The release label applies only to Amazon EMR releases version 4.0 and later. Earlier versions use `AmiVersion` .
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	//
	// `TERMINATE_AT_INSTANCE_HOUR` indicates that Amazon EMR terminates nodes at the instance-hour boundary, regardless of when the request to terminate the instance was submitted. This option is only available with Amazon EMR 5.1.0 and later and is the default for clusters created using that version. `TERMINATE_AT_TASK_COMPLETION` indicates that Amazon EMR adds nodes to a deny list and drains tasks from nodes before terminating the Amazon EC2 instances, regardless of the instance-hour boundary. With either behavior, Amazon EMR removes the least active nodes first and blocks instance termination if it could lead to HDFS corruption. `TERMINATE_AT_TASK_COMPLETION` is available only in Amazon EMR version 4.1.0 and later, and is the default for versions of Amazon EMR earlier than 5.1.0.
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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster.
	//
	// If this value is set to `true` , all IAM users of that AWS account can view and manage the cluster if they have the proper policy permissions set. If this value is `false` , only the IAM user that created the cluster can view and manage it. This value can be changed using the SetVisibleToAllUsers action.
	//
	// > When you create clusters directly through the EMR console or API, this value is set to `true` by default. However, for `AWS::EMR::Cluster` resources in CloudFormation, the default is `false` .
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_CfnCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	if err := validateNewCfnClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnCluster",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
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
		"aws-cdk-lib.aws_emr.CfnCluster",
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

func (c *jsiiProxy_CfnCluster) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnCluster) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnCluster) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnCluster) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

