package mixinsawsmedialive


// The configuration of key provider settings.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyProviderSettingsProperty := &KeyProviderSettingsProperty{
//   	StaticKeySettings: &StaticKeySettingsProperty{
//   		KeyProviderServer: &InputLocationProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Uri: jsii.String("uri"),
//   			Username: jsii.String("username"),
//   		},
//   		StaticKeyValue: jsii.String("staticKeyValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-keyprovidersettings.html
//
type CfnChannelPropsMixin_KeyProviderSettingsProperty struct {
	// The configuration of static key settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-keyprovidersettings.html#cfn-medialive-channel-keyprovidersettings-statickeysettings
	//
	StaticKeySettings interface{} `field:"optional" json:"staticKeySettings" yaml:"staticKeySettings"`
}

