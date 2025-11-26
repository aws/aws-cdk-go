package previewawsredshiftserverlessmixins


// A snapshot object that contains databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotProperty := &SnapshotProperty{
//   	AdminUsername: jsii.String("adminUsername"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	NamespaceArn: jsii.String("namespaceArn"),
//   	NamespaceName: jsii.String("namespaceName"),
//   	OwnerAccount: jsii.String("ownerAccount"),
//   	RetentionPeriod: jsii.Number(123),
//   	SnapshotArn: jsii.String("snapshotArn"),
//   	SnapshotCreateTime: jsii.String("snapshotCreateTime"),
//   	SnapshotName: jsii.String("snapshotName"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html
//
type CfnSnapshotPropsMixin_SnapshotProperty struct {
	// The username of the database within a snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-adminusername
	//
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// The unique identifier of the KMS key used to encrypt the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-namespacearn
	//
	NamespaceArn *string `field:"optional" json:"namespaceArn" yaml:"namespaceArn"`
	// The name of the namepsace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The owner AWS ;
	//
	// account of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-owneraccount
	//
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// The retention period of the snapshot created by the scheduled action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-retentionperiod
	//
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The Amazon Resource Name (ARN) of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotarn
	//
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// The timestamp of when the snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotcreatetime
	//
	SnapshotCreateTime *string `field:"optional" json:"snapshotCreateTime" yaml:"snapshotCreateTime"`
	// The name of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotname
	//
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The status of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

