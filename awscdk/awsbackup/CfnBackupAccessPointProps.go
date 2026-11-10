package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBackupAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPointPolicy interface{}
//
//   cfnBackupAccessPointProps := &CfnBackupAccessPointProps{
//   	RecoveryPointArn: jsii.String("recoveryPointArn"),
//
//   	// the properties below are optional
//   	AccessPointInTime: jsii.String("accessPointInTime"),
//   	AccessPointPolicy: accessPointPolicy,
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html
//
type CfnBackupAccessPointProps struct {
	// The Amazon Resource Name (ARN) of the recovery point for which to create the backup access point.
	//
	// The recovery point must be an Amazon S3 recovery point in the AVAILABLE, STOPPED, or COMPLETED state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html#cfn-backup-backupaccesspoint-recoverypointarn
	//
	RecoveryPointArn *string `field:"required" json:"recoveryPointArn" yaml:"recoveryPointArn"`
	// For continuous (point-in-time) recovery points, the timestamp in ISO 8601 format (for example 2021-11-27T03:30:27Z) that determines the point-in-time view of the backup data.
	//
	// The timestamp must be within the continuous backup's retention period at the time the access point is created. Do not specify this for snapshot recovery points. The service reports it back under AccessPointMetadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html#cfn-backup-backupaccesspoint-accesspointintime
	//
	AccessPointInTime *string `field:"optional" json:"accessPointInTime" yaml:"accessPointInTime"`
	// An optional resource-based policy document, in JSON, to apply to the underlying Amazon S3 access point.
	//
	// The policy controls how backup data can be accessed through the access point. If you do not specify a policy, access is governed by the caller's IAM permissions. The serialized policy may not exceed 200000 characters. This value is not returned by the service and cannot be updated in place.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html#cfn-backup-backupaccesspoint-accesspointpolicy
	//
	AccessPointPolicy interface{} `field:"optional" json:"accessPointPolicy" yaml:"accessPointPolicy"`
	// The name of the backup access point.
	//
	// This name is shared with the Amazon S3 access point namespace. It must be unique within your account and Region and cannot conflict with an existing Amazon S3 access point. If omitted, a name is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html#cfn-backup-backupaccesspoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to assign to the backup access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupaccesspoint.html#cfn-backup-backupaccesspoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

