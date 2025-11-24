package awsimagebuilderalpha


// The parameter type for the workflow parameter.
// Experimental.
type WorkflowParameterType string

const (
	// Indicates the workflow parameter has a boolean value.
	// Experimental.
	WorkflowParameterType_BOOLEAN WorkflowParameterType = "BOOLEAN"
	// Indicates the workflow parameter has an integer value.
	// Experimental.
	WorkflowParameterType_INTEGER WorkflowParameterType = "INTEGER"
	// Indicates the workflow parameter has a string value.
	// Experimental.
	WorkflowParameterType_STRING WorkflowParameterType = "STRING"
	// Indicates the workflow parameter has a string list value.
	// Experimental.
	WorkflowParameterType_STRING_LIST WorkflowParameterType = "STRING_LIST"
)

