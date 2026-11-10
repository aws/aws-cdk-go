package awsbackup


// Metadata the service reports for the backup access point.
//
// Once the access point reaches the AVAILABLE status this carries the Amazon S3 access point ARN and alias; for continuous recovery points it also echoes the AccessPointInTime that was supplied at create time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPointMetadataProperty := &AccessPointMetadataProperty{
//   	AccessPointInTime: jsii.String("accessPointInTime"),
//   	S3AccessPointAlias: jsii.String("s3AccessPointAlias"),
//   	S3AccessPointArn: jsii.String("s3AccessPointArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupaccesspoint-accesspointmetadata.html
//
type CfnBackupAccessPoint_AccessPointMetadataProperty struct {
	// For a continuous recovery point, the point in time the backup access point exposes, echoed back by the service as supplied at create time.
	//
	// Absent for snapshot recovery points. A plain string rather than a date-time, because it is one entry of the service's string-to-string metadata map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupaccesspoint-accesspointmetadata.html#cfn-backup-backupaccesspoint-accesspointmetadata-accesspointintime
	//
	AccessPointInTime *string `field:"optional" json:"accessPointInTime" yaml:"accessPointInTime"`
	// The alias of the Amazon S3 access point that Backup provisioned.
	//
	// Use this with standard Amazon S3 read APIs to access the backup data. Populated once the backup access point reaches the AVAILABLE status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupaccesspoint-accesspointmetadata.html#cfn-backup-backupaccesspoint-accesspointmetadata-s3accesspointalias
	//
	S3AccessPointAlias *string `field:"optional" json:"s3AccessPointAlias" yaml:"s3AccessPointAlias"`
	// The Amazon Resource Name (ARN) of the Amazon S3 access point that Backup provisioned.
	//
	// Use this with standard Amazon S3 read APIs to access the backup data. Populated once the backup access point reaches the AVAILABLE status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupaccesspoint-accesspointmetadata.html#cfn-backup-backupaccesspoint-accesspointmetadata-s3accesspointarn
	//
	S3AccessPointArn *string `field:"optional" json:"s3AccessPointArn" yaml:"s3AccessPointArn"`
}

