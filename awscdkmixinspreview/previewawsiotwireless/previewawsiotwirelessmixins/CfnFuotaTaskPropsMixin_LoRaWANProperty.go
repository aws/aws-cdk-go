package previewawsiotwirelessmixins


// The LoRaWAN information used with a FUOTA task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loRaWANProperty := &LoRaWANProperty{
//   	RfRegion: jsii.String("rfRegion"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html
//
type CfnFuotaTaskPropsMixin_LoRaWANProperty struct {
	// The frequency band (RFRegion) value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html#cfn-iotwireless-fuotatask-lorawan-rfregion
	//
	RfRegion *string `field:"optional" json:"rfRegion" yaml:"rfRegion"`
	// Start time of a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-fuotatask-lorawan.html#cfn-iotwireless-fuotatask-lorawan-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

