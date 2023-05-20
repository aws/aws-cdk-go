package awscdk


// A Stack Synthesizer, obtained from `IReusableStackSynthesizer.`.
//
// Just a type alias with a very concrete contract.
type IBoundStackSynthesizer interface {
	IStackSynthesizer
}

// The jsii proxy for IBoundStackSynthesizer
type jsiiProxy_IBoundStackSynthesizer struct {
	jsiiProxy_IStackSynthesizer
}

