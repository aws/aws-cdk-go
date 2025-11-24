package mixinsawssagemaker


// Details of a restricted instance group in a SageMaker HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterRestrictedInstanceGroupProperty := &ClusterRestrictedInstanceGroupProperty{
//   	CurrentCount: jsii.Number(123),
//   	EnvironmentConfig: &EnvironmentConfigProperty{
//   		FSxLustreConfig: &FSxLustreConfigProperty{
//   			PerUnitStorageThroughput: jsii.Number(123),
//   			SizeInGiB: jsii.Number(123),
//   		},
//   	},
//   	ExecutionRole: jsii.String("executionRole"),
//   	InstanceCount: jsii.Number(123),
//   	InstanceGroupName: jsii.String("instanceGroupName"),
//   	InstanceStorageConfigs: []interface{}{
//   		&ClusterInstanceStorageConfigProperty{
//   			EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   				RootVolume: jsii.Boolean(false),
//   				VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   				VolumeSizeInGb: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
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
//   	TrainingPlanArn: jsii.String("trainingPlanArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html
//
type CfnClusterPropsMixin_ClusterRestrictedInstanceGroupProperty struct {
	// The number of instances that are currently in the restricted instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-currentcount
	//
	CurrentCount *float64 `field:"optional" json:"currentCount" yaml:"currentCount"`
	// The configuration for the restricted instance groups (RIG) environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-environmentconfig
	//
	EnvironmentConfig interface{} `field:"optional" json:"environmentConfig" yaml:"environmentConfig"`
	// The execution role for the instance group to assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The number of instances you specified to add to the restricted instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The name of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-instancegroupname
	//
	InstanceGroupName *string `field:"optional" json:"instanceGroupName" yaml:"instanceGroupName"`
	// The instance storage configuration for the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-instancestorageconfigs
	//
	InstanceStorageConfigs interface{} `field:"optional" json:"instanceStorageConfigs" yaml:"instanceStorageConfigs"`
	// The instance type of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-onstartdeephealthchecks
	//
	OnStartDeepHealthChecks *[]*string `field:"optional" json:"onStartDeepHealthChecks" yaml:"onStartDeepHealthChecks"`
	// Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to.
	//
	// You can control access to and from your resources by configuring a VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-overridevpcconfig
	//
	OverrideVpcConfig interface{} `field:"optional" json:"overrideVpcConfig" yaml:"overrideVpcConfig"`
	// The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading.
	//
	// For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// The Amazon Resource Name (ARN) of the training plan to use for this cluster restricted instance group.
	//
	// For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.html#cfn-sagemaker-cluster-clusterrestrictedinstancegroup-trainingplanarn
	//
	TrainingPlanArn *string `field:"optional" json:"trainingPlanArn" yaml:"trainingPlanArn"`
}

