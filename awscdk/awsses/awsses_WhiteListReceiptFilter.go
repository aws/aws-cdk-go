package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

// An allow list receipt filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   whiteListReceiptFilter := awscdk.Aws_ses.NewWhiteListReceiptFilter(this, jsii.String("MyWhiteListReceiptFilter"), &whiteListReceiptFilterProps{
//   	ips: []*string{
//   		jsii.String("ips"),
//   	},
//   })
//
// Deprecated: use `AllowListReceiptFilter`.
type WhiteListReceiptFilter interface {
	AllowListReceiptFilter
	// The construct tree node associated with this construct.
	// Deprecated: use `AllowListReceiptFilter`.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use `AllowListReceiptFilter`.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use `AllowListReceiptFilter`.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use `AllowListReceiptFilter`.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use `AllowListReceiptFilter`.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use `AllowListReceiptFilter`.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: use `AllowListReceiptFilter`.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use `AllowListReceiptFilter`.
	Validate() *[]*string
}

// The jsii proxy struct for WhiteListReceiptFilter
type jsiiProxy_WhiteListReceiptFilter struct {
	jsiiProxy_AllowListReceiptFilter
}

func (j *jsiiProxy_WhiteListReceiptFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: use `AllowListReceiptFilter`.
func NewWhiteListReceiptFilter(scope constructs.Construct, id *string, props *WhiteListReceiptFilterProps) WhiteListReceiptFilter {
	_init_.Initialize()

	if err := validateNewWhiteListReceiptFilterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WhiteListReceiptFilter{}

	_jsii_.Create(
		"monocdk.aws_ses.WhiteListReceiptFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use `AllowListReceiptFilter`.
func NewWhiteListReceiptFilter_Override(w WhiteListReceiptFilter, scope constructs.Construct, id *string, props *WhiteListReceiptFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses.WhiteListReceiptFilter",
		[]interface{}{scope, id, props},
		w,
	)
}

// Return whether the given object is a Construct.
// Deprecated: use `AllowListReceiptFilter`.
func WhiteListReceiptFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWhiteListReceiptFilter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.WhiteListReceiptFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WhiteListReceiptFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WhiteListReceiptFilter) OnSynthesize(session constructs.ISynthesisSession) {
	if err := w.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WhiteListReceiptFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WhiteListReceiptFilter) Prepare() {
	_jsii_.InvokeVoid(
		w,
		"prepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WhiteListReceiptFilter) Synthesize(session awscdk.ISynthesisSession) {
	if err := w.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"synthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WhiteListReceiptFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WhiteListReceiptFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

