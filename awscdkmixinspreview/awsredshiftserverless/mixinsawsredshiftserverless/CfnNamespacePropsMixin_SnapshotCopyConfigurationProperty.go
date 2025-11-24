package mixinsawsredshiftserverless


// The object that you configure to copy snapshots from one namespace to a namespace in another AWS Region .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snapshotCopyConfigurationProperty := &SnapshotCopyConfigurationProperty{
//   	DestinationKmsKeyId: jsii.String("destinationKmsKeyId"),
//   	DestinationRegion: jsii.String("destinationRegion"),
//   	SnapshotRetentionPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html
//
type CfnNamespacePropsMixin_SnapshotCopyConfigurationProperty struct {
	// The ID of the KMS key to use to encrypt your snapshots in the destination AWS Region .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-destinationkmskeyid
	//
	DestinationKmsKeyId *string `field:"optional" json:"destinationKmsKeyId" yaml:"destinationKmsKeyId"`
	// The destination AWS Region to copy snapshots to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-destinationregion
	//
	DestinationRegion *string `field:"optional" json:"destinationRegion" yaml:"destinationRegion"`
	// The retention period of snapshots that are copied to the destination AWS Region .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-snapshotretentionperiod
	//
	SnapshotRetentionPeriod *float64 `field:"optional" json:"snapshotRetentionPeriod" yaml:"snapshotRetentionPeriod"`
}

