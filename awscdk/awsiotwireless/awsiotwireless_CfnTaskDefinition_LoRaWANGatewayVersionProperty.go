package awsiotwireless


// LoRaWANGatewayVersion object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANGatewayVersionProperty := &loRaWANGatewayVersionProperty{
//   	model: jsii.String("model"),
//   	packageVersion: jsii.String("packageVersion"),
//   	station: jsii.String("station"),
//   }
//
type CfnTaskDefinition_LoRaWANGatewayVersionProperty struct {
	// The model number of the wireless gateway.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The version of the wireless gateway firmware.
	PackageVersion *string `field:"optional" json:"packageVersion" yaml:"packageVersion"`
	// The basic station version of the wireless gateway.
	Station *string `field:"optional" json:"station" yaml:"station"`
}

