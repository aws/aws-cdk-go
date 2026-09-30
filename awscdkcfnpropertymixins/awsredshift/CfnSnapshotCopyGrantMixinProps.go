package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSnapshotCopyGrantPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSnapshotCopyGrantMixinProps := &CfnSnapshotCopyGrantMixinProps{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotCopyGrantName: jsii.String("snapshotCopyGrantName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotcopygrant.html
//
type CfnSnapshotCopyGrantMixinProps struct {
	// The unique identifier of the encrypted symmetric key to which to grant Amazon Redshift permission.
	//
	// If no key is specified, the default key is used. Marked writeOnly because Create accepts aliases/key-ids but Describe returns the resolved key ARN, which would fail drift detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotcopygrant.html#cfn-redshift-snapshotcopygrant-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the snapshot copy grant.
	//
	// This name must be unique in the region for the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotcopygrant.html#cfn-redshift-snapshotcopygrant-snapshotcopygrantname
	//
	SnapshotCopyGrantName *string `field:"optional" json:"snapshotCopyGrantName" yaml:"snapshotCopyGrantName"`
	// A list of tag instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotcopygrant.html#cfn-redshift-snapshotcopygrant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

