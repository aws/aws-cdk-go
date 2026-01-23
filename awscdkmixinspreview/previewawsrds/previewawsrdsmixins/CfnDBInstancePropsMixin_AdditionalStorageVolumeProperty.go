package previewawsrdsmixins


// Contains details about an additional storage volume for a DB instance.
//
// RDS support additional storage volumes for RDS for Oracle and RDS for SQL Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalStorageVolumeProperty := &AdditionalStorageVolumeProperty{
//   	AllocatedStorage: jsii.String("allocatedStorage"),
//   	Iops: jsii.Number(123),
//   	MaxAllocatedStorage: jsii.Number(123),
//   	StorageThroughput: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	VolumeName: jsii.String("volumeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html
//
type CfnDBInstancePropsMixin_AdditionalStorageVolumeProperty struct {
	// The amount of storage allocated for the additional storage volume, in gibibytes (GiB).
	//
	// The minimum is 20 GiB. The maximum is 65,536 GiB (64 TiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-allocatedstorage
	//
	AllocatedStorage *string `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The number of I/O operations per second (IOPS) provisioned for the additional storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The upper limit in gibibytes (GiB) to which RDS can automatically scale the storage of the additional storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-maxallocatedstorage
	//
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The storage throughput value for the additional storage volume, in mebibytes per second (MiBps).
	//
	// This setting applies only to the General Purpose SSD gp3 storage type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-storagethroughput
	//
	StorageThroughput *float64 `field:"optional" json:"storageThroughput" yaml:"storageThroughput"`
	// The storage type for the additional storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The name of the additional storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-additionalstoragevolume.html#cfn-rds-dbinstance-additionalstoragevolume-volumename
	//
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

