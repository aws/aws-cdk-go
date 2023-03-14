package awsfsx


// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS file system.
//
// The default is 3 IOPS per GB of storage capacity, but you can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how the amount was provisioned (by the customer or by the system).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskIopsConfigurationProperty := &diskIopsConfigurationProperty{
//   	iops: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   }
//
type CfnFileSystem_DiskIopsConfigurationProperty struct {
	// The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Specifies whether the number of IOPS for the file system is using the system default ( `AUTOMATIC` ) or was provisioned by the customer ( `USER_PROVISIONED` ).
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

