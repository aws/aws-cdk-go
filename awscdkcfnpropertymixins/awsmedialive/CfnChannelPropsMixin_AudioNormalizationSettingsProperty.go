package awsmedialive


// The settings for normalizing video.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audioNormalizationSettingsProperty := &AudioNormalizationSettingsProperty{
//   	Algorithm: jsii.String("algorithm"),
//   	AlgorithmControl: jsii.String("algorithmControl"),
//   	PeakCalculation: jsii.String("peakCalculation"),
//   	PeakLimiterThreshold: jsii.Number(123),
//   	TargetLkfs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html
//
type CfnChannelPropsMixin_AudioNormalizationSettingsProperty struct {
	// The audio normalization algorithm to use.
	//
	// itu17701 conforms to the CALM Act specification. itu17702 conforms to the EBU R-128 specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html#cfn-medialive-channel-audionormalizationsettings-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// When set to correctAudio, the output audio is corrected using the chosen algorithm.
	//
	// If set to measureOnly, the audio is measured but not adjusted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html#cfn-medialive-channel-audionormalizationsettings-algorithmcontrol
	//
	AlgorithmControl *string `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html#cfn-medialive-channel-audionormalizationsettings-peakcalculation
	//
	PeakCalculation *string `field:"optional" json:"peakCalculation" yaml:"peakCalculation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html#cfn-medialive-channel-audionormalizationsettings-peaklimiterthreshold
	//
	PeakLimiterThreshold *float64 `field:"optional" json:"peakLimiterThreshold" yaml:"peakLimiterThreshold"`
	// The Target LKFS(loudness) to adjust volume to.
	//
	// If no value is entered, a default value is used according to the chosen algorithm. The CALM Act (1770-1) recommends a target of -24 LKFS. The EBU R-128 specification (1770-2) recommends a target of -23 LKFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audionormalizationsettings.html#cfn-medialive-channel-audionormalizationsettings-targetlkfs
	//
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

