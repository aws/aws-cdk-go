package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   var myRole role
//
//   cr.NewAwsCustomResource(this, jsii.String("Customized"), &AwsCustomResourceProps{
//   	Role: myRole,
//   	 // must be assumable by the `lambda.amazonaws.com` service principal
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	 // defaults to 2 minutes
//   	LogGroup: logs.NewLogGroup(this, jsii.String("AwsCustomResourceLogs"), &LogGroupProps{
//   		Retention: logs.RetentionDays_ONE_DAY,
//   	}),
//   	FunctionName: jsii.String("my-custom-name"),
//   	 // defaults to a CloudFormation generated name
//   	RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	 // defaults to `RemovalPolicy.DESTROY`
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type Duration interface {
	// Returns stringified number of duration.
	FormatTokenToNumber() *string
	// Checks if duration is a token or a resolvable object.
	IsUnresolved() *bool
	// Substract two Durations together.
	Minus(rhs Duration) Duration
	// Add two Durations together.
	Plus(rhs Duration) Duration
	// Return the total number of days in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Days.
	ToDays(opts *TimeConversionOptions) *float64
	// Return the total number of hours in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Hours.
	ToHours(opts *TimeConversionOptions) *float64
	// Turn this duration into a human-readable string.
	ToHumanString() *string
	// Return an ISO 8601 representation of this period.
	//
	// Returns: a string starting with 'P' describing the period.
	// See: https://www.iso.org/standard/70907.html
	//
	ToIsoString() *string
	// Return the total number of milliseconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Milliseconds.
	ToMilliseconds(opts *TimeConversionOptions) *float64
	// Return the total number of minutes in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Minutes.
	ToMinutes(opts *TimeConversionOptions) *float64
	// Return the total number of seconds in this Duration.
	//
	// Returns: the value of this `Duration` expressed in Seconds.
	ToSeconds(opts *TimeConversionOptions) *float64
	// Returns a string representation of this `Duration`.
	//
	// This is is never the right function to use when you want to use the `Duration`
	// object in a template. Use `toSeconds()`, `toMinutes()`, `toDays()`, etc. instead.
	ToString() *string
	// Returns unit of the duration.
	UnitLabel() *string
}

// The jsii proxy struct for Duration
type jsiiProxy_Duration struct {
	_ byte // padding
}

// Create a Duration representing an amount of days.
//
// Returns: a new `Duration` representing `amount` Days.
func Duration_Days(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_DaysParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"days",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of hours.
//
// Returns: a new `Duration` representing `amount` Hours.
func Duration_Hours(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_HoursParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"hours",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of milliseconds.
//
// Returns: a new `Duration` representing `amount` ms.
func Duration_Millis(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_MillisParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of minutes.
//
// Returns: a new `Duration` representing `amount` Minutes.
func Duration_Minutes(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_MinutesParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"minutes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Parse a period formatted according to the ISO 8601 standard.
//
// Returns: the parsed `Duration`.
// See: https://www.iso.org/standard/70907.html
//
func Duration_Parse(duration *string) Duration {
	_init_.Initialize()

	if err := validateDuration_ParseParameters(duration); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
		"parse",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of seconds.
//
// Returns: a new `Duration` representing `amount` Seconds.
func Duration_Seconds(amount *float64) Duration {
	_init_.Initialize()

	if err := validateDuration_SecondsParameters(amount); err != nil {
		panic(err)
	}
	var returns Duration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Duration",
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

