package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Which Dolby E program to decode from a selected audio track.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   dolbyEProgramSelection := medialive_alpha.DolbyEProgramSelection_ALL_CHANNELS()
//
// Experimental.
type DolbyEProgramSelection interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for DolbyEProgramSelection
type jsiiProxy_DolbyEProgramSelection struct {
	_ byte // padding
}

func (j *jsiiProxy_DolbyEProgramSelection) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func DolbyEProgramSelection_Of(value *string) DolbyEProgramSelection {
	_init_.Initialize()

	if err := validateDolbyEProgramSelection_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DolbyEProgramSelection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DolbyEProgramSelection_ALL_CHANNELS() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"ALL_CHANNELS",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_1() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_1",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_2() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_2",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_3() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_3",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_4() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_4",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_5() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_5",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_6() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_6",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_7() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_7",
		&returns,
	)
	return returns
}

func DolbyEProgramSelection_PROGRAM_8() DolbyEProgramSelection {
	_init_.Initialize()
	var returns DolbyEProgramSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DolbyEProgramSelection",
		"PROGRAM_8",
		&returns,
	)
	return returns
}

