package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CMAF Ingest SCTE-35 type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   scte35Type := medialive_alpha.Scte35Type_Of(jsii.String("value"))
//
// Experimental.
type Scte35Type interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Scte35Type
type jsiiProxy_Scte35Type struct {
	_ byte // padding
}

func (j *jsiiProxy_Scte35Type) Value() *string {
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
func Scte35Type_Of(value *string) Scte35Type {
	_init_.Initialize()

	if err := validateScte35Type_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Scte35Type

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Scte35Type",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Scte35Type_NONE() Scte35Type {
	_init_.Initialize()
	var returns Scte35Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35Type",
		"NONE",
		&returns,
	)
	return returns
}

func Scte35Type_SCTE_35_WITHOUT_SEGMENTATION() Scte35Type {
	_init_.Initialize()
	var returns Scte35Type
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35Type",
		"SCTE_35_WITHOUT_SEGMENTATION",
		&returns,
	)
	return returns
}

