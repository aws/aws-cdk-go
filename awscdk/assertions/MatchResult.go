package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The result of `Match.test()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var target interface{}
//
//   matchResult := awscdk.Assertions.NewMatchResult(target)
//
type MatchResult interface {
	// The cost of the failures so far.
	FailCost() *float64
	// The number of failures.
	FailCount() *float64
	// Whether the match is a success.
	IsSuccess() *bool
	// The target for which this result was generated.
	Target() interface{}
	// Compose the results of a previous match as a subtree.
	Compose(id *string, inner MatchResult) MatchResult
	// Prepare the result to be analyzed.
	//
	// This API *must* be called prior to analyzing these results.
	Finished() MatchResult
	// Does the result contain any failures.
	//
	// If not, the result is a success.
	HasFailed() *bool
	// DEPRECATED.
	// Deprecated: use recordFailure().
	Push(matcher Matcher, path *[]*string, message *string) MatchResult
	// Record a capture against in this match result.
	RecordCapture(options *MatchCapture)
	// Record a new failure into this result at a specific path.
	RecordFailure(failure *MatchFailure) MatchResult
	// Do a deep render of the match result, showing the structure mismatches in context.
	RenderMismatch() *string
	// Render the failed match in a presentable way.
	//
	// Prefer using `renderMismatch` over this method. It is left for backwards
	// compatibility for test suites that expect it, but `renderMismatch()` will
	// produce better output.
	ToHumanStrings() *[]*string
}

// The jsii proxy struct for MatchResult
type jsiiProxy_MatchResult struct {
	_ byte // padding
}

func (j *jsiiProxy_MatchResult) FailCost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failCost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchResult) FailCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchResult) IsSuccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchResult) Target() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


func NewMatchResult(target interface{}) MatchResult {
	_init_.Initialize()

	if err := validateNewMatchResultParameters(target); err != nil {
		panic(err)
	}
	j := jsiiProxy_MatchResult{}

	_jsii_.Create(
		"aws-cdk-lib.assertions.MatchResult",
		[]interface{}{target},
		&j,
	)

	return &j
}

func NewMatchResult_Override(m MatchResult, target interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.MatchResult",
		[]interface{}{target},
		m,
	)
}

func (m *jsiiProxy_MatchResult) Compose(id *string, inner MatchResult) MatchResult {
	if err := m.validateComposeParameters(id, inner); err != nil {
		panic(err)
	}
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"compose",
		[]interface{}{id, inner},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) Finished() MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"finished",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) HasFailed() *bool {
	var returns *bool

	_jsii_.Invoke(
		m,
		"hasFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) Push(matcher Matcher, path *[]*string, message *string) MatchResult {
	if err := m.validatePushParameters(matcher, path, message); err != nil {
		panic(err)
	}
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"push",
		[]interface{}{matcher, path, message},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) RecordCapture(options *MatchCapture) {
	if err := m.validateRecordCaptureParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"recordCapture",
		[]interface{}{options},
	)
}

func (m *jsiiProxy_MatchResult) RecordFailure(failure *MatchFailure) MatchResult {
	if err := m.validateRecordFailureParameters(failure); err != nil {
		panic(err)
	}
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"recordFailure",
		[]interface{}{failure},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) RenderMismatch() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"renderMismatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchResult) ToHumanStrings() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"toHumanStrings",
		nil, // no parameters
		&returns,
	)

	return returns
}

