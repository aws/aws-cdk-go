package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotProps := &CfnSnapshotProps{
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	NamespaceName: jsii.String("namespaceName"),
//   	RetentionPeriod: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-snapshot.html
//
type CfnSnapshotProps struct {
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-snapshot.html#cfn-redshiftserverless-snapshot-snapshotname
	//
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// The namespace the snapshot is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-snapshot.html#cfn-redshiftserverless-snapshot-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The retention period of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-snapshot.html#cfn-redshiftserverless-snapshot-retentionperiod
	//
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-snapshot.html#cfn-redshiftserverless-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

