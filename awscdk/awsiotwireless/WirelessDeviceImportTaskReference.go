package awsiotwireless


// A reference to a WirelessDeviceImportTask resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wirelessDeviceImportTaskReference := &WirelessDeviceImportTaskReference{
//   	WirelessDeviceImportTaskArn: jsii.String("wirelessDeviceImportTaskArn"),
//   	WirelessDeviceImportTaskId: jsii.String("wirelessDeviceImportTaskId"),
//   }
//
type WirelessDeviceImportTaskReference struct {
	// The ARN of the WirelessDeviceImportTask resource.
	WirelessDeviceImportTaskArn *string `field:"required" json:"wirelessDeviceImportTaskArn" yaml:"wirelessDeviceImportTaskArn"`
	// The Id of the WirelessDeviceImportTask resource.
	WirelessDeviceImportTaskId *string `field:"required" json:"wirelessDeviceImportTaskId" yaml:"wirelessDeviceImportTaskId"`
}

