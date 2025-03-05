package awssagemaker


// The configuration information of the instance group within the HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterInstanceGroupProperty := &ClusterInstanceGroupProperty{
//   	ExecutionRole: jsii.String("executionRole"),
//   	InstanceCount: jsii.Number(123),
//   	InstanceGroupName: jsii.String("instanceGroupName"),
//   	InstanceType: jsii.String("instanceType"),
//   	LifeCycleConfig: &ClusterLifeCycleConfigProperty{
//   		OnCreate: jsii.String("onCreate"),
//   		SourceS3Uri: jsii.String("sourceS3Uri"),
//   	},
//
//   	// the properties below are optional
//   	CurrentCount: jsii.Number(123),
//   	InstanceStorageConfigs: []interface{}{
//   		&ClusterInstanceStorageConfigProperty{
//   			EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   				VolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	OnStartDeepHealthChecks: []*string{
//   		jsii.String("onStartDeepHealthChecks"),
//   	},
//   	OverrideVpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	ThreadsPerCore: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html
//
type CfnCluster_ClusterInstanceGroupProperty struct {
	// The execution role for the instance group to assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// The number of instances in an instance group of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancecount
	//
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// The name of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancegroupname
	//
	InstanceGroupName *string `field:"required" json:"instanceGroupName" yaml:"instanceGroupName"`
	// The instance type of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The lifecycle configuration for a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-lifecycleconfig
	//
	LifeCycleConfig interface{} `field:"required" json:"lifeCycleConfig" yaml:"lifeCycleConfig"`
	// The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-currentcount
	//
	CurrentCount *float64 `field:"optional" json:"currentCount" yaml:"currentCount"`
	// The configurations of additional storage specified to the instance group where the instance (node) is launched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancestorageconfigs
	//
	InstanceStorageConfigs interface{} `field:"optional" json:"instanceStorageConfigs" yaml:"instanceStorageConfigs"`
	// A flag indicating whether deep health checks should be performed when the HyperPod cluster instance group is created or updated.
	//
	// Deep health checks are comprehensive, invasive tests that validate the health of the underlying hardware and infrastructure components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-onstartdeephealthchecks
	//
	OnStartDeepHealthChecks *[]*string `field:"optional" json:"onStartDeepHealthChecks" yaml:"onStartDeepHealthChecks"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-overridevpcconfig
	//
	OverrideVpcConfig interface{} `field:"optional" json:"overrideVpcConfig" yaml:"overrideVpcConfig"`
	// The number of threads per CPU core you specified under `CreateCluster` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

