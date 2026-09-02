package awsmedialivealpha


// Audio normalization settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var audioNormalizationAlgorithm AudioNormalizationAlgorithm
//   var audioNormalizationAlgorithmControl AudioNormalizationAlgorithmControl
//   var audioNormalizationPeakCalculation AudioNormalizationPeakCalculation
//
//   audioNormalizationSettings := &AudioNormalizationSettings{
//   	Algorithm: audioNormalizationAlgorithm,
//   	AlgorithmControl: audioNormalizationAlgorithmControl,
//   	PeakCalculation: audioNormalizationPeakCalculation,
//   	PeakLimiterThreshold: jsii.Number(123),
//   	TargetLkfs: jsii.Number(123),
//   }
//
// Experimental.
type AudioNormalizationSettings struct {
	// The normalization algorithm.
	// Default: - service default.
	//
	// Experimental.
	Algorithm AudioNormalizationAlgorithm `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Whether to correct or only measure.
	// Default: - service default.
	//
	// Experimental.
	AlgorithmControl AudioNormalizationAlgorithmControl `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// Whether to use a peak limiter and how to calculate peak levels.
	// Default: - service default.
	//
	// Experimental.
	PeakCalculation AudioNormalizationPeakCalculation `field:"optional" json:"peakCalculation" yaml:"peakCalculation"`
	// The peak limiter threshold in dBFS.
	//
	// Only used when peak limiting is enabled.
	// Default: - service default.
	//
	// Experimental.
	PeakLimiterThreshold *float64 `field:"optional" json:"peakLimiterThreshold" yaml:"peakLimiterThreshold"`
	// The target loudness in LKFS.
	//
	// CALM Act recommends -24, EBU R-128 recommends -23.
	// Default: - service default.
	//
	// Experimental.
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

