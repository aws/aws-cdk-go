package awsmedialive


// The configuration of key provider settings.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyProviderSettingsProperty := &keyProviderSettingsProperty{
//   	staticKeySettings: &staticKeySettingsProperty{
//   		keyProviderServer: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		staticKeyValue: jsii.String("staticKeyValue"),
//   	},
//   }
//
type CfnChannel_KeyProviderSettingsProperty struct {
	// The configuration of static key settings.
	StaticKeySettings interface{} `field:"optional" json:"staticKeySettings" yaml:"staticKeySettings"`
}

