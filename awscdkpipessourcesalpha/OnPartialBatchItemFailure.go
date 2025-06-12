package awscdkpipessourcesalpha


// Define how to handle item process failures.
// Experimental.
type OnPartialBatchItemFailure string

const (
	// EventBridge halves each batch and will retry each half until all the records are processed or there is one failed message left in the batch.
	// Experimental.
	OnPartialBatchItemFailure_AUTOMATIC_BISECT OnPartialBatchItemFailure = "AUTOMATIC_BISECT"
)

