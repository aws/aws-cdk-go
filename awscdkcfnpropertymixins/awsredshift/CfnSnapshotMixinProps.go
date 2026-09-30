package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSnapshotMixinProps := &CfnSnapshotMixinProps{
//   	ClusterIdentifier: jsii.String("clusterIdentifier"),
//   	ManualSnapshotRetentionPeriod: jsii.Number(123),
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshot.html
//
type CfnSnapshotMixinProps struct {
	// The cluster identifier for which you want a snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshot.html#cfn-redshift-snapshot-clusteridentifier
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The number of days that a manual snapshot is retained.
	//
	// If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3653.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshot.html#cfn-redshift-snapshot-manualsnapshotretentionperiod
	//
	ManualSnapshotRetentionPeriod *float64 `field:"optional" json:"manualSnapshotRetentionPeriod" yaml:"manualSnapshotRetentionPeriod"`
	// A unique identifier for the snapshot that you are requesting.
	//
	// This identifier must be unique for all snapshots within the Amazon Web Services account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshot.html#cfn-redshift-snapshot-snapshotidentifier
	//
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// A list of tag instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshot.html#cfn-redshift-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

