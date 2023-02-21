package regioninfo

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A database of regional information.
//
// Example:
//   type myFact struct {
//   	region
//   	name
//   	value
//   }
//
//   regionInfo.fact.register(NewMyFact())
//
// Experimental.
type Fact interface {
}

// The jsii proxy struct for Fact
type jsiiProxy_Fact struct {
	_ byte // padding
}

// Retrieves a fact from this Fact database.
//
// Returns: the fact value if it is known, and `undefined` otherwise.
// Experimental.
func Fact_Find(region *string, name *string) *string {
	_init_.Initialize()

	if err := validateFact_FindParameters(region, name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.region_info.Fact",
		"find",
		[]interface{}{region, name},
		&returns,
	)

	return returns
}

// Registers a new fact in this Fact database.
// Experimental.
func Fact_Register(fact IFact, allowReplacing *bool) {
	_init_.Initialize()

	if err := validateFact_RegisterParameters(fact); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.region_info.Fact",
		"register",
		[]interface{}{fact, allowReplacing},
	)
}

// Retrieve a fact from the Fact database.
//
// (retrieval will fail if the specified region or
// fact name does not exist.)
// Experimental.
func Fact_RequireFact(region *string, name *string) *string {
	_init_.Initialize()

	if err := validateFact_RequireFactParameters(region, name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.region_info.Fact",
		"requireFact",
		[]interface{}{region, name},
		&returns,
	)

	return returns
}

// Removes a fact from the database.
// Experimental.
func Fact_Unregister(region *string, name *string, value *string) {
	_init_.Initialize()

	if err := validateFact_UnregisterParameters(region, name); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.region_info.Fact",
		"unregister",
		[]interface{}{region, name, value},
	)
}

func Fact_Regions() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.region_info.Fact",
		"regions",
		&returns,
	)
	return returns
}

