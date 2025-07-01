package awssagemaker


// Specifies a rolling deployment strategy for updating a SageMaker endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingUpdatePolicyProperty := &RollingUpdatePolicyProperty{
//   	MaximumBatchSize: &CapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	WaitIntervalInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   	RollbackMaximumBatchSize: &CapacitySizeProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html
//
type CfnEndpoint_RollingUpdatePolicyProperty struct {
	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet.
	//
	// Value must be between 5% to 50% of the variant's total instance count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html#cfn-sagemaker-endpoint-rollingupdatepolicy-maximumbatchsize
	//
	MaximumBatchSize interface{} `field:"required" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html#cfn-sagemaker-endpoint-rollingupdatepolicy-waitintervalinseconds
	//
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// The time limit for the total deployment.
	//
	// Exceeding this limit causes a timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html#cfn-sagemaker-endpoint-rollingupdatepolicy-maximumexecutiontimeoutinseconds
	//
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Batch size for rollback to the old endpoint fleet.
	//
	// Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100% of total capacity which means to bring up the whole capacity of the old fleet at once during rollback.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-rollingupdatepolicy.html#cfn-sagemaker-endpoint-rollingupdatepolicy-rollbackmaximumbatchsize
	//
	RollbackMaximumBatchSize interface{} `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
}

