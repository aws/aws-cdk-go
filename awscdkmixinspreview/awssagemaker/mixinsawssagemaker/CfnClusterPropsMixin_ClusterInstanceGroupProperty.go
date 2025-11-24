package mixinsawssagemaker


// The configuration information of the instance group within the HyperPod cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   clusterInstanceGroupProperty := &ClusterInstanceGroupProperty{
//   	CurrentCount: jsii.Number(123),
//   	ExecutionRole: jsii.String("executionRole"),
//   	ImageId: jsii.String("imageId"),
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
//   	LifeCycleConfig: &ClusterLifeCycleConfigProperty{
//   		OnCreate: jsii.String("onCreate"),
//   		SourceS3Uri: jsii.String("sourceS3Uri"),
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
//   	ScheduledUpdateConfig: &ScheduledUpdateConfigProperty{
//   		DeploymentConfig: &DeploymentConfigProperty{
//   			AutoRollbackConfiguration: []interface{}{
//   				&AlarmDetailsProperty{
//   					AlarmName: jsii.String("alarmName"),
//   				},
//   			},
//   			RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   				MaximumBatchSize: &CapacitySizeConfigProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   				RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			WaitIntervalInSeconds: jsii.Number(123),
//   		},
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	ThreadsPerCore: jsii.Number(123),
//   	TrainingPlanArn: jsii.String("trainingPlanArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html
//
type CfnClusterPropsMixin_ClusterInstanceGroupProperty struct {
	// The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-currentcount
	//
	CurrentCount *float64 `field:"optional" json:"currentCount" yaml:"currentCount"`
	// The execution role for the instance group to assume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// AMI Id to be used for launching EC2 instances - HyperPodPublicAmiId or CustomAmiId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-imageid
	//
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The number of instances in an instance group of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The name of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancegroupname
	//
	InstanceGroupName *string `field:"optional" json:"instanceGroupName" yaml:"instanceGroupName"`
	// The configurations of additional storage specified to the instance group where the instance (node) is launched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancestorageconfigs
	//
	InstanceStorageConfigs interface{} `field:"optional" json:"instanceStorageConfigs" yaml:"instanceStorageConfigs"`
	// The instance type of the instance group of a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The lifecycle configuration for a SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-lifecycleconfig
	//
	LifeCycleConfig interface{} `field:"optional" json:"lifeCycleConfig" yaml:"lifeCycleConfig"`
	// A flag indicating whether deep health checks should be performed when the HyperPod cluster instance group is created or updated.
	//
	// Deep health checks are comprehensive, invasive tests that validate the health of the underlying hardware and infrastructure components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-onstartdeephealthchecks
	//
	OnStartDeepHealthChecks *[]*string `field:"optional" json:"onStartDeepHealthChecks" yaml:"onStartDeepHealthChecks"`
	// The customized Amazon VPC configuration at the instance group level that overrides the default Amazon VPC configuration of the SageMaker HyperPod cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-overridevpcconfig
	//
	OverrideVpcConfig interface{} `field:"optional" json:"overrideVpcConfig" yaml:"overrideVpcConfig"`
	// The configuration object of the schedule that SageMaker follows when updating the AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-scheduledupdateconfig
	//
	ScheduledUpdateConfig interface{} `field:"optional" json:"scheduledUpdateConfig" yaml:"scheduledUpdateConfig"`
	// The number of threads per CPU core you specified under `CreateCluster` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
	// The Amazon Resource Name (ARN) of the training plan to use for this cluster instance group.
	//
	// For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusterinstancegroup.html#cfn-sagemaker-cluster-clusterinstancegroup-trainingplanarn
	//
	TrainingPlanArn *string `field:"optional" json:"trainingPlanArn" yaml:"trainingPlanArn"`
}

