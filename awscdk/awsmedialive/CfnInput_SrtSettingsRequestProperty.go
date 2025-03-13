package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtSettingsRequestProperty := &SrtSettingsRequestProperty{
//   	SrtCallerSources: []interface{}{
//   		&SrtCallerSourceRequestProperty{
//   			Decryption: &SrtCallerDecryptionRequestProperty{
//   				Algorithm: jsii.String("algorithm"),
//   				PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   			},
//   			MinimumLatency: jsii.Number(123),
//   			SrtListenerAddress: jsii.String("srtListenerAddress"),
//   			SrtListenerPort: jsii.String("srtListenerPort"),
//   			StreamId: jsii.String("streamId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtsettingsrequest.html
//
type CfnInput_SrtSettingsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtsettingsrequest.html#cfn-medialive-input-srtsettingsrequest-srtcallersources
	//
	SrtCallerSources interface{} `field:"optional" json:"srtCallerSources" yaml:"srtCallerSources"`
}

