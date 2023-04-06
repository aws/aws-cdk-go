package awsiotwireless


// The signature used to verify the update firmware.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANUpdateGatewayTaskCreateProperty := &LoRaWANUpdateGatewayTaskCreateProperty{
//   	CurrentVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   	SigKeyCrc: jsii.Number(123),
//   	UpdateSignature: jsii.String("updateSignature"),
//   	UpdateVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   }
//
type CfnTaskDefinition_LoRaWANUpdateGatewayTaskCreateProperty struct {
	// The version of the gateways that should receive the update.
	CurrentVersion interface{} `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// The CRC of the signature private key to check.
	SigKeyCrc *float64 `field:"optional" json:"sigKeyCrc" yaml:"sigKeyCrc"`
	// The signature used to verify the update firmware.
	UpdateSignature *string `field:"optional" json:"updateSignature" yaml:"updateSignature"`
	// The firmware version to update the gateway to.
	UpdateVersion interface{} `field:"optional" json:"updateVersion" yaml:"updateVersion"`
}

