package awscdkpipesalpha


// Reserved pipe variables.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-input-transformation.html#input-transform-reserved
//
// Experimental.
type PipeVariable string

const (
	// The Amazon Resource Name (ARN) of the pipe.
	// Experimental.
	PipeVariable_ARN PipeVariable = "ARN"
	// The name of the pipe.
	// Experimental.
	PipeVariable_NAME PipeVariable = "NAME"
	// The ARN of the event source of the pipe.
	// Experimental.
	PipeVariable_SOURCE_ARN PipeVariable = "SOURCE_ARN"
	// The ARN of the enrichment of the pipe.
	// Experimental.
	PipeVariable_ENRICHMENT_ARN PipeVariable = "ENRICHMENT_ARN"
	// The ARN of the target of the pipe.
	// Experimental.
	PipeVariable_TARGET_ARN PipeVariable = "TARGET_ARN"
	// The time at which the event was received by the input transformer.
	//
	// This is an ISO 8601 timestamp. This time is different for the enrichment input transformer and the target input transformer, depending on when the enrichment completed processing the event.
	// Experimental.
	PipeVariable_EVENT_INGESTION_TIME PipeVariable = "EVENT_INGESTION_TIME"
	// The event as received by the input transformer.
	// Experimental.
	PipeVariable_EVENT PipeVariable = "EVENT"
	// The same as aws.pipes.event, but the variable only has a value if the original payload, either from the source or returned by the enrichment, is JSON. If the pipe has an encoded field, such as the Amazon SQS body field or the Kinesis data, those fields are decoded and turned into valid JSON. Because it isn't escaped, the variable can only be used as a value for a JSON field. For more information, see Implicit body data parsing.
	// Experimental.
	PipeVariable_EVENT_JSON PipeVariable = "EVENT_JSON"
)

