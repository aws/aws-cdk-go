package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Timezone applied to the timestamps in a Nielsen NAES II/NW watermark.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   nielsenWatermarkTimezone := medialive_alpha.NielsenWatermarkTimezone_AMERICA_PUERTO_RICO()
//
// Experimental.
type NielsenWatermarkTimezone interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NielsenWatermarkTimezone
type jsiiProxy_NielsenWatermarkTimezone struct {
	_ byte // padding
}

func (j *jsiiProxy_NielsenWatermarkTimezone) Value() *string {
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
func NielsenWatermarkTimezone_Of(value *string) NielsenWatermarkTimezone {
	_init_.Initialize()

	if err := validateNielsenWatermarkTimezone_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NielsenWatermarkTimezone

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NielsenWatermarkTimezone_AMERICA_PUERTO_RICO() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"AMERICA_PUERTO_RICO",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_ALASKA() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_ALASKA",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_ARIZONA() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_ARIZONA",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_CENTRAL() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_CENTRAL",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_EASTERN() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_EASTERN",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_HAWAII() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_HAWAII",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_MOUNTAIN() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_MOUNTAIN",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_PACIFIC() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_PACIFIC",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_US_SAMOA() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"US_SAMOA",
		&returns,
	)
	return returns
}

func NielsenWatermarkTimezone_UTC() NielsenWatermarkTimezone {
	_init_.Initialize()
	var returns NielsenWatermarkTimezone
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenWatermarkTimezone",
		"UTC",
		&returns,
	)
	return returns
}

