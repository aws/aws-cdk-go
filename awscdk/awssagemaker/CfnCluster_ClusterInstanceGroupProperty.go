package awssagemaker


// Details of an instance group in a SageMaker HyperPod cluster.
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
	// The number of instances you specified to add to the instance group of a SageMaker HyperPod cluster.
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
	// The instance storage configuration for the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancestorageconfigs
	//
	InstanceStorageConfigs interface{} `field:"optional" json:"instanceStorageConfigs" yaml:"instanceStorageConfigs"`
	// Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-onstartdeephealthchecks
	//
	OnStartDeepHealthChecks *[]*string `field:"optional" json:"onStartDeepHealthChecks" yaml:"onStartDeepHealthChecks"`
	// The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading.
	//
	// For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

