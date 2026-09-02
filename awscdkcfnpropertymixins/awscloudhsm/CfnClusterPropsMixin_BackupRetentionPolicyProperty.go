package awscloudhsm


// A policy that defines how the service retains backups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   backupRetentionPolicyProperty := &BackupRetentionPolicyProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudhsm-cluster-backupretentionpolicy.html
//
type CfnClusterPropsMixin_BackupRetentionPolicyProperty struct {
	// The type of backup retention policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudhsm-cluster-backupretentionpolicy.html#cfn-cloudhsm-cluster-backupretentionpolicy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Use a value between 7 - 379.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudhsm-cluster-backupretentionpolicy.html#cfn-cloudhsm-cluster-backupretentionpolicy-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

