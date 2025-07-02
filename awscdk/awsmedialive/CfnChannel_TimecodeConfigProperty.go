package awsmedialive


// The configuration of the timecode in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timecodeConfigProperty := &TimecodeConfigProperty{
//   	Source: jsii.String("source"),
//   	SyncThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeconfig.html
//
type CfnChannel_TimecodeConfigProperty struct {
	// Identifies the source for the timecode that will be associated with the channel outputs.
	//
	// Embedded (embedded): Initialize the output timecode with timecode from the source. If no embedded timecode is detected in the source, the system falls back to using "Start at 0" (zerobased). System Clock (systemclock): Use the UTC time. Start at 0 (zerobased): The time of the first frame of the channel will be 00:00:00:00.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeconfig.html#cfn-medialive-channel-timecodeconfig-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The threshold in frames beyond which output timecode is resynchronized to the input timecode.
	//
	// Discrepancies below this threshold are permitted to avoid unnecessary discontinuities in the output timecode. There is no timecode sync when this is not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-timecodeconfig.html#cfn-medialive-channel-timecodeconfig-syncthreshold
	//
	SyncThreshold *float64 `field:"optional" json:"syncThreshold" yaml:"syncThreshold"`
}

