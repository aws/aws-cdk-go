package mixinsawssagemaker


// The deployment configuration for an endpoint that hosts inference components.
//
// The configuration includes the desired deployment strategy and rollback settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inferenceComponentDeploymentConfigProperty := &InferenceComponentDeploymentConfigProperty{
//   	AutoRollbackConfiguration: &AutoRollbackConfigurationProperty{
//   		Alarms: []interface{}{
//   			&AlarmProperty{
//   				AlarmName: jsii.String("alarmName"),
//   			},
//   		},
//   	},
//   	RollingUpdatePolicy: &InferenceComponentRollingUpdatePolicyProperty{
//   		MaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   		RollbackMaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.Number(123),
//   		},
//   		WaitIntervalInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentDeploymentConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig-autorollbackconfiguration
	//
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// Specifies a rolling deployment strategy for updating a SageMaker AI endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig-rollingupdatepolicy
	//
	RollingUpdatePolicy interface{} `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

