package awsiotwireless


// UpdateWirelessGatewayTaskCreate object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateWirelessGatewayTaskCreateProperty := &UpdateWirelessGatewayTaskCreateProperty{
//   	LoRaWan: &LoRaWANUpdateGatewayTaskCreateProperty{
//   		CurrentVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   		SigKeyCrc: jsii.Number(123),
//   		UpdateSignature: jsii.String("updateSignature"),
//   		UpdateVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   	},
//   	UpdateDataRole: jsii.String("updateDataRole"),
//   	UpdateDataSource: jsii.String("updateDataSource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html
//
type CfnTaskDefinition_UpdateWirelessGatewayTaskCreateProperty struct {
	// The properties that relate to the LoRaWAN wireless gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-lorawan
	//
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The IAM role used to read data from the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-updatedatarole
	//
	UpdateDataRole *string `field:"optional" json:"updateDataRole" yaml:"updateDataRole"`
	// The link to the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.html#cfn-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate-updatedatasource
	//
	UpdateDataSource *string `field:"optional" json:"updateDataSource" yaml:"updateDataSource"`
}

