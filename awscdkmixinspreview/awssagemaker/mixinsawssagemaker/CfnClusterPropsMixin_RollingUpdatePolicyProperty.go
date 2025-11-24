package mixinsawssagemaker


// Specifies a rolling deployment strategy for updating a SageMaker endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rollingUpdatePolicyProperty := &RollingUpdatePolicyProperty{
//   	MaximumBatchSize: &CapacitySizeConfigProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-rollingupdatepolicy.html
//
type CfnClusterPropsMixin_RollingUpdatePolicyProperty struct {
	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet.
	//
	// Value must be between 5% to 50% of the variant's total instance count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-rollingupdatepolicy.html#cfn-sagemaker-cluster-rollingupdatepolicy-maximumbatchsize
	//
	MaximumBatchSize interface{} `field:"optional" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// Batch size for rollback to the old endpoint fleet.
	//
	// Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100% of total capacity which means to bring up the whole capacity of the old fleet at once during rollback.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-rollingupdatepolicy.html#cfn-sagemaker-cluster-rollingupdatepolicy-rollbackmaximumbatchsize
	//
	RollbackMaximumBatchSize interface{} `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
}

