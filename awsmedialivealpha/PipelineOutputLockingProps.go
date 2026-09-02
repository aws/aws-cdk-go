package awsmedialivealpha


// Properties for pipeline output locking.
//
// Example:
//   // Video-aligned pipeline locking — useful when sources lack reliable timecodes
//   locking := medialive.OutputLocking_Pipeline(&PipelineOutputLockingProps{
//   	Method: medialive.PipelineLockingMethod_VIDEO_ALIGNMENT(),
//   })
//
// Experimental.
type PipelineOutputLockingProps struct {
	// A custom epoch (ISO-8601 timestamp) to lock outputs to.
	// Default: - service default.
	//
	// Experimental.
	CustomEpoch *string `field:"optional" json:"customEpoch" yaml:"customEpoch"`
	// The method MediaLive uses to synchronise the pipelines.
	// Default: - SOURCE_TIMECODE, applied by MediaLive.
	//
	// Experimental.
	Method PipelineLockingMethod `field:"optional" json:"method" yaml:"method"`
}

