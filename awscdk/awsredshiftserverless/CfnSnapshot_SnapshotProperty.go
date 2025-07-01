package awsredshiftserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnSnapshot_SnapshotProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-adminusername
	//
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-namespacearn
	//
	NamespaceArn *string `field:"optional" json:"namespaceArn" yaml:"namespaceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-owneraccount
	//
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-retentionperiod
	//
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotarn
	//
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotcreatetime
	//
	SnapshotCreateTime *string `field:"optional" json:"snapshotCreateTime" yaml:"snapshotCreateTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-snapshotname
	//
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-snapshot-snapshot.html#cfn-redshiftserverless-snapshot-snapshot-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

