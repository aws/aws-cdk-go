package awsiotwireless


// The LoRaWAN information used with a FUOTA task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANProperty := &loRaWANProperty{
//   	rfRegion: jsii.String("rfRegion"),
//
//   	// the properties below are optional
//   	startTime: jsii.String("startTime"),
//   }
//
type CfnFuotaTask_LoRaWANProperty struct {
	// The frequency band (RFRegion) value.
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
	// Start time of a FUOTA task.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

