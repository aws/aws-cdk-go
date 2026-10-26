package awsmedialive


// The configured sources for this SRT input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtSettingsProperty := &SrtSettingsProperty{
//   	SrtCallerSources: []interface{}{
//   		&SrtCallerSourceProperty{
//   			Decryption: &SrtCallerDecryptionProperty{
//   				Algorithm: jsii.String("algorithm"),
//   				PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   			},
//   			MinimumLatency: jsii.Number(123),
//   			SrtListenerAddress: jsii.String("srtListenerAddress"),
//   			SrtListenerPort: jsii.String("srtListenerPort"),
//   			StreamId: jsii.String("streamId"),
//   		},
//   	},
//   	SrtListenerSettings: &SrtListenerSettingsProperty{
//   		Decryption: &SrtListenerDecryptionProperty{
//   			Algorithm: jsii.String("algorithm"),
//   			PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   		},
//   		MinimumLatency: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtsettings.html
//
type CfnInput_SrtSettingsProperty struct {
	// Placeholder documentation for __listOfSrtCallerSource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtsettings.html#cfn-medialive-input-srtsettings-srtcallersources
	//
	SrtCallerSources interface{} `field:"optional" json:"srtCallerSources" yaml:"srtCallerSources"`
	// The configuration for a source that uses SRT as the connection protocol.
	//
	// In terms of establishing the connection, MediaLive is the listener and the upstream system is the caller.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtsettings.html#cfn-medialive-input-srtsettings-srtlistenersettings
	//
	SrtListenerSettings interface{} `field:"optional" json:"srtListenerSettings" yaml:"srtListenerSettings"`
}

