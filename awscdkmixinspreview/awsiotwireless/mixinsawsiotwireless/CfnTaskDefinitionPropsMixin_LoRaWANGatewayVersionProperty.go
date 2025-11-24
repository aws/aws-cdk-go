package mixinsawsiotwireless


// LoRaWANGatewayVersion object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loRaWANGatewayVersionProperty := &LoRaWANGatewayVersionProperty{
//   	Model: jsii.String("model"),
//   	PackageVersion: jsii.String("packageVersion"),
//   	Station: jsii.String("station"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html
//
type CfnTaskDefinitionPropsMixin_LoRaWANGatewayVersionProperty struct {
	// The model number of the wireless gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html#cfn-iotwireless-taskdefinition-lorawangatewayversion-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The version of the wireless gateway firmware.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html#cfn-iotwireless-taskdefinition-lorawangatewayversion-packageversion
	//
	PackageVersion *string `field:"optional" json:"packageVersion" yaml:"packageVersion"`
	// The basic station version of the wireless gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html#cfn-iotwireless-taskdefinition-lorawangatewayversion-station
	//
	Station *string `field:"optional" json:"station" yaml:"station"`
}

