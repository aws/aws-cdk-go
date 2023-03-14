package awsmedialive


// The settings for normalizing video.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioNormalizationSettingsProperty := &audioNormalizationSettingsProperty{
//   	algorithm: jsii.String("algorithm"),
//   	algorithmControl: jsii.String("algorithmControl"),
//   	targetLkfs: jsii.Number(123),
//   }
//
type CfnChannel_AudioNormalizationSettingsProperty struct {
	// The audio normalization algorithm to use.
	//
	// itu17701 conforms to the CALM Act specification. itu17702 conforms to the EBU R-128 specification.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// When set to correctAudio, the output audio is corrected using the chosen algorithm.
	//
	// If set to measureOnly, the audio is measured but not adjusted.
	AlgorithmControl *string `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// The Target LKFS(loudness) to adjust volume to.
	//
	// If no value is entered, a default value is used according to the chosen algorithm. The CALM Act (1770-1) recommends a target of -24 LKFS. The EBU R-128 specification (1770-2) recommends a target of -23 LKFS.
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

