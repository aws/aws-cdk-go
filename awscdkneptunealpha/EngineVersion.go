package awscdkneptunealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible Instances Types to use in Neptune cluster used for defining `DatabaseClusterProps.engineVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import neptune_alpha "github.com/aws/aws-cdk-go/awscdkneptunealpha"
//
//   engineVersion := neptune_alpha.EngineVersion_V1_0_1_0()
//
// Experimental.
type EngineVersion interface {
	// the engine version of Neptune.
	// Experimental.
	Version() *string
}

// The jsii proxy struct for EngineVersion
type jsiiProxy_EngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_EngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Constructor for specifying a custom engine version.
// Experimental.
func NewEngineVersion(version *string) EngineVersion {
	_init_.Initialize()

	if err := validateNewEngineVersionParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_EngineVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		[]interface{}{version},
		&j,
	)

	return &j
}

// Constructor for specifying a custom engine version.
// Experimental.
func NewEngineVersion_Override(e EngineVersion, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		[]interface{}{version},
		e,
	)
}

func EngineVersion_V1_0_1_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_1_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_1_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_1_2",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_2_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_2_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_2_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_2_2",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_3_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_3_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_4_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_4_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_4_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_4_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_0_5_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_0_5_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_1_0_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_1_0_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_1_1_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_1_1_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_2_0_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_2_0_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_2_0_1() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_2_0_1",
		&returns,
	)
	return returns
}

func EngineVersion_V1_2_0_2() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_2_0_2",
		&returns,
	)
	return returns
}

func EngineVersion_V1_2_1_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_2_1_0",
		&returns,
	)
	return returns
}

func EngineVersion_V1_3_0_0() EngineVersion {
	_init_.Initialize()
	var returns EngineVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-neptune-alpha.EngineVersion",
		"V1_3_0_0",
		&returns,
	)
	return returns
}

