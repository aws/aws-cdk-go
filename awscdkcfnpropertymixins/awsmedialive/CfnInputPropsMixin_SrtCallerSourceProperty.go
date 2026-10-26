package awsmedialive


// The configuration for a source that uses SRT as the connection protocol.
//
// In terms of establishing the connection, MediaLive is always caller and the upstream system is always the listener. In terms of transmission of the source content, MediaLive is always the receiver and the upstream system is always the sender.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   srtCallerSourceProperty := &SrtCallerSourceProperty{
//   	Decryption: &SrtCallerDecryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		PassphraseSecretArn: jsii.String("passphraseSecretArn"),
//   	},
//   	MinimumLatency: jsii.Number(123),
//   	SrtListenerAddress: jsii.String("srtListenerAddress"),
//   	SrtListenerPort: jsii.String("srtListenerPort"),
//   	StreamId: jsii.String("streamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html
//
type CfnInputPropsMixin_SrtCallerSourceProperty struct {
	// The decryption settings for the SRT caller source.
	//
	// Present only if the source has decryption enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html#cfn-medialive-input-srtcallersource-decryption
	//
	Decryption interface{} `field:"optional" json:"decryption" yaml:"decryption"`
	// The preferred latency (in milliseconds) for implementing packet loss and recovery.
	//
	// Packet recovery is a key feature of SRT.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html#cfn-medialive-input-srtcallersource-minimumlatency
	//
	MinimumLatency *float64 `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// The IP address at the upstream system (the listener) that MediaLive (the caller) connects to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html#cfn-medialive-input-srtcallersource-srtlisteneraddress
	//
	SrtListenerAddress *string `field:"optional" json:"srtListenerAddress" yaml:"srtListenerAddress"`
	// The port at the upstream system (the listener) that MediaLive (the caller) connects to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html#cfn-medialive-input-srtcallersource-srtlistenerport
	//
	SrtListenerPort *string `field:"optional" json:"srtListenerPort" yaml:"srtListenerPort"`
	// The stream ID, if the upstream system uses this identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-srtcallersource.html#cfn-medialive-input-srtcallersource-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

