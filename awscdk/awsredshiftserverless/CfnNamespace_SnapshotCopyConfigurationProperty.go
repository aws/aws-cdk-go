package awsredshiftserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotCopyConfigurationProperty := &SnapshotCopyConfigurationProperty{
//   	DestinationRegion: jsii.String("destinationRegion"),
//
//   	// the properties below are optional
//   	DestinationKmsKeyId: jsii.String("destinationKmsKeyId"),
//   	SnapshotRetentionPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html
//
type CfnNamespace_SnapshotCopyConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-destinationregion
	//
	DestinationRegion *string `field:"required" json:"destinationRegion" yaml:"destinationRegion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-destinationkmskeyid
	//
	DestinationKmsKeyId *string `field:"optional" json:"destinationKmsKeyId" yaml:"destinationKmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-namespace-snapshotcopyconfiguration.html#cfn-redshiftserverless-namespace-snapshotcopyconfiguration-snapshotretentionperiod
	//
	SnapshotRetentionPeriod *float64 `field:"optional" json:"snapshotRetentionPeriod" yaml:"snapshotRetentionPeriod"`
}

