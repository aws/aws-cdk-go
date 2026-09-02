package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSnapshotProps := &CfnDBSnapshotProps{
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	DbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//
//   	// the properties below are optional
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
type CfnDBSnapshotProps struct {
	// The identifier of the DB instance that you want to create the snapshot of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-dbinstanceidentifier
	//
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The identifier for the DB snapshot.
	//
	// Must contain from 1 to 255 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-dbsnapshotidentifier
	//
	DbSnapshotIdentifier *string `field:"required" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// The tags to be assigned to the DB snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsnapshot.html#cfn-rds-dbsnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

