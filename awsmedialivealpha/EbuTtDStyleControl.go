package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether EBU-TT-D includes source style information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ebuTtDStyleControl := medialive_alpha.EbuTtDStyleControl_Of(jsii.String("value"))
//
// Experimental.
type EbuTtDStyleControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for EbuTtDStyleControl
type jsiiProxy_EbuTtDStyleControl struct {
	_ byte // padding
}

func (j *jsiiProxy_EbuTtDStyleControl) Value() *string {
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
func EbuTtDStyleControl_Of(value *string) EbuTtDStyleControl {
	_init_.Initialize()

	if err := validateEbuTtDStyleControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns EbuTtDStyleControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.EbuTtDStyleControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func EbuTtDStyleControl_EXCLUDE() EbuTtDStyleControl {
	_init_.Initialize()
	var returns EbuTtDStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.EbuTtDStyleControl",
		"EXCLUDE",
		&returns,
	)
	return returns
}

func EbuTtDStyleControl_INCLUDE() EbuTtDStyleControl {
	_init_.Initialize()
	var returns EbuTtDStyleControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.EbuTtDStyleControl",
		"INCLUDE",
		&returns,
	)
	return returns
}

