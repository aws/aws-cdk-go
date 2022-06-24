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
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

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

func (j *jsiiProxy_CfnCluster) SetAdditionalInfo(val interface{}) {
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetApplications(val interface{}) {
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAutoScalingRole(val *string) {
	_jsii_.Set(
		j,
		"autoScalingRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetBootstrapActions(val interface{}) {
	_jsii_.Set(
		j,
		"bootstrapActions",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEbsRootVolumeSize(val *float64) {
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetInstances(val interface{}) {
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetJobFlowRole(val *string) {
	_jsii_.Set(
		j,
		"jobFlowRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKerberosAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"kerberosAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLogEncryptionKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetManagedScalingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"managedScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetReleaseLabel(val *string) {
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetScaleDownBehavior(val *string) {
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetStepConcurrencyLevel(val *float64) {
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSteps(val interface{}) {
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetVisibleToAllUsers(val interface{}) {
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
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `Application` is a property of `AWS::EMR::Cluster` .
//
// The `Application` property type defines the open-source big data applications for EMR to install and configure when a cluster is created.
//
// With Amazon EMR release version 4.0 and later, the only accepted parameter is the application `Name` . To pass arguments to these applications, you use configuration classifications specified using JSON objects in a `Configuration` property. For more information, see [Configuring Applications](https://docs.aws.amazon.com//emr/latest/ReleaseGuide/emr-configure-apps.html) .
//
// With earlier Amazon EMR releases, the application is any AWS or third-party software that you can add to the cluster. You can specify the version of the application and arguments to pass to it. Amazon EMR accepts and forwards the argument list to the corresponding installation script as a bootstrap action argument.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationProperty := &applicationProperty{
//   	additionalInfo: map[string]*string{
//   		"additionalInfoKey": jsii.String("additionalInfo"),
//   	},
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnCluster_ApplicationProperty struct {
	// This option is for advanced users only.
	//
	// This is meta information about clusters and applications that are used for testing and troubleshooting.
	AdditionalInfo interface{} `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Arguments for Amazon EMR to pass to the application.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the application.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the application.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
//
// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	constraints: &scalingConstraintsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	rules: []interface{}{
//   		&scalingRuleProperty{
//   			action: &scalingActionProperty{
//   				simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   					scalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					adjustmentType: jsii.String("adjustmentType"),
//   					coolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				market: jsii.String("market"),
//   			},
//   			name: jsii.String("name"),
//   			trigger: &scalingTriggerProperty{
//   				cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					metricName: jsii.String("metricName"),
//   					period: jsii.Number(123),
//   					threshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					dimensions: []interface{}{
//   						&metricDimensionProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					evaluationPeriods: jsii.Number(123),
//   					namespace: jsii.String("namespace"),
//   					statistic: jsii.String("statistic"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   }
//
type CfnCluster_AutoScalingPolicyProperty struct {
	// The upper and lower EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance group to grow above or below these limits.
	Constraints interface{} `field:"required" json:"constraints" yaml:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// `BootstrapActionConfig` is a property of `AWS::EMR::Cluster` that can be used to run bootstrap actions on EMR clusters.
//
// You can use a bootstrap action to install software and configure EC2 instances for all cluster nodes before EMR installs and configures open-source big data applications on cluster instances. For more information, see [Create Bootstrap Actions to Install Additional Software](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-plan-bootstrap.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstrapActionConfigProperty := &bootstrapActionConfigProperty{
//   	name: jsii.String("name"),
//   	scriptBootstrapAction: &scriptBootstrapActionConfigProperty{
//   		path: jsii.String("path"),
//
//   		// the properties below are optional
//   		args: []*string{
//   			jsii.String("args"),
//   		},
//   	},
//   }
//
type CfnCluster_BootstrapActionConfigProperty struct {
	// The name of the bootstrap action.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The script run by the bootstrap action.
	ScriptBootstrapAction interface{} `field:"required" json:"scriptBootstrapAction" yaml:"scriptBootstrapAction"`
}

// `CloudWatchAlarmDefinition` is a subproperty of the `ScalingTrigger` property, which determines when to trigger an automatic scaling activity.
//
// Scaling activity begins when you satisfy the defined alarm conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchAlarmDefinitionProperty := &cloudWatchAlarmDefinitionProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	metricName: jsii.String("metricName"),
//   	period: jsii.Number(123),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&metricDimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluationPeriods: jsii.Number(123),
//   	namespace: jsii.String("namespace"),
//   	statistic: jsii.String("statistic"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnCluster_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by `MetricName` is compared to the value specified by `Threshold` .
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// EMR CloudWatch metrics are emitted every five minutes (300 seconds), so if an EMR CloudWatch metric is specified, specify `300` .
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// A CloudWatch metric dimension.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	//
	// The default value is `1` .
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	//
	// The default is `AWS/ElasticMapReduce` .
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	//
	// The default is `AVERAGE` .
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for `Unit` must correspond to the units specified in the CloudWatch metric.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

// The EC2 unit limits for a managed scaling policy.
//
// The managed scaling activity of a cluster can not be above or below these limits. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeLimitsProperty := &computeLimitsProperty{
//   	maximumCapacityUnits: jsii.Number(123),
//   	minimumCapacityUnits: jsii.Number(123),
//   	unitType: jsii.String("unitType"),
//
//   	// the properties below are optional
//   	maximumCoreCapacityUnits: jsii.Number(123),
//   	maximumOnDemandCapacityUnits: jsii.Number(123),
//   }
//
type CfnCluster_ComputeLimitsProperty struct {
	// The upper boundary of EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. Managed scaling activities are not allowed beyond this boundary. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	MaximumCapacityUnits *float64 `field:"required" json:"maximumCapacityUnits" yaml:"maximumCapacityUnits"`
	// The lower boundary of EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. Managed scaling activities are not allowed beyond this boundary. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	MinimumCapacityUnits *float64 `field:"required" json:"minimumCapacityUnits" yaml:"minimumCapacityUnits"`
	// The unit type used for specifying a managed scaling policy.
	UnitType *string `field:"required" json:"unitType" yaml:"unitType"`
	// The upper boundary of EC2 units for core node type in a cluster.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. The core units are not allowed to scale beyond this boundary. The parameter is used to split capacity allocation between core and task nodes.
	MaximumCoreCapacityUnits *float64 `field:"optional" json:"maximumCoreCapacityUnits" yaml:"maximumCoreCapacityUnits"`
	// The upper boundary of On-Demand EC2 units.
	//
	// It is measured through vCPU cores or instances for instance groups and measured through units for instance fleets. The On-Demand units are not allowed to scale beyond this boundary. The parameter is used to split capacity allocation between On-Demand and Spot Instances.
	MaximumOnDemandCapacityUnits *float64 `field:"optional" json:"maximumOnDemandCapacityUnits" yaml:"maximumOnDemandCapacityUnits"`
}

// > Used only with Amazon EMR release 4.0 and later.
//
// `Configuration` is a subproperty of `InstanceFleetConfig` or `InstanceGroupConfig` . `Configuration` specifies optional configurations for customizing open-source big data applications and environment parameters. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) in the *Amazon EMR Release Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   configurationProperty := &configurationProperty{
//   	classification: jsii.String("classification"),
//   	configurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
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
//   }
//
type CfnCluster_ConfigurationProperty struct {
	// The classification within a configuration.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// A list of additional configurations to apply within a configuration object.
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// A list of additional configurations to apply within a configuration object.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

// `EbsBlockDeviceConfig` is a subproperty of the `EbsConfiguration` property type.
//
// `EbsBlockDeviceConfig` defines the number and type of EBS volumes to associate with all EC2 instances in an EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceConfigProperty := &ebsBlockDeviceConfigProperty{
//   	volumeSpecification: &volumeSpecificationProperty{
//   		sizeInGb: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//
//   		// the properties below are optional
//   		iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	volumesPerInstance: jsii.Number(123),
//   }
//
type CfnCluster_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

// `EbsConfiguration` is a subproperty of `InstanceFleetConfig` or `InstanceGroupConfig` .
//
// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsConfigurationProperty := &ebsConfigurationProperty{
//   	ebsBlockDeviceConfigs: []interface{}{
//   		&ebsBlockDeviceConfigProperty{
//   			volumeSpecification: &volumeSpecificationProperty{
//   				sizeInGb: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//
//   				// the properties below are optional
//   				iops: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			volumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   }
//
type CfnCluster_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}

// The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed.
//
// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hadoopJarStepConfigProperty := &hadoopJarStepConfigProperty{
//   	jar: jsii.String("jar"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	mainClass: jsii.String("mainClass"),
//   	stepProperties: []interface{}{
//   		&keyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCluster_HadoopJarStepConfigProperty struct {
	// A path to a JAR file run during the step.
	Jar *string `field:"required" json:"jar" yaml:"jar"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key-value pairs to your main function.
	StepProperties interface{} `field:"optional" json:"stepProperties" yaml:"stepProperties"`
}

// Use `InstanceFleetConfig` to define instance fleets for an EMR cluster.
//
// A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   instanceFleetConfigProperty := &instanceFleetConfigProperty{
//   	instanceTypeConfigs: []interface{}{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
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
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   		},
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: jsii.String("timeoutAction"),
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
//   }
//
type CfnCluster_InstanceFleetConfigProperty struct {
	// The instance type configurations that define the EC2 instances in the instance fleet.
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

// `InstanceFleetProvisioningSpecification` is a subproperty of `InstanceFleetConfig` .
//
// `InstanceFleetProvisioningSpecification` defines the launch specification for Spot instances in an instance fleet, which determines the defined duration and provisioning timeout behavior for Spot instances.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &instanceFleetProvisioningSpecificationsProperty{
//   	onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	spotSpecification: &spotProvisioningSpecificationProperty{
//   		timeoutAction: jsii.String("timeoutAction"),
//   		timeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		blockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR version 5.12.1 and later.
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// The launch specification for Spot Instances in the fleet, which determines the defined duration, provisioning timeout behavior, and allocation strategy.
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

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

// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// `InstanceTypeConfig` is a sub-property of `InstanceFleetConfig` . `InstanceTypeConfig` determines the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   instanceTypeConfigProperty := &instanceTypeConfigProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	bidPrice: jsii.String("bidPrice"),
//   	bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
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
//   	weightedCapacity: jsii.Number(123),
//   }
//
type CfnCluster_InstanceTypeConfigProperty struct {
	// An EC2 instance type, such as `m3.xlarge` .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each EC2 Spot Instance type as defined by `InstanceType` .
	//
	// Expressed in USD. If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price, for each EC2 Spot Instance as defined by `InstanceType` .
	//
	// Expressed as a number (for example, 20 specifies 20%). If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the instance type.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each instance as defined by `InstanceType` .
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig` .
	//
	// This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

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
	// Specifies whether to lock the cluster to prevent the Amazon EC2 instances from being terminated by API call, user intervention, or in the event of a job-flow error.
	TerminationProtected interface{} `field:"optional" json:"terminationProtected" yaml:"terminationProtected"`
}

// `KerberosAttributes` is a property of the `AWS::EMR::Cluster` resource.
//
// `KerberosAttributes` define the cluster-specific Kerberos configuration when Kerberos authentication is enabled using a security configuration. The cluster-specific configuration must be compatible with the security configuration. For more information see [Use Kerberos Authentication](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in the *EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kerberosAttributesProperty := &kerberosAttributesProperty{
//   	kdcAdminPassword: jsii.String("kdcAdminPassword"),
//   	realm: jsii.String("realm"),
//
//   	// the properties below are optional
//   	adDomainJoinPassword: jsii.String("adDomainJoinPassword"),
//   	adDomainJoinUser: jsii.String("adDomainJoinUser"),
//   	crossRealmTrustPrincipalPassword: jsii.String("crossRealmTrustPrincipalPassword"),
//   }
//
type CfnCluster_KerberosAttributesProperty struct {
	// The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster.
	KdcAdminPassword *string `field:"required" json:"kdcAdminPassword" yaml:"kdcAdminPassword"`
	// The name of the Kerberos realm to which all nodes in a cluster belong.
	//
	// For example, `EC2.INTERNAL` .
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// The Active Directory password for `ADDomainJoinUser` .
	AdDomainJoinPassword *string `field:"optional" json:"adDomainJoinPassword" yaml:"adDomainJoinPassword"`
	// Required only when establishing a cross-realm trust with an Active Directory domain.
	//
	// A user with sufficient privileges to join resources to the domain.
	AdDomainJoinUser *string `field:"optional" json:"adDomainJoinUser" yaml:"adDomainJoinUser"`
	// Required only when establishing a cross-realm trust with a KDC in a different realm.
	//
	// The cross-realm principal password, which must be identical across realms.
	CrossRealmTrustPrincipalPassword *string `field:"optional" json:"crossRealmTrustPrincipalPassword" yaml:"crossRealmTrustPrincipalPassword"`
}

// `KeyValue` is a subproperty of the `HadoopJarStepConfig` property type.
//
// `KeyValue` is used to pass parameters to a step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueProperty := &keyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnCluster_KeyValueProperty struct {
	// The unique identifier of a key-value pair.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value part of the identified key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Managed scaling policy for an Amazon EMR cluster.
//
// The policy specifies the limits for resources that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingPolicyProperty := &managedScalingPolicyProperty{
//   	computeLimits: &computeLimitsProperty{
//   		maximumCapacityUnits: jsii.Number(123),
//   		minimumCapacityUnits: jsii.Number(123),
//   		unitType: jsii.String("unitType"),
//
//   		// the properties below are optional
//   		maximumCoreCapacityUnits: jsii.Number(123),
//   		maximumOnDemandCapacityUnits: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_ManagedScalingPolicyProperty struct {
	// The EC2 unit limits for a managed scaling policy.
	//
	// The managed scaling activity of a cluster is not allowed to go above or below these limits. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	ComputeLimits interface{} `field:"optional" json:"computeLimits" yaml:"computeLimits"`
}

// `MetricDimension` is a subproperty of the `CloudWatchAlarmDefinition` property type.
//
// `MetricDimension` specifies a CloudWatch dimension, which is specified with a `Key` `Value` pair. The key is known as a `Name` in CloudWatch. By default, Amazon EMR uses one dimension whose `Key` is `JobFlowID` and `Value` is a variable representing the cluster ID, which is `${emr.clusterId}` . This enables the automatic scaling rule for EMR to bootstrap when the cluster ID becomes available during cluster creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnCluster_MetricDimensionProperty struct {
	// The dimension name.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The dimension value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR version 5.12.1 and later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onDemandProvisioningSpecificationProperty := &onDemandProvisioningSpecificationProperty{
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   }
//
type CfnCluster_OnDemandProvisioningSpecificationProperty struct {
	// Specifies the strategy to use in launching On-Demand instance fleets.
	//
	// Currently, the only option is `lowest-price` (the default), which launches the lowest price first.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
}

// `PlacementType` is a property of the `AWS::EMR::Cluster` resource.
//
// `PlacementType` determines the Amazon EC2 Availability Zone configuration of the cluster (job flow).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementTypeProperty := &placementTypeProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   }
//
type CfnCluster_PlacementTypeProperty struct {
	// The Amazon EC2 Availability Zone for the cluster.
	//
	// `AvailabilityZone` is used for uniform instance groups, while `AvailabilityZones` (plural) is used for instance fleets.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
}

// `ScalingAction` is a subproperty of the `ScalingRule` property type.
//
// `ScalingAction` determines the type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingActionProperty := &scalingActionProperty{
//   	simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   		scalingAdjustment: jsii.Number(123),
//
//   		// the properties below are optional
//   		adjustmentType: jsii.String("adjustmentType"),
//   		coolDown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	market: jsii.String("market"),
//   }
//
type CfnCluster_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	SimpleScalingPolicyConfiguration interface{} `field:"required" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	Market *string `field:"optional" json:"market" yaml:"market"`
}

// `ScalingConstraints` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingConstraints` defines the upper and lower EC2 instance limits for an automatic scaling policy. Automatic scaling activities triggered by automatic scaling rules will not cause an instance group to grow above or shrink below these limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConstraintsProperty := &scalingConstraintsProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
type CfnCluster_ScalingConstraintsProperty struct {
	// The upper boundary of EC2 instances in an instance group beyond which scaling activities are not allowed to grow.
	//
	// Scale-out activities will not add instances beyond this boundary.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The lower boundary of EC2 instances in an instance group below which scaling activities are not allowed to shrink.
	//
	// Scale-in activities will not terminate instances below this boundary.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

// `ScalingRule` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingRule` defines the scale-in or scale-out rules for scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments. The automatic scaling policy for an instance group can comprise one or more automatic scaling rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingRuleProperty := &scalingRuleProperty{
//   	action: &scalingActionProperty{
//   		simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			adjustmentType: jsii.String("adjustmentType"),
//   			coolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		market: jsii.String("market"),
//   	},
//   	name: jsii.String("name"),
//   	trigger: &scalingTriggerProperty{
//   		cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			metricName: jsii.String("metricName"),
//   			period: jsii.Number(123),
//   			threshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			evaluationPeriods: jsii.Number(123),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnCluster_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// `ScalingTrigger` is a subproperty of the `ScalingRule` property type.
//
// `ScalingTrigger` determines the conditions that trigger an automatic scaling activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingTriggerProperty := &scalingTriggerProperty{
//   	cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		metricName: jsii.String("metricName"),
//   		period: jsii.Number(123),
//   		threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		evaluationPeriods: jsii.Number(123),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnCluster_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
	CloudWatchAlarmDefinition interface{} `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}

// `ScriptBootstrapActionConfig` is a subproperty of the `BootstrapActionConfig` property type.
//
// `ScriptBootstrapActionConfig` specifies the arguments and location of the bootstrap script for EMR to run on all cluster nodes before it installs open-source big data applications on them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptBootstrapActionConfigProperty := &scriptBootstrapActionConfigProperty{
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   }
//
type CfnCluster_ScriptBootstrapActionConfigProperty struct {
	// Location in Amazon S3 of the script to run during a bootstrap action.
	Path *string `field:"required" json:"path" yaml:"path"`
	// A list of command line arguments to pass to the bootstrap action script.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

// `SimpleScalingPolicyConfiguration` is a subproperty of the `ScalingAction` property type.
//
// `SimpleScalingPolicyConfiguration` determines how an automatic scaling action adds or removes instances, the cooldown period, and the number of EC2 instances that are added each time the CloudWatch metric alarm condition is satisfied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleScalingPolicyConfigurationProperty := &simpleScalingPolicyConfigurationProperty{
//   	scalingAdjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	adjustmentType: jsii.String("adjustmentType"),
//   	coolDown: jsii.Number(123),
//   }
//
type CfnCluster_SimpleScalingPolicyConfigurationProperty struct {
	// The amount by which to scale in or scale out, based on the specified `AdjustmentType` .
	//
	// A positive value adds to the instance group's EC2 instance count while a negative number removes instances. If `AdjustmentType` is set to `EXACT_CAPACITY` , the number should only be a positive integer. If `AdjustmentType` is set to `PERCENT_CHANGE_IN_CAPACITY` , the value should express the percentage as an integer. For example, -20 indicates a decrease in 20% increments of cluster capacity.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The way in which EC2 instances are added (if `ScalingAdjustment` is a positive number) or terminated (if `ScalingAdjustment` is a negative number) each time the scaling activity is triggered.
	//
	// `CHANGE_IN_CAPACITY` is the default. `CHANGE_IN_CAPACITY` indicates that the EC2 instance count increments or decrements by `ScalingAdjustment` , which should be expressed as an integer. `PERCENT_CHANGE_IN_CAPACITY` indicates the instance count increments or decrements by the percentage specified by `ScalingAdjustment` , which should be expressed as an integer. For example, 20 indicates an increase in 20% increments of cluster capacity. `EXACT_CAPACITY` indicates the scaling activity results in an instance group with the number of EC2 instances specified by `ScalingAdjustment` , which should be expressed as a positive integer.
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
	//
	// The default value is 0.
	CoolDown *float64 `field:"optional" json:"coolDown" yaml:"coolDown"`
}

// `SpotProvisioningSpecification` is a subproperty of the `InstanceFleetProvisioningSpecifications` property type.
//
// `SpotProvisioningSpecification` determines the launch specification for Spot instances in the instance fleet, which includes the defined duration and provisioning timeout behavior.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotProvisioningSpecificationProperty := &spotProvisioningSpecificationProperty{
//   	timeoutAction: jsii.String("timeoutAction"),
//   	timeoutDurationMinutes: jsii.Number(123),
//
//   	// the properties below are optional
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	blockDurationMinutes: jsii.Number(123),
//   }
//
type CfnCluster_SpotProvisioningSpecificationProperty struct {
	// The action to take when `TargetSpotCapacity` has not been fulfilled when the `TimeoutDurationMinutes` has expired;
	//
	// that is, when all Spot Instances could not be provisioned within the Spot provisioning timeout. Valid values are `TERMINATE_CLUSTER` and `SWITCH_TO_ON_DEMAND` . SWITCH_TO_ON_DEMAND specifies that if no Spot Instances are available, On-Demand Instances should be provisioned to fulfill any remaining Spot capacity.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// The spot provisioning timeout period in minutes.
	//
	// If Spot Instances are not provisioned within this time period, the `TimeOutAction` is taken. Minimum value is 5 and maximum value is 1440. The timeout applies only during initial provisioning, when the cluster is first created.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Specifies the strategy to use in launching Spot Instance fleets.
	//
	// Currently, the only option is capacity-optimized (the default), which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The defined duration for Spot Instances (also known as Spot blocks) in minutes.
	//
	// When specified, the Spot Instance does not terminate before the defined duration expires, and defined duration pricing for Spot Instances applies. Valid values are 60, 120, 180, 240, 300, or 360. The duration period starts as soon as a Spot Instance receives its instance ID. At the end of the duration, Amazon EC2 marks the Spot Instance for termination and provides a Spot Instance termination notice, which gives the instance a two-minute warning before it terminates.
	//
	// > Spot Instances with a defined duration (also known as Spot blocks) are no longer available to new customers from July 1, 2021. For customers who have previously used the feature, we will continue to support Spot Instances with a defined duration until December 31, 2022.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}

// `StepConfig` is a property of the `AWS::EMR::Cluster` resource.
//
// The `StepConfig` property type specifies a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepConfigProperty := &stepConfigProperty{
//   	hadoopJarStep: &hadoopJarStepConfigProperty{
//   		jar: jsii.String("jar"),
//
//   		// the properties below are optional
//   		args: []*string{
//   			jsii.String("args"),
//   		},
//   		mainClass: jsii.String("mainClass"),
//   		stepProperties: []interface{}{
//   			&keyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	actionOnFailure: jsii.String("actionOnFailure"),
//   }
//
type CfnCluster_StepConfigProperty struct {
	// The JAR file used for the step.
	HadoopJarStep interface{} `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// The name of the step.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
}

// `VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type.
//
// `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeSpecificationProperty := &volumeSpecificationProperty{
//   	sizeInGb: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//
//   	// the properties below are optional
//   	iops: jsii.Number(123),
//   }
//
type CfnCluster_VolumeSpecificationProperty struct {
	// The volume size, in gibibytes (GiB).
	//
	// This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, and standard.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
}

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

// A CloudFormation `AWS::EMR::InstanceFleetConfig`.
//
// Use `InstanceFleetConfig` to define instance fleets for an EMR cluster. A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. > You can currently only add a task instance fleet to a cluster with this resource. If you use this resource, CloudFormation waits for the cluster launch to complete before adding the task instance fleet to the cluster. In order to add a task instance fleet to the cluster as part of the cluster launch and minimize delays in provisioning task nodes, use the `TaskInstanceFleets` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceFleetConfig := awscdk.Aws_emr.NewCfnInstanceFleetConfig(this, jsii.String("MyCfnInstanceFleetConfig"), &cfnInstanceFleetConfigProps{
//   	clusterId: jsii.String("clusterId"),
//   	instanceFleetType: jsii.String("instanceFleetType"),
//
//   	// the properties below are optional
//   	instanceTypeConfigs: []interface{}{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
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
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   		},
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: jsii.String("timeoutAction"),
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
//   })
//
type CfnInstanceFleetConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The unique identifier of the EMR cluster.
	ClusterId() *string
	SetClusterId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The node type that the instance fleet hosts.
	//
	// *Allowed Values* : TASK.
	InstanceFleetType() *string
	SetInstanceFleetType(val *string)
	// `InstanceTypeConfigs` determine the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	InstanceTypeConfigs() interface{}
	SetInstanceTypeConfigs(val interface{})
	// The launch specification for the instance fleet.
	LaunchSpecifications() interface{}
	SetLaunchSpecifications(val interface{})
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
	// The friendly name of the instance fleet.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetOnDemandCapacity() *float64
	SetTargetOnDemandCapacity(val *float64)
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetSpotCapacity() *float64
	SetTargetSpotCapacity(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnInstanceFleetConfig
type jsiiProxy_CfnInstanceFleetConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) InstanceFleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceFleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) InstanceTypeConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) LaunchSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfig(scope constructs.Construct, id *string, props *CfnInstanceFleetConfigProps) CfnInstanceFleetConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceFleetConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfig_Override(c CfnInstanceFleetConfig, scope constructs.Construct, id *string, props *CfnInstanceFleetConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetInstanceFleetType(val *string) {
	_jsii_.Set(
		j,
		"instanceFleetType",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetInstanceTypeConfigs(val interface{}) {
	_jsii_.Set(
		j,
		"instanceTypeConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetLaunchSpecifications(val interface{}) {
	_jsii_.Set(
		j,
		"launchSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetTargetOnDemandCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetTargetSpotCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInstanceFleetConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstanceFleetConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
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
func CfnInstanceFleetConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceFleetConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceFleetConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// > Used only with Amazon EMR release 4.0 and later.
//
// `Configuration` specifies optional configurations for customizing open-source big data applications and environment parameters. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) in the *Amazon EMR Release Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   configurationProperty := &configurationProperty{
//   	classification: jsii.String("classification"),
//   	configurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
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
//   }
//
type CfnInstanceFleetConfig_ConfigurationProperty struct {
	// The classification within a configuration.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Within a configuration classification, a set of properties that represent the settings that you want to change in the configuration file.
	//
	// Duplicates not allowed.
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// A list of additional configurations to apply within a configuration object.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

// `EbsBlockDeviceConfig` is a subproperty of the `EbsConfiguration` property type.
//
// `EbsBlockDeviceConfig` defines the number and type of EBS volumes to associate with all EC2 instances in an EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceConfigProperty := &ebsBlockDeviceConfigProperty{
//   	volumeSpecification: &volumeSpecificationProperty{
//   		sizeInGb: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//
//   		// the properties below are optional
//   		iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	volumesPerInstance: jsii.Number(123),
//   }
//
type CfnInstanceFleetConfig_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsConfigurationProperty := &ebsConfigurationProperty{
//   	ebsBlockDeviceConfigs: []interface{}{
//   		&ebsBlockDeviceConfigProperty{
//   			volumeSpecification: &volumeSpecificationProperty{
//   				sizeInGb: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//
//   				// the properties below are optional
//   				iops: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			volumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   }
//
type CfnInstanceFleetConfig_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}

// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// `InstanceTypeConfig` is a sub-property of `InstanceFleetConfig` . `InstanceTypeConfig` determines the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &instanceFleetProvisioningSpecificationsProperty{
//   	onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	spotSpecification: &spotProvisioningSpecificationProperty{
//   		timeoutAction: jsii.String("timeoutAction"),
//   		timeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		blockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnInstanceFleetConfig_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR version 5.12.1 and later.
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// The launch specification for Spot Instances in the fleet, which determines the defined duration, provisioning timeout behavior, and allocation strategy.
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

// `InstanceType` config is a subproperty of `InstanceFleetConfig` .
//
// An instance type configuration specifies each instance type in an instance fleet. The configuration determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   instanceTypeConfigProperty := &instanceTypeConfigProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	bidPrice: jsii.String("bidPrice"),
//   	bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
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
//   	weightedCapacity: jsii.Number(123),
//   }
//
type CfnInstanceFleetConfig_InstanceTypeConfigProperty struct {
	// An EC2 instance type, such as `m3.xlarge` .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each EC2 Spot Instance type as defined by `InstanceType` .
	//
	// Expressed in USD. If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price, for each EC2 Spot Instance as defined by `InstanceType` .
	//
	// Expressed as a number (for example, 20 specifies 20%). If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// > Amazon EMR releases 4.x or later.
	//
	// An optional configuration specification to be used when provisioning cluster instances, which can include configurations for applications and software bundled with Amazon EMR. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) .
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the instance type.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each instance as defined by `InstanceType` .
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig` .
	//
	// This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR version 5.12.1 and later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onDemandProvisioningSpecificationProperty := &onDemandProvisioningSpecificationProperty{
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   }
//
type CfnInstanceFleetConfig_OnDemandProvisioningSpecificationProperty struct {
	// Specifies the strategy to use in launching On-Demand instance fleets.
	//
	// Currently, the only option is `lowest-price` (the default), which launches the lowest price first.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
}

// `SpotProvisioningSpecification` is a subproperty of the `InstanceFleetProvisioningSpecifications` property type.
//
// `SpotProvisioningSpecification` determines the launch specification for Spot instances in the instance fleet, which includes the defined duration and provisioning timeout behavior.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotProvisioningSpecificationProperty := &spotProvisioningSpecificationProperty{
//   	timeoutAction: jsii.String("timeoutAction"),
//   	timeoutDurationMinutes: jsii.Number(123),
//
//   	// the properties below are optional
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	blockDurationMinutes: jsii.Number(123),
//   }
//
type CfnInstanceFleetConfig_SpotProvisioningSpecificationProperty struct {
	// The action to take when `TargetSpotCapacity` has not been fulfilled when the `TimeoutDurationMinutes` has expired;
	//
	// that is, when all Spot Instances could not be provisioned within the Spot provisioning timeout. Valid values are `TERMINATE_CLUSTER` and `SWITCH_TO_ON_DEMAND` . SWITCH_TO_ON_DEMAND specifies that if no Spot Instances are available, On-Demand Instances should be provisioned to fulfill any remaining Spot capacity.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// The spot provisioning timeout period in minutes.
	//
	// If Spot Instances are not provisioned within this time period, the `TimeOutAction` is taken. Minimum value is 5 and maximum value is 1440. The timeout applies only during initial provisioning, when the cluster is first created.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Specifies the strategy to use in launching Spot Instance fleets.
	//
	// Currently, the only option is capacity-optimized (the default), which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// The defined duration for Spot Instances (also known as Spot blocks) in minutes.
	//
	// When specified, the Spot Instance does not terminate before the defined duration expires, and defined duration pricing for Spot Instances applies. Valid values are 60, 120, 180, 240, 300, or 360. The duration period starts as soon as a Spot Instance receives its instance ID. At the end of the duration, Amazon EC2 marks the Spot Instance for termination and provides a Spot Instance termination notice, which gives the instance a two-minute warning before it terminates.
	//
	// > Spot Instances with a defined duration (also known as Spot blocks) are no longer available to new customers from July 1, 2021. For customers who have previously used the feature, we will continue to support Spot Instances with a defined duration until December 31, 2022.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}

// `VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type.
//
// `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeSpecificationProperty := &volumeSpecificationProperty{
//   	sizeInGb: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//
//   	// the properties below are optional
//   	iops: jsii.Number(123),
//   }
//
type CfnInstanceFleetConfig_VolumeSpecificationProperty struct {
	// The volume size, in gibibytes (GiB).
	//
	// This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, and standard.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
}

// Properties for defining a `CfnInstanceFleetConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceFleetConfigProps := &cfnInstanceFleetConfigProps{
//   	clusterId: jsii.String("clusterId"),
//   	instanceFleetType: jsii.String("instanceFleetType"),
//
//   	// the properties below are optional
//   	instanceTypeConfigs: []interface{}{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
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
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   		},
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: jsii.String("timeoutAction"),
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
//   }
//
type CfnInstanceFleetConfigProps struct {
	// The unique identifier of the EMR cluster.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The node type that the instance fleet hosts.
	//
	// *Allowed Values* : TASK.
	InstanceFleetType *string `field:"required" json:"instanceFleetType" yaml:"instanceFleetType"`
	// `InstanceTypeConfigs` determine the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

// A CloudFormation `AWS::EMR::InstanceGroupConfig`.
//
// Use `InstanceGroupConfig` to define instance groups for an EMR cluster. A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > You can currently only add task instance groups to a cluster with this resource. If you use this resource, CloudFormation waits for the cluster launch to complete before adding the task instance group to the cluster. In order to add task instance groups to the cluster as part of the cluster launch and minimize delays in provisioning task nodes, use the `TaskInstanceGroups` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceGroupConfig := awscdk.Aws_emr.NewCfnInstanceGroupConfig(this, jsii.String("MyCfnInstanceGroupConfig"), &cfnInstanceGroupConfigProps{
//   	instanceCount: jsii.Number(123),
//   	instanceRole: jsii.String("instanceRole"),
//   	instanceType: jsii.String("instanceType"),
//   	jobFlowId: jsii.String("jobFlowId"),
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
//   })
//
type CfnInstanceGroupConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
	//
	// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
	AutoScalingPolicy() interface{}
	SetAutoScalingPolicy(val interface{})
	// If specified, indicates that the instance group uses Spot Instances.
	//
	// This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.
	BidPrice() *string
	SetBidPrice(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// > Amazon EMR releases 4.x or later.
	//
	// The list of configurations supplied for an EMR cluster instance group. You can specify a separate configuration for each instance group (master, core, and task).
	Configurations() interface{}
	SetConfigurations(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The custom AMI ID to use for the provisioned instance group.
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
	EbsConfiguration() interface{}
	SetEbsConfiguration(val interface{})
	// Target number of instances for the instance group.
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	// The role of the instance group in the cluster.
	//
	// *Allowed Values* : TASK.
	InstanceRole() *string
	SetInstanceRole(val *string)
	// The EC2 instance type for all instances in the instance group.
	InstanceType() *string
	SetInstanceType(val *string)
	// The ID of an Amazon EMR cluster that you want to associate this instance group with.
	JobFlowId() *string
	SetJobFlowId(val *string)
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
	// Market type of the EC2 instances used to create a cluster node.
	Market() *string
	SetMarket(val *string)
	// Friendly name given to the instance group.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnInstanceGroupConfig
type jsiiProxy_CfnInstanceGroupConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstanceGroupConfig) AutoScalingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Configurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) EbsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) JobFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Market() *string {
	var returns *string
	_jsii_.Get(
		j,
		"market",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfig(scope constructs.Construct, id *string, props *CfnInstanceGroupConfigProps) CfnInstanceGroupConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceGroupConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfig_Override(c CfnInstanceGroupConfig, scope constructs.Construct, id *string, props *CfnInstanceGroupConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetAutoScalingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"autoScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetEbsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"ebsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceRole(val *string) {
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetJobFlowId(val *string) {
	_jsii_.Set(
		j,
		"jobFlowId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetMarket(val *string) {
	_jsii_.Set(
		j,
		"market",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInstanceGroupConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstanceGroupConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
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
func CfnInstanceGroupConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceGroupConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric.
//
// For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	constraints: &scalingConstraintsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	rules: []interface{}{
//   		&scalingRuleProperty{
//   			action: &scalingActionProperty{
//   				simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   					scalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					adjustmentType: jsii.String("adjustmentType"),
//   					coolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				market: jsii.String("market"),
//   			},
//   			name: jsii.String("name"),
//   			trigger: &scalingTriggerProperty{
//   				cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					metricName: jsii.String("metricName"),
//   					period: jsii.Number(123),
//   					threshold: jsii.Number(123),
//
//   					// the properties below are optional
//   					dimensions: []interface{}{
//   						&metricDimensionProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					evaluationPeriods: jsii.Number(123),
//   					namespace: jsii.String("namespace"),
//   					statistic: jsii.String("statistic"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   }
//
type CfnInstanceGroupConfig_AutoScalingPolicyProperty struct {
	// The upper and lower EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance group to grow above or below these limits.
	Constraints interface{} `field:"required" json:"constraints" yaml:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// `CloudWatchAlarmDefinition` is a subproperty of the `ScalingTrigger` property, which determines when to trigger an automatic scaling activity.
//
// Scaling activity begins when you satisfy the defined alarm conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchAlarmDefinitionProperty := &cloudWatchAlarmDefinitionProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	metricName: jsii.String("metricName"),
//   	period: jsii.Number(123),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&metricDimensionProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluationPeriods: jsii.Number(123),
//   	namespace: jsii.String("namespace"),
//   	statistic: jsii.String("statistic"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnInstanceGroupConfig_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by `MetricName` is compared to the value specified by `Threshold` .
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// EMR CloudWatch metrics are emitted every five minutes (300 seconds), so if an EMR CloudWatch metric is specified, specify `300` .
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// A CloudWatch metric dimension.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	//
	// The default value is `1` .
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	//
	// The default is `AWS/ElasticMapReduce` .
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	//
	// The default is `AVERAGE` .
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for `Unit` must correspond to the units specified in the CloudWatch metric.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

// `Configurations` is a property of the `AWS::EMR::Cluster` resource that specifies the configuration of applications on an Amazon EMR cluster.
//
// Configurations are optional. You can use them to have EMR customize applications and software bundled with Amazon EMR when a cluster is created. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) .
//
// > Applies only to Amazon EMR releases 4.0 and later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   configurationProperty := &configurationProperty{
//   	classification: jsii.String("classification"),
//   	configurationProperties: map[string]*string{
//   		"configurationPropertiesKey": jsii.String("configurationProperties"),
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
//   }
//
type CfnInstanceGroupConfig_ConfigurationProperty struct {
	// The classification within a configuration.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Within a configuration classification, a set of properties that represent the settings that you want to change in the configuration file.
	//
	// Duplicates not allowed.
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// A list of additional configurations to apply within a configuration object.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

// Configuration of requested EBS block device associated with the instance group with count of volumes that will be associated to every instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceConfigProperty := &ebsBlockDeviceConfigProperty{
//   	volumeSpecification: &volumeSpecificationProperty{
//   		sizeInGb: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//
//   		// the properties below are optional
//   		iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	volumesPerInstance: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}

// The Amazon EBS configuration of a cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsConfigurationProperty := &ebsConfigurationProperty{
//   	ebsBlockDeviceConfigs: []interface{}{
//   		&ebsBlockDeviceConfigProperty{
//   			volumeSpecification: &volumeSpecificationProperty{
//   				sizeInGb: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//
//   				// the properties below are optional
//   				iops: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			volumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   }
//
type CfnInstanceGroupConfig_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}

// `MetricDimension` is a subproperty of the `CloudWatchAlarmDefinition` property type.
//
// `MetricDimension` specifies a CloudWatch dimension, which is specified with a `Key` `Value` pair. The key is known as a `Name` in CloudWatch. By default, Amazon EMR uses one dimension whose `Key` is `JobFlowID` and `Value` is a variable representing the cluster ID, which is `${emr.clusterId}` . This enables the automatic scaling rule for EMR to bootstrap when the cluster ID becomes available during cluster creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnInstanceGroupConfig_MetricDimensionProperty struct {
	// The dimension name.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The dimension value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// `ScalingAction` is a subproperty of the `ScalingRule` property type.
//
// `ScalingAction` determines the type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingActionProperty := &scalingActionProperty{
//   	simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   		scalingAdjustment: jsii.Number(123),
//
//   		// the properties below are optional
//   		adjustmentType: jsii.String("adjustmentType"),
//   		coolDown: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	market: jsii.String("market"),
//   }
//
type CfnInstanceGroupConfig_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	SimpleScalingPolicyConfiguration interface{} `field:"required" json:"simpleScalingPolicyConfiguration" yaml:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	Market *string `field:"optional" json:"market" yaml:"market"`
}

// `ScalingConstraints` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingConstraints` defines the upper and lower EC2 instance limits for an automatic scaling policy. Automatic scaling activities triggered by automatic scaling rules will not cause an instance group to grow above or shrink below these limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConstraintsProperty := &scalingConstraintsProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_ScalingConstraintsProperty struct {
	// The upper boundary of EC2 instances in an instance group beyond which scaling activities are not allowed to grow.
	//
	// Scale-out activities will not add instances beyond this boundary.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The lower boundary of EC2 instances in an instance group below which scaling activities are not allowed to shrink.
	//
	// Scale-in activities will not terminate instances below this boundary.
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
}

// `ScalingRule` is a subproperty of the `AutoScalingPolicy` property type.
//
// `ScalingRule` defines the scale-in or scale-out rules for scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments. The automatic scaling policy for an instance group can comprise one or more automatic scaling rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingRuleProperty := &scalingRuleProperty{
//   	action: &scalingActionProperty{
//   		simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			adjustmentType: jsii.String("adjustmentType"),
//   			coolDown: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		market: jsii.String("market"),
//   	},
//   	name: jsii.String("name"),
//   	trigger: &scalingTriggerProperty{
//   		cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			metricName: jsii.String("metricName"),
//   			period: jsii.Number(123),
//   			threshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			evaluationPeriods: jsii.Number(123),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnInstanceGroupConfig_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	Trigger interface{} `field:"required" json:"trigger" yaml:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// `ScalingTrigger` is a subproperty of the `ScalingRule` property type.
//
// `ScalingTrigger` determines the conditions that trigger an automatic scaling activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingTriggerProperty := &scalingTriggerProperty{
//   	cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		metricName: jsii.String("metricName"),
//   		period: jsii.Number(123),
//   		threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		evaluationPeriods: jsii.Number(123),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnInstanceGroupConfig_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
	CloudWatchAlarmDefinition interface{} `field:"required" json:"cloudWatchAlarmDefinition" yaml:"cloudWatchAlarmDefinition"`
}

// `SimpleScalingPolicyConfiguration` is a subproperty of the `ScalingAction` property type.
//
// `SimpleScalingPolicyConfiguration` determines how an automatic scaling action adds or removes instances, the cooldown period, and the number of EC2 instances that are added each time the CloudWatch metric alarm condition is satisfied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   simpleScalingPolicyConfigurationProperty := &simpleScalingPolicyConfigurationProperty{
//   	scalingAdjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	adjustmentType: jsii.String("adjustmentType"),
//   	coolDown: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_SimpleScalingPolicyConfigurationProperty struct {
	// The amount by which to scale in or scale out, based on the specified `AdjustmentType` .
	//
	// A positive value adds to the instance group's EC2 instance count while a negative number removes instances. If `AdjustmentType` is set to `EXACT_CAPACITY` , the number should only be a positive integer. If `AdjustmentType` is set to `PERCENT_CHANGE_IN_CAPACITY` , the value should express the percentage as an integer. For example, -20 indicates a decrease in 20% increments of cluster capacity.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The way in which EC2 instances are added (if `ScalingAdjustment` is a positive number) or terminated (if `ScalingAdjustment` is a negative number) each time the scaling activity is triggered.
	//
	// `CHANGE_IN_CAPACITY` is the default. `CHANGE_IN_CAPACITY` indicates that the EC2 instance count increments or decrements by `ScalingAdjustment` , which should be expressed as an integer. `PERCENT_CHANGE_IN_CAPACITY` indicates the instance count increments or decrements by the percentage specified by `ScalingAdjustment` , which should be expressed as an integer. For example, 20 indicates an increase in 20% increments of cluster capacity. `EXACT_CAPACITY` indicates the scaling activity results in an instance group with the number of EC2 instances specified by `ScalingAdjustment` , which should be expressed as a positive integer.
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
	//
	// The default value is 0.
	CoolDown *float64 `field:"optional" json:"coolDown" yaml:"coolDown"`
}

// `VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type.
//
// `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeSpecificationProperty := &volumeSpecificationProperty{
//   	sizeInGb: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//
//   	// the properties below are optional
//   	iops: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_VolumeSpecificationProperty struct {
	// The volume size, in gibibytes (GiB).
	//
	// This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, and standard.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
}

// Properties for defining a `CfnInstanceGroupConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceGroupConfigProps := &cfnInstanceGroupConfigProps{
//   	instanceCount: jsii.Number(123),
//   	instanceRole: jsii.String("instanceRole"),
//   	instanceType: jsii.String("instanceType"),
//   	jobFlowId: jsii.String("jobFlowId"),
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
type CfnInstanceGroupConfigProps struct {
	// Target number of instances for the instance group.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The role of the instance group in the cluster.
	//
	// *Allowed Values* : TASK.
	InstanceRole *string `field:"required" json:"instanceRole" yaml:"instanceRole"`
	// The EC2 instance type for all instances in the instance group.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The ID of an Amazon EMR cluster that you want to associate this instance group with.
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
	// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
	//
	// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
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
	// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// Market type of the EC2 instances used to create a cluster node.
	Market *string `field:"optional" json:"market" yaml:"market"`
	// Friendly name given to the instance group.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A CloudFormation `AWS::EMR::SecurityConfiguration`.
//
// Use a `SecurityConfiguration` resource to configure data encryption, Kerberos authentication (available in Amazon EMR release version 5.10.0 and later), and Amazon S3 authorization for EMRFS (available in EMR 5.10.0 and later). You can re-use a security configuration for any number of clusters in your account. For more information and example security configuration JSON objects, see [Create a Security Configuration](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-create-security-configuration.html) in the *Amazon EMR Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityConfiguration interface{}
//
//   cfnSecurityConfiguration := awscdk.Aws_emr.NewCfnSecurityConfiguration(this, jsii.String("MyCfnSecurityConfiguration"), &cfnSecurityConfigurationProps{
//   	securityConfiguration: securityConfiguration,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   })
//
type CfnSecurityConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The name of the security configuration.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The security configuration details in JSON format.
	SecurityConfiguration() interface{}
	SetSecurityConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnSecurityConfiguration
type jsiiProxy_CfnSecurityConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) SecurityConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::SecurityConfiguration`.
func NewCfnSecurityConfiguration(scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) CfnSecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnSecurityConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::SecurityConfiguration`.
func NewCfnSecurityConfiguration_Override(c CfnSecurityConfiguration, scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetSecurityConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSecurityConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSecurityConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
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
func CfnSecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityConfiguration interface{}
//
//   cfnSecurityConfigurationProps := &cfnSecurityConfigurationProps{
//   	securityConfiguration: securityConfiguration,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnSecurityConfigurationProps struct {
	// The security configuration details in JSON format.
	SecurityConfiguration interface{} `field:"required" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The name of the security configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A CloudFormation `AWS::EMR::Step`.
//
// Use `Step` to specify a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to a cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStep := awscdk.Aws_emr.NewCfnStep(this, jsii.String("MyCfnStep"), &cfnStepProps{
//   	actionOnFailure: jsii.String("actionOnFailure"),
//   	hadoopJarStep: &hadoopJarStepConfigProperty{
//   		jar: jsii.String("jar"),
//
//   		// the properties below are optional
//   		args: []*string{
//   			jsii.String("args"),
//   		},
//   		mainClass: jsii.String("mainClass"),
//   		stepProperties: []interface{}{
//   			&keyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	jobFlowId: jsii.String("jobFlowId"),
//   	name: jsii.String("name"),
//   })
//
type CfnStep interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// This specifies what action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	ActionOnFailure() *string
	SetActionOnFailure(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed.
	//
	// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
	HadoopJarStep() interface{}
	SetHadoopJarStep(val interface{})
	// A string that uniquely identifies the cluster (job flow).
	JobFlowId() *string
	SetJobFlowId(val *string)
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
	// The name of the cluster step.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStep
type jsiiProxy_CfnStep struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStep) ActionOnFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) HadoopJarStep() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hadoopJarStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) JobFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Step`.
func NewCfnStep(scope constructs.Construct, id *string, props *CfnStepProps) CfnStep {
	_init_.Initialize()

	j := jsiiProxy_CfnStep{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStep",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Step`.
func NewCfnStep_Override(c CfnStep, scope constructs.Construct, id *string, props *CfnStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStep",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStep) SetActionOnFailure(val *string) {
	_jsii_.Set(
		j,
		"actionOnFailure",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetHadoopJarStep(val interface{}) {
	_jsii_.Set(
		j,
		"hadoopJarStep",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetJobFlowId(val *string) {
	_jsii_.Set(
		j,
		"jobFlowId",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStep_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStep_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
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
func CfnStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStep_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStep",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStep) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStep) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStep) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStep) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStep) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStep) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStep) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStep) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStep) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStep) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStep) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStep) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStep) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStep) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A job flow step consisting of a JAR file whose main function will be executed.
//
// The main function submits a job for Hadoop to execute and waits for the job to finish or fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hadoopJarStepConfigProperty := &hadoopJarStepConfigProperty{
//   	jar: jsii.String("jar"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	mainClass: jsii.String("mainClass"),
//   	stepProperties: []interface{}{
//   		&keyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStep_HadoopJarStepConfigProperty struct {
	// A path to a JAR file run during the step.
	Jar *string `field:"required" json:"jar" yaml:"jar"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	StepProperties interface{} `field:"optional" json:"stepProperties" yaml:"stepProperties"`
}

// `KeyValue` is a subproperty of the `HadoopJarStepConfig` property type.
//
// `KeyValue` is used to pass parameters to a step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueProperty := &keyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnStep_KeyValueProperty struct {
	// The unique identifier of a key-value pair.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value part of the identified key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnStep`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStepProps := &cfnStepProps{
//   	actionOnFailure: jsii.String("actionOnFailure"),
//   	hadoopJarStep: &hadoopJarStepConfigProperty{
//   		jar: jsii.String("jar"),
//
//   		// the properties below are optional
//   		args: []*string{
//   			jsii.String("args"),
//   		},
//   		mainClass: jsii.String("mainClass"),
//   		stepProperties: []interface{}{
//   			&keyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	jobFlowId: jsii.String("jobFlowId"),
//   	name: jsii.String("name"),
//   }
//
type CfnStepProps struct {
	// This specifies what action to take when the cluster step fails.
	//
	// Possible values are `CANCEL_AND_WAIT` and `CONTINUE` .
	ActionOnFailure *string `field:"required" json:"actionOnFailure" yaml:"actionOnFailure"`
	// The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed.
	//
	// The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.
	HadoopJarStep interface{} `field:"required" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// A string that uniquely identifies the cluster (job flow).
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
	// The name of the cluster step.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// A CloudFormation `AWS::EMR::Studio`.
//
// The `AWS::EMR::Studio` resource specifies an Amazon EMR Studio. An EMR Studio is a web-based, integrated development environment for fully managed Jupyter notebooks that run on Amazon EMR clusters. For more information, see the [*Amazon EMR Management Guide*](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudio := awscdk.Aws_emr.NewCfnStudio(this, jsii.String("MyCfnStudio"), &cfnStudioProps{
//   	authMode: jsii.String("authMode"),
//   	defaultS3Location: jsii.String("defaultS3Location"),
//   	engineSecurityGroupId: jsii.String("engineSecurityGroupId"),
//   	name: jsii.String("name"),
//   	serviceRole: jsii.String("serviceRole"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   	workspaceSecurityGroupId: jsii.String("workspaceSecurityGroupId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	idpAuthUrl: jsii.String("idpAuthUrl"),
//   	idpRelayStateParameterName: jsii.String("idpRelayStateParameterName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userRole: jsii.String("userRole"),
//   })
//
type CfnStudio interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the Amazon EMR Studio.
	//
	// For example: `arn:aws:elasticmapreduce:us-east-1:653XXXXXXXXX:studio/es-EXAMPLE12345678XXXXXXXXXXX` .
	AttrArn() *string
	// The ID of the Amazon EMR Studio.
	//
	// For example: `es-EXAMPLE12345678XXXXXXXXXXX` .
	AttrStudioId() *string
	// The unique access URL of the Amazon EMR Studio.
	//
	// For example: `https://es-EXAMPLE12345678XXXXXXXXXXX.emrstudio-prod.us-east-1.amazonaws.com` .
	AttrUrl() *string
	// Specifies whether the Studio authenticates users using AWS SSO or IAM.
	AuthMode() *string
	SetAuthMode(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon S3 location to back up EMR Studio Workspaces and notebook files.
	DefaultS3Location() *string
	SetDefaultS3Location(val *string)
	// A detailed description of the Amazon EMR Studio.
	Description() *string
	SetDescription(val *string)
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by `VpcId` .
	EngineSecurityGroupId() *string
	SetEngineSecurityGroupId(val *string)
	// Your identity provider's authentication endpoint.
	//
	// Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	IdpAuthUrl() *string
	SetIdpAuthUrl(val *string)
	// The name of your identity provider's `RelayState` parameter.
	IdpRelayStateParameterName() *string
	SetIdpRelayStateParameterName(val *string)
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
	// A descriptive name for the Amazon EMR Studio.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role that will be assumed by the Amazon EMR Studio.
	//
	// The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of subnet IDs to associate with the Amazon EMR Studio.
	//
	// A Studio can have a maximum of 5 subnets. The subnets must belong to the VPC specified by `VpcId` . Studio users can create a Workspace in any of the specified subnets.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The Amazon Resource Name (ARN) of the IAM user role that will be assumed by users and groups logged in to a Studio.
	//
	// The permissions attached to this IAM role can be scoped down for each user or group using session policies. You only need to specify `UserRole` when you set `AuthMode` to `SSO` .
	UserRole() *string
	SetUserRole(val *string)
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	VpcId() *string
	SetVpcId(val *string)
	// The ID of the Workspace security group associated with the Amazon EMR Studio.
	//
	// The Workspace security group allows outbound network traffic to resources in the Engine security group and to the internet.
	WorkspaceSecurityGroupId() *string
	SetWorkspaceSecurityGroupId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStudio
type jsiiProxy_CfnStudio struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudio) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrStudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStudioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) DefaultS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) EngineSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) IdpAuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpAuthUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) IdpRelayStateParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpRelayStateParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Studio`.
func NewCfnStudio(scope constructs.Construct, id *string, props *CfnStudioProps) CfnStudio {
	_init_.Initialize()

	j := jsiiProxy_CfnStudio{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Studio`.
func NewCfnStudio_Override(c CfnStudio, scope constructs.Construct, id *string, props *CfnStudioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudio",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudio) SetAuthMode(val *string) {
	_jsii_.Set(
		j,
		"authMode",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetDefaultS3Location(val *string) {
	_jsii_.Set(
		j,
		"defaultS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetEngineSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"engineSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetIdpAuthUrl(val *string) {
	_jsii_.Set(
		j,
		"idpAuthUrl",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetIdpRelayStateParameterName(val *string) {
	_jsii_.Set(
		j,
		"idpRelayStateParameterName",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetUserRole(val *string) {
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetWorkspaceSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"workspaceSecurityGroupId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStudio_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStudio_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
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
func CfnStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudio_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudio) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStudio) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStudio) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStudio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStudio) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStudio) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStudio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStudio) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStudio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudio) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudio) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStudio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioProps := &cfnStudioProps{
//   	authMode: jsii.String("authMode"),
//   	defaultS3Location: jsii.String("defaultS3Location"),
//   	engineSecurityGroupId: jsii.String("engineSecurityGroupId"),
//   	name: jsii.String("name"),
//   	serviceRole: jsii.String("serviceRole"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   	workspaceSecurityGroupId: jsii.String("workspaceSecurityGroupId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	idpAuthUrl: jsii.String("idpAuthUrl"),
//   	idpRelayStateParameterName: jsii.String("idpRelayStateParameterName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userRole: jsii.String("userRole"),
//   }
//
type CfnStudioProps struct {
	// Specifies whether the Studio authenticates users using AWS SSO or IAM.
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The Amazon S3 location to back up EMR Studio Workspaces and notebook files.
	DefaultS3Location *string `field:"required" json:"defaultS3Location" yaml:"defaultS3Location"`
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by `VpcId` .
	EngineSecurityGroupId *string `field:"required" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// A descriptive name for the Amazon EMR Studio.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role that will be assumed by the Amazon EMR Studio.
	//
	// The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// A list of subnet IDs to associate with the Amazon EMR Studio.
	//
	// A Studio can have a maximum of 5 subnets. The subnets must belong to the VPC specified by `VpcId` . Studio users can create a Workspace in any of the specified subnets.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the Workspace security group associated with the Amazon EMR Studio.
	//
	// The Workspace security group allows outbound network traffic to resources in the Engine security group and to the internet.
	WorkspaceSecurityGroupId *string `field:"required" json:"workspaceSecurityGroupId" yaml:"workspaceSecurityGroupId"`
	// A detailed description of the Amazon EMR Studio.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Your identity provider's authentication endpoint.
	//
	// Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	IdpAuthUrl *string `field:"optional" json:"idpAuthUrl" yaml:"idpAuthUrl"`
	// The name of your identity provider's `RelayState` parameter.
	IdpRelayStateParameterName *string `field:"optional" json:"idpRelayStateParameterName" yaml:"idpRelayStateParameterName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the IAM user role that will be assumed by users and groups logged in to a Studio.
	//
	// The permissions attached to this IAM role can be scoped down for each user or group using session policies. You only need to specify `UserRole` when you set `AuthMode` to `SSO` .
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
}

// A CloudFormation `AWS::EMR::StudioSessionMapping`.
//
// The `AWS::EMR::StudioSessionMapping` resource is an Amazon EMR resource type that maps a user or group to the Amazon EMR Studio specified by `StudioId` , and applies a session policy that defines Studio permissions for that user or group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioSessionMapping := awscdk.Aws_emr.NewCfnStudioSessionMapping(this, jsii.String("MyCfnStudioSessionMapping"), &cfnStudioSessionMappingProps{
//   	identityName: jsii.String("identityName"),
//   	identityType: jsii.String("identityType"),
//   	sessionPolicyArn: jsii.String("sessionPolicyArn"),
//   	studioId: jsii.String("studioId"),
//   })
//
type CfnStudioSessionMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the user or group.
	//
	// For more information, see [UserName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName) and [DisplayName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName) in the *AWS SSO Identity Store API Reference* .
	IdentityName() *string
	SetIdentityName(val *string)
	// Specifies whether the identity to map to the Amazon EMR Studio is a user or a group.
	IdentityType() *string
	SetIdentityType(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group.
	//
	// Session policies refine Studio user permissions without the need to use multiple IAM user roles. For more information, see [Create an EMR Studio user role with session policies](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-user-role.html) in the *Amazon EMR Management Guide* .
	SessionPolicyArn() *string
	SetSessionPolicyArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId() *string
	SetStudioId(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStudioSessionMapping
type jsiiProxy_CfnStudioSessionMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) IdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) SessionPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMapping(scope constructs.Construct, id *string, props *CfnStudioSessionMappingProps) CfnStudioSessionMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnStudioSessionMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMapping_Override(c CfnStudioSessionMapping, scope constructs.Construct, id *string, props *CfnStudioSessionMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetIdentityName(val *string) {
	_jsii_.Set(
		j,
		"identityName",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetIdentityType(val *string) {
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetSessionPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"sessionPolicyArn",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStudioSessionMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStudioSessionMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
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
func CfnStudioSessionMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudioSessionMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStudioSessionMapping) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStudioSessionMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioSessionMappingProps := &cfnStudioSessionMappingProps{
//   	identityName: jsii.String("identityName"),
//   	identityType: jsii.String("identityType"),
//   	sessionPolicyArn: jsii.String("sessionPolicyArn"),
//   	studioId: jsii.String("studioId"),
//   }
//
type CfnStudioSessionMappingProps struct {
	// The name of the user or group.
	//
	// For more information, see [UserName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName) and [DisplayName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName) in the *AWS SSO Identity Store API Reference* .
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
	// Specifies whether the identity to map to the Amazon EMR Studio is a user or a group.
	IdentityType *string `field:"required" json:"identityType" yaml:"identityType"`
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group.
	//
	// Session policies refine Studio user permissions without the need to use multiple IAM user roles. For more information, see [Create an EMR Studio user role with session policies](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-user-role.html) in the *Amazon EMR Management Guide* .
	SessionPolicyArn *string `field:"required" json:"sessionPolicyArn" yaml:"sessionPolicyArn"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
}

