package previewawsiotwirelessmixins


// Sidewalk-related information about a wireless device import task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sidewalkProperty := &SidewalkProperty{
//   	DeviceCreationFile: jsii.String("deviceCreationFile"),
//   	DeviceCreationFileList: []*string{
//   		jsii.String("deviceCreationFileList"),
//   	},
//   	Role: jsii.String("role"),
//   	SidewalkManufacturingSn: jsii.String("sidewalkManufacturingSn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html
//
type CfnWirelessDeviceImportTaskPropsMixin_SidewalkProperty struct {
	// The CSV file contained in an S3 bucket that's used for adding devices to an import task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-devicecreationfile
	//
	DeviceCreationFile *string `field:"optional" json:"deviceCreationFile" yaml:"deviceCreationFile"`
	// List of Sidewalk devices that are added to the import task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-devicecreationfilelist
	//
	DeviceCreationFileList *[]*string `field:"optional" json:"deviceCreationFileList" yaml:"deviceCreationFileList"`
	// The IAM role that allows  to access the CSV file in the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The Sidewalk manufacturing serial number (SMSN) of the Sidewalk device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-sidewalkmanufacturingsn
	//
	SidewalkManufacturingSn *string `field:"optional" json:"sidewalkManufacturingSn" yaml:"sidewalkManufacturingSn"`
}

