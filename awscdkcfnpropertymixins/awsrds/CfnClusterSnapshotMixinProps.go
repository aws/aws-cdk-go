package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnClusterSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnClusterSnapshotMixinProps := &CfnClusterSnapshotMixinProps{
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterSnapshotIdentifier: jsii.String("dbClusterSnapshotIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-clustersnapshot.html
//
type CfnClusterSnapshotMixinProps struct {
	// The identifier of the DB cluster to create a snapshot for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-clustersnapshot.html#cfn-rds-clustersnapshot-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The identifier for the DB cluster snapshot.
	//
	// Must contain from 1 to 63 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-clustersnapshot.html#cfn-rds-clustersnapshot-dbclustersnapshotidentifier
	//
	DbClusterSnapshotIdentifier *string `field:"optional" json:"dbClusterSnapshotIdentifier" yaml:"dbClusterSnapshotIdentifier"`
	// The tags to be assigned to the DB cluster snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-clustersnapshot.html#cfn-rds-clustersnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

