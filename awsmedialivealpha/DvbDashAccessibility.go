package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// DVB DASH accessibility signaling for an audio output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   dvbDashAccessibility := medialive_alpha.DvbDashAccessibility_CLEAN_FEED()
//
// Experimental.
type DvbDashAccessibility interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for DvbDashAccessibility
type jsiiProxy_DvbDashAccessibility struct {
	_ byte // padding
}

func (j *jsiiProxy_DvbDashAccessibility) Value() *string {
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
func DvbDashAccessibility_Of(value *string) DvbDashAccessibility {
	_init_.Initialize()

	if err := validateDvbDashAccessibility_OfParameters(value); err != nil {
		panic(err)
	}
	var returns DvbDashAccessibility

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DvbDashAccessibility_CLEAN_FEED() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"CLEAN_FEED",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_DIRECTORS_COMMENTARY() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"DIRECTORS_COMMENTARY",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_EDUCATIONAL_NOTES() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"EDUCATIONAL_NOTES",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_HARD_OF_HEARING() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"HARD_OF_HEARING",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_MAIN_PROGRAM() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"MAIN_PROGRAM",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_SUPPLEMENTAL_COMMENTARY() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"SUPPLEMENTAL_COMMENTARY",
		&returns,
	)
	return returns
}

func DvbDashAccessibility_VISUALLY_IMPAIRED() DvbDashAccessibility {
	_init_.Initialize()
	var returns DvbDashAccessibility
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.DvbDashAccessibility",
		"VISUALLY_IMPAIRED",
		&returns,
	)
	return returns
}

