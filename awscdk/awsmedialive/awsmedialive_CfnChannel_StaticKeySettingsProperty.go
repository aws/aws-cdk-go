package awsmedialive


// The static key settings.
//
// The parent of this entity is KeyProviderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticKeySettingsProperty := &staticKeySettingsProperty{
//   	keyProviderServer: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	staticKeyValue: jsii.String("staticKeyValue"),
//   }
//
type CfnChannel_StaticKeySettingsProperty struct {
	// The URL of the license server that is used for protecting content.
	KeyProviderServer interface{} `field:"optional" json:"keyProviderServer" yaml:"keyProviderServer"`
	// The static key value as a 32 character hexadecimal string.
	StaticKeyValue *string `field:"optional" json:"staticKeyValue" yaml:"staticKeyValue"`
}

