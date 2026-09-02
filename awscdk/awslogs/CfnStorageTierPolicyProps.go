package awslogs


// Properties for defining a `CfnStorageTierPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStorageTierPolicyProps := &CfnStorageTierPolicyProps{
//   	StorageTier: jsii.String("storageTier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-storagetierpolicy.html
//
type CfnStorageTierPolicyProps struct {
	// The storage tier to apply.
	//
	// Only INTELLIGENT_TIERING is accepted for creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-storagetierpolicy.html#cfn-logs-storagetierpolicy-storagetier
	//
	StorageTier *string `field:"required" json:"storageTier" yaml:"storageTier"`
}

