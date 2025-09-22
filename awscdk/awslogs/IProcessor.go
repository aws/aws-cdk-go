package awslogs


// Interface representing a single processor in a CloudWatch Logs transformer.
//
// A log transformer is a series of processors, where each processor applies one type of transformation
// to the log events. The processors work one after another, in the order that they are listed, like a pipeline.
type IProcessor interface {
}

// The jsii proxy for IProcessor
type jsiiProxy_IProcessor struct {
	_ byte // padding
}

