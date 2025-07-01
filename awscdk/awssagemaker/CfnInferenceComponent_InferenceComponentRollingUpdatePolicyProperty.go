package awssagemaker


// Specifies a rolling deployment strategy for updating a SageMaker AI inference component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentRollingUpdatePolicyProperty := &InferenceComponentRollingUpdatePolicyProperty{
//   	MaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   	RollbackMaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	WaitIntervalInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.html
//
type CfnInferenceComponent_InferenceComponentRollingUpdatePolicyProperty struct {
	// The batch size for each rolling step in the deployment process.
	//
	// For each step, SageMaker AI provisions capacity on the new endpoint fleet, routes traffic to that fleet, and terminates capacity on the old endpoint fleet. The value must be between 5% to 50% of the copy count of the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.html#cfn-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy-maximumbatchsize
	//
	MaximumBatchSize interface{} `field:"optional" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// The time limit for the total deployment.
	//
	// Exceeding this limit causes a timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.html#cfn-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy-maximumexecutiontimeoutinseconds
	//
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// The batch size for a rollback to the old endpoint fleet.
	//
	// If this field is absent, the value is set to the default, which is 100% of the total capacity. When the default is used, SageMaker AI provisions the entire capacity of the old fleet at once during rollback.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.html#cfn-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy-rollbackmaximumbatchsize
	//
	RollbackMaximumBatchSize interface{} `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
	// The length of the baking period, during which SageMaker AI monitors alarms for each batch on the new fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy.html#cfn-sagemaker-inferencecomponent-inferencecomponentrollingupdatepolicy-waitintervalinseconds
	//
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

