package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Object respresenting an Aspect application.
//
// Stores the Aspect, the pointer to the construct it was applied
// to, and the priority value of that Aspect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var aspect IAspect
//   var construct Construct
//
//   aspectApplication := cdk.NewAspectApplication(construct, aspect, jsii.Number(123))
//
type AspectApplication interface {
	// The Aspect that was applied.
	Aspect() IAspect
	// The construct that the Aspect was applied to.
	Construct() constructs.IConstruct
	// Gets the priority value.
	//
	// Sets the priority value.
	Priority() *float64
	SetPriority(val *float64)
}

// The jsii proxy struct for AspectApplication
type jsiiProxy_AspectApplication struct {
	_ byte // padding
}

func (j *jsiiProxy_AspectApplication) Aspect() IAspect {
	var returns IAspect
	_jsii_.Get(
		j,
		"aspect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AspectApplication) Construct() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"construct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AspectApplication) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}


// Initializes AspectApplication object.
func NewAspectApplication(construct constructs.IConstruct, aspect IAspect, priority *float64) AspectApplication {
	_init_.Initialize()

	if err := validateNewAspectApplicationParameters(construct, aspect, priority); err != nil {
		panic(err)
	}
	j := jsiiProxy_AspectApplication{}

	_jsii_.Create(
		"aws-cdk-lib.AspectApplication",
		[]interface{}{construct, aspect, priority},
		&j,
	)

	return &j
}

// Initializes AspectApplication object.
func NewAspectApplication_Override(a AspectApplication, construct constructs.IConstruct, aspect IAspect, priority *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.AspectApplication",
		[]interface{}{construct, aspect, priority},
		a,
	)
}

func (j *jsiiProxy_AspectApplication)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

