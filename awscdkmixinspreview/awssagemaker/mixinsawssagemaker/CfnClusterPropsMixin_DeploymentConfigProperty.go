package mixinsawssagemaker


// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deploymentConfigProperty := &DeploymentConfigProperty{
//   	AutoRollbackConfiguration: []interface{}{
//   		&AlarmDetailsProperty{
//   			AlarmName: jsii.String("alarmName"),
//   		},
//   	},
//   	RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   		MaximumBatchSize: &CapacitySizeConfigProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	WaitIntervalInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-deploymentconfig.html
//
type CfnClusterPropsMixin_DeploymentConfigProperty struct {
	// Automatic rollback configuration for handling endpoint deployment failures and recovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-deploymentconfig.html#cfn-sagemaker-cluster-deploymentconfig-autorollbackconfiguration
	//
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// Specifies a rolling deployment strategy for updating a SageMaker endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-deploymentconfig.html#cfn-sagemaker-cluster-deploymentconfig-rollingupdatepolicy
	//
	RollingUpdatePolicy interface{} `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
	// The duration in seconds that SageMaker waits before updating more instances in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-deploymentconfig.html#cfn-sagemaker-cluster-deploymentconfig-waitintervalinseconds
	//
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

