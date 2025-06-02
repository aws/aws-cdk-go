package awsiotwireless


// List of FPorts assigned for different LoRaWAN application packages to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fPortsProperty := &FPortsProperty{
//   	Applications: []interface{}{
//   		&ApplicationProperty{
//   			DestinationName: jsii.String("destinationName"),
//   			FPort: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-fports.html
//
type CfnWirelessDevice_FPortsProperty struct {
	// LoRaWAN application configuration, which can be used to perform geolocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-fports.html#cfn-iotwireless-wirelessdevice-fports-applications
	//
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
}

