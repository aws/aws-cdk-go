package mixinsawsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multicastInputSettingsProperty := &MulticastInputSettingsProperty{
//   	SourceIpAddress: jsii.String("sourceIpAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multicastinputsettings.html
//
type CfnChannelPropsMixin_MulticastInputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multicastinputsettings.html#cfn-medialive-channel-multicastinputsettings-sourceipaddress
	//
	SourceIpAddress *string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
}

