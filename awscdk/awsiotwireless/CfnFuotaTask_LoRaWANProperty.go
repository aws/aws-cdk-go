package awsiotwireless


// The LoRaWAN information used with a FUOTA task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANProperty := &LoRaWANProperty{
//   	RfRegion: jsii.String("rfRegion"),
//
//   	// the properties below are optional
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html
//
type CfnFuotaTask_LoRaWANProperty struct {
	// The frequency band (RFRegion) value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html#cfn-iotwireless-fuotatask-lorawan-rfregion
	//
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
	// Start time of a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html#cfn-iotwireless-fuotatask-lorawan-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

