package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBSnapshotPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDBSnapshotMixinProps := &CfnDBSnapshotMixinProps{
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	DbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html
//
type CfnDBSnapshotMixinProps struct {
	// The identifier of the DB instance that you want to create the snapshot of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-dbinstanceidentifier
	//
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The identifier for the DB snapshot.
	//
	// Must contain from 1 to 255 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-dbsnapshotidentifier
	//
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// The tags to be assigned to the DB snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

