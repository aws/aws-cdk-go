package previewawsiotwirelessmixins


// LoRaWANUpdateGatewayTaskEntry object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loRaWANUpdateGatewayTaskEntryProperty := &LoRaWANUpdateGatewayTaskEntryProperty{
//   	CurrentVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   	UpdateVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html
//
type CfnTaskDefinitionPropsMixin_LoRaWANUpdateGatewayTaskEntryProperty struct {
	// The version of the gateways that should receive the update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html#cfn-iotwireless-taskdefinition-lorawanupdategatewaytaskentry-currentversion
	//
	CurrentVersion interface{} `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// The firmware version to update the gateway to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.html#cfn-iotwireless-taskdefinition-lorawanupdategatewaytaskentry-updateversion
	//
	UpdateVersion interface{} `field:"optional" json:"updateVersion" yaml:"updateVersion"`
}

