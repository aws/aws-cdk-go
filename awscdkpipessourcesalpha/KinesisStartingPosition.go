package awscdkpipessourcesalpha


// The position in a Kinesis stream from which to start reading.
//
// Example:
//   var sourceStream stream
//   var targetQueue queue
//
//
//   pipeSource := sources.NewKinesisSource(sourceStream, &KinesisSourceParameters{
//   	StartingPosition: sources.KinesisStartingPosition_LATEST,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type KinesisStartingPosition string

const (
	// Start streaming at the last untrimmed record in the shard, which is the oldest data record in the shard.
	// Experimental.
	KinesisStartingPosition_TRIM_HORIZON KinesisStartingPosition = "TRIM_HORIZON"
	// Start streaming just after the most recent record in the shard, so that you always read the most recent data in the shard.
	// Experimental.
	KinesisStartingPosition_LATEST KinesisStartingPosition = "LATEST"
	// Start streaming from the position denoted by the time stamp specified in the `startingPositionTimestamp` field.
	// Experimental.
	KinesisStartingPosition_AT_TIMESTAMP KinesisStartingPosition = "AT_TIMESTAMP"
)

