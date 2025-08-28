package awsemr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html
//
type CfnClusterProps struct {
	// A specification of the number and type of Amazon EC2 instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-instances
	//
	Instances interface{} `field:"required" json:"instances" yaml:"instances"`
	// Also called instance profile and Amazon EC2 role.
	//
	// An IAM role for an Amazon EMR cluster. The Amazon EC2 instances of the cluster assume this role. The default role is `EMR_EC2_DefaultRole` . In order to use the default role, you must have already created it using the AWS CLI or console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-jobflowrole
	//
	JobFlowRole *string `field:"required" json:"jobFlowRole" yaml:"jobFlowRole"`
	// The name of the cluster.
	//
	// This parameter can't contain the characters <, >, $, |, or ` (backtick).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IAM role that Amazon EMR assumes in order to access AWS resources on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-cluster.html#cfn-emr-cluster-servicerole
	//
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
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

