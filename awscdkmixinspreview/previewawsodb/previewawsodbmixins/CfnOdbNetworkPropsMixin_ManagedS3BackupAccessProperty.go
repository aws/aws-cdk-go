package previewawsodbmixins


// The configuration for managed Amazon S3 backup access from the ODB network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedS3BackupAccessProperty := &ManagedS3BackupAccessProperty{
//   	Ipv4Addresses: []*string{
//   		jsii.String("ipv4Addresses"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-manageds3backupaccess.html
//
type CfnOdbNetworkPropsMixin_ManagedS3BackupAccessProperty struct {
	// The IPv4 addresses for the managed Amazon S3 backup access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-manageds3backupaccess.html#cfn-odb-odbnetwork-manageds3backupaccess-ipv4addresses
	//
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// The status of the managed Amazon S3 backup access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-manageds3backupaccess.html#cfn-odb-odbnetwork-manageds3backupaccess-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

