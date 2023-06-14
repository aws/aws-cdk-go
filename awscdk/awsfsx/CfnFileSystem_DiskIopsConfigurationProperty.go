package awsfsx


// The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP or FSx for OpenZFS file system.
//
// By default, Amazon FSx automatically provisions 3 IOPS per GB of storage capacity. You can provision additional IOPS per GB of storage. The configuration consists of the total number of provisioned SSD IOPS and how it is was provisioned, or the mode (by the customer or by Amazon FSx).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskIopsConfigurationProperty := &DiskIopsConfigurationProperty{
//   	Iops: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   }
//
type CfnFileSystem_DiskIopsConfigurationProperty struct {
	// The total number of SSD IOPS provisioned for the file system.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Specifies whether the file system is using the `AUTOMATIC` setting of SSD IOPS of 3 IOPS per GB of storage capacity, , or if it using a `USER_PROVISIONED` value.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

