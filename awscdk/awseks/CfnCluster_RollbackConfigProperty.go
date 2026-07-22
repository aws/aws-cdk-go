package awseks


// The rollback configuration to use for the cluster version rollback.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollbackConfigProperty := &RollbackConfigProperty{
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-rollbackconfig.html
//
type CfnCluster_RollbackConfigProperty struct {
	// The timeout in minutes for the version rollback operation.
	//
	// If not specified, defaults to 720 minutes (12 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-rollbackconfig.html#cfn-eks-cluster-rollbackconfig-timeoutminutes
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

