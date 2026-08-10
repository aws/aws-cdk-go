package awsstoragegateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnTapePoolPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTapePoolMixinProps := &CfnTapePoolMixinProps{
//   	PoolName: jsii.String("poolName"),
//   	RetentionLockTimeInDays: jsii.Number(123),
//   	RetentionLockType: jsii.String("retentionLockType"),
//   	StorageClass: jsii.String("storageClass"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html
//
type CfnTapePoolMixinProps struct {
	// The name of the custom tape pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html#cfn-storagegateway-tapepool-poolname
	//
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// Tape retention lock time in days (up to 36,500 days / 100 years).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html#cfn-storagegateway-tapepool-retentionlocktimeindays
	//
	RetentionLockTimeInDays *float64 `field:"optional" json:"retentionLockTimeInDays" yaml:"retentionLockTimeInDays"`
	// Tape retention lock type.
	//
	// Governance mode allows authorized removal; compliance mode prevents all removal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html#cfn-storagegateway-tapepool-retentionlocktype
	//
	RetentionLockType *string `field:"optional" json:"retentionLockType" yaml:"retentionLockType"`
	// The storage class associated with the custom pool (S3 Glacier or S3 Glacier Deep Archive).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html#cfn-storagegateway-tapepool-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// A list of up to 50 tags for the tape pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-tapepool.html#cfn-storagegateway-tapepool-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

