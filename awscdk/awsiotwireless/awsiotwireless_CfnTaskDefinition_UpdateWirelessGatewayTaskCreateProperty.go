package awsiotwireless


// UpdateWirelessGatewayTaskCreate object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateWirelessGatewayTaskCreateProperty := &updateWirelessGatewayTaskCreateProperty{
//   	loRaWan: &loRaWANUpdateGatewayTaskCreateProperty{
//   		currentVersion: &loRaWANGatewayVersionProperty{
//   			model: jsii.String("model"),
//   			packageVersion: jsii.String("packageVersion"),
//   			station: jsii.String("station"),
//   		},
//   		sigKeyCrc: jsii.Number(123),
//   		updateSignature: jsii.String("updateSignature"),
//   		updateVersion: &loRaWANGatewayVersionProperty{
//   			model: jsii.String("model"),
//   			packageVersion: jsii.String("packageVersion"),
//   			station: jsii.String("station"),
//   		},
//   	},
//   	updateDataRole: jsii.String("updateDataRole"),
//   	updateDataSource: jsii.String("updateDataSource"),
//   }
//
type CfnTaskDefinition_UpdateWirelessGatewayTaskCreateProperty struct {
	// The properties that relate to the LoRaWAN wireless gateway.
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The IAM role used to read data from the S3 bucket.
	UpdateDataRole *string `field:"optional" json:"updateDataRole" yaml:"updateDataRole"`
	// The link to the S3 bucket.
	UpdateDataSource *string `field:"optional" json:"updateDataSource" yaml:"updateDataSource"`
}

