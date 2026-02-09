package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDatabaseSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabaseSnapshotProps := &CfnDatabaseSnapshotProps{
//   	RelationalDatabaseName: jsii.String("relationalDatabaseName"),
//   	RelationalDatabaseSnapshotName: jsii.String("relationalDatabaseSnapshotName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-databasesnapshot.html
//
type CfnDatabaseSnapshotProps struct {
	// The name of the database on which to base your new snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-databasesnapshot.html#cfn-lightsail-databasesnapshot-relationaldatabasename
	//
	RelationalDatabaseName *string `field:"required" json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// The name for your new database snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-databasesnapshot.html#cfn-lightsail-databasesnapshot-relationaldatabasesnapshotname
	//
	RelationalDatabaseSnapshotName *string `field:"required" json:"relationalDatabaseSnapshotName" yaml:"relationalDatabaseSnapshotName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-databasesnapshot.html#cfn-lightsail-databasesnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

