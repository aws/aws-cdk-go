package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
)

// An event pattern matcher.
//
// Example:
//   rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
//   	EventPattern: &EventPattern{
//   		Detail: map[string]interface{}{
//   			"object": map[string][]*string{
//   				// Matchers may appear at any level
//   				"size": events.Match_greaterThan(jsii.Number(1024)),
//   			},
//
//   			// 'OR' condition
//   			"source-storage-class": events.Match_anyOf(events.Match_prefix(jsii.String("GLACIER")), events.Match_exactString(jsii.String("DEEP_ARCHIVE"))),
//   		},
//   		DetailType: events.Match_EqualsIgnoreCase(jsii.String("object created")),
//
//   		// If you prefer, you can use a low level array of strings, as directly consumed by EventBridge
//   		Source: []*string{
//   			jsii.String("aws.s3"),
//   		},
//
//   		Region: events.Match_AnythingButPrefix(jsii.String("us-gov")),
//   	},
//   })
//
type Match interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// A representation of this matcher as a list of strings.
	AsList() *[]*string
	// Produce the Token's value at resolution time.
	Resolve(context awscdk.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	ToString() *string
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_Match) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


// Matches an event if any of the provided matchers do.
//
// Only numeric matchers are accepted.
func Match_AllOf(matchers ...interface{}) *[]*string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range matchers {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"allOf",
		args,
		&returns,
	)

	return returns
}

// Matches an event if any of the provided matchers does.
func Match_AnyOf(matchers ...interface{}) *[]*string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range matchers {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"anyOf",
		args,
		&returns,
	)

	return returns
}

// Matches anything except what's provided in the rule.
//
// The list of provided values must contain
// only strings or only numbers.
func Match_AnythingBut(values ...interface{}) *[]*string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range values {
		args = append(args, a)
	}

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"anythingBut",
		args,
		&returns,
	)

	return returns
}

// Matches any string that doesn't start with the given prefix.
func Match_AnythingButPrefix(prefix *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_AnythingButPrefixParameters(prefix); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"anythingButPrefix",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

// Matches IPv4 and IPv6 network addresses using the Classless Inter-Domain Routing (CIDR) format.
func Match_Cidr(range_ *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_CidrParameters(range_); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"cidr",
		[]interface{}{range_},
		&returns,
	)

	return returns
}

// Matches when the field is absent from the JSON of the event.
func Match_DoesNotExist() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"doesNotExist",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches numbers equal to the provided value.
func Match_Equal(value *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_EqualParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"equal",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches a string, regardless of case, in the JSON of the event.
func Match_EqualsIgnoreCase(value *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_EqualsIgnoreCaseParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"equalsIgnoreCase",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches a string, exactly, in the JSON of the event.
func Match_ExactString(value *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_ExactStringParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"exactString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches when the field is present in the JSON of the event.
func Match_Exists() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"exists",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches numbers greater than the provided value.
func Match_GreaterThan(value *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_GreaterThanParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"greaterThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches numbers greater than, or equal to, the provided value.
func Match_GreaterThanOrEqual(value *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_GreaterThanOrEqualParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"greaterThanOrEqual",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches numbers inside a closed numeric interval. Equivalent to:.
//
// Match.allOf(Match.greaterThanOrEqual(lower), Match.lessThanOrEqual(upper))
func Match_Interval(lower *float64, upper *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_IntervalParameters(lower, upper); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"interval",
		[]interface{}{lower, upper},
		&returns,
	)

	return returns
}

// Matches IPv4 and IPv6 network addresses using the Classless Inter-Domain Routing (CIDR) format.
//
// Alias of `cidr()`.
func Match_IpAddressRange(range_ *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_IpAddressRangeParameters(range_); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"ipAddressRange",
		[]interface{}{range_},
		&returns,
	)

	return returns
}

// Matches numbers less than the provided value.
func Match_LessThan(value *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_LessThanParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"lessThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches numbers less than, or equal to, the provided value.
func Match_LessThanOrEqual(value *float64) *[]*string {
	_init_.Initialize()

	if err := validateMatch_LessThanOrEqualParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"lessThanOrEqual",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches a null value in the JSON of the event.
func Match_NullValue() *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"nullValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches strings with the given prefix in the JSON of the event.
func Match_Prefix(value *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_PrefixParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"prefix",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Matches strings with the given suffix in the JSON of the event.
func Match_Suffix(value *string) *[]*string {
	_init_.Initialize()

	if err := validateMatch_SuffixParameters(value); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Match",
		"suffix",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Match) AsList() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"asList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Match) Resolve(context awscdk.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Match) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

