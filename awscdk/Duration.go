package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a length of time.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative, or as an unresolved number token.
//
// When the amount is passed as a token, unit conversion is not possible.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
//   	EventPattern: &EventPattern{
//   		Source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.AddTarget(targets.NewLambdaFunction(fn, &LambdaFunctionProps{
//   	DeadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	MaxEventAge: cdk.Duration_Hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	RetryAttempts: jsii.Number(2),
//   }))
//
// Experimental.
type Duration interface {
	// Returns stringified number of duration.
	// Experimental.
	FormatTokenToNumber() *string
	// Checks if duration is a token or a resolvable object.
	// Experimental.
	IsUnresolved() *bool
	// Substract two Durations together.
	// Experimental.
	Minus(rhs Duration) Duration
	// Add two Durations together.
	// Experimental.
	Plus(rhs Duration) Duration
	// Return the total number of days in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Days.
	// Experimental.
	ToDays(opts *TimeConversionOptions) *float64
	// Return the total number of hours in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Hours.
	// Experimental.
	ToHours(opts *TimeConversionOptions) *float64
	// Turn this duration into a human-readable string.
	// Experimental.
	ToHumanString() *string
	// Return an ISO 8601 representation of this period.
	//
	// Returns: a string starting with 'P' describing the period.
	// See: https://www.iso.org/fr/standard/70907.html
	//
	// Experimental.
	ToIsoString() *string
	// Return an ISO 8601 representation of this period.
	//
	// Returns: a string starting with 'P' describing the period.
	// See: https://www.iso.org/fr/standard/70907.html
	//
	// Deprecated: Use `toIsoString()` instead.
	ToISOString() *string
	// Return the total number of milliseconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Milliseconds.
	// Experimental.
	ToMilliseconds(opts *TimeConversionOptions) *float64
	// Return the total number of minutes in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Minutes.
	// Experimental.
	ToMinutes(opts *TimeConversionOptions) *float64
	// Return the total number of seconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Seconds.
	// Experimental.
	ToSeconds(opts *TimeConversionOptions) *float64
	// Returns a string representation of this `Duration`.
	//
	// This is is never the right function to use when you want to use the `Duration`
	// object in a template. Use `toSeconds()`, `toMinutes()`, `toDays()`, etc. instead.
	// Experimental.
	ToString() *string
	// Returns unit of the duration.
	// Experimental.
	UnitLabel() *string
}

// The jsii proxy struct for Duration
type jsiiProxy_Duration struct {
	_ byte // padding
}

// Create a Duration representing an amount of days.
//
// Returns: a new `Duration` representing `amount` Days.
// Experimental.
func Duration_Days(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_DaysParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"days",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of hours.
//
// Returns: a new `Duration` representing `amount` Hours.
// Experimental.
func Duration_Hours(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_HoursParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"hours",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of milliseconds.
//
// Returns: a new `Duration` representing `amount` ms.
// Experimental.
func Duration_Millis(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_MillisParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of minutes.
//
// Returns: a new `Duration` representing `amount` Minutes.
// Experimental.
func Duration_Minutes(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_MinutesParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"minutes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Parse a period formatted according to the ISO 8601 standard.
//
// Returns: the parsed `Duration`.
// See: https://www.iso.org/fr/standard/70907.html
//
// Experimental.
func Duration_Parse(duration *string) Duration {
	_init_.Initialize()

	if err := validateDuration_ParseParameters(duration); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"parse",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of seconds.
//
// Returns: a new `Duration` representing `amount` Seconds.
// Experimental.
func Duration_Seconds(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_SecondsParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"monocdk.Duration",
		"seconds",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) FormatTokenToNumber() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"formatTokenToNumber",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) IsUnresolved() *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"isUnresolved",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) Minus(rhs Duration) Duration {
	if err := d.validateMinusParameters(rhs); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.Invoke(
		d,
		"minus",
		[]interface{}{rhs},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) Plus(rhs Duration) Duration {
	if err := d.validatePlusParameters(rhs); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.Invoke(
		d,
		"plus",
		[]interface{}{rhs},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToDays(opts *TimeConversionOptions) *float64 {
	if err := d.validateToDaysParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"toDays",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToHours(opts *TimeConversionOptions) *float64 {
	if err := d.validateToHoursParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"toHours",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToHumanString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toHumanString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToIsoString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toIsoString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToISOString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toISOString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToMilliseconds(opts *TimeConversionOptions) *float64 {
	if err := d.validateToMillisecondsParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMilliseconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToMinutes(opts *TimeConversionOptions) *float64 {
	if err := d.validateToMinutesParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMinutes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToSeconds(opts *TimeConversionOptions) *float64 {
	if err := d.validateToSecondsParameters(opts); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"toSeconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Duration) UnitLabel() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"unitLabel",
		nil, // no parameters
		&returns,
	)

	return returns
}

