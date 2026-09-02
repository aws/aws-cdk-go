package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Sub-GOP length mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   subgopLength := medialive_alpha.SubgopLength_Of(jsii.String("value"))
//
// Experimental.
type SubgopLength interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SubgopLength
type jsiiProxy_SubgopLength struct {
	_ byte // padding
}

func (j *jsiiProxy_SubgopLength) Value() *string {
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
func SubgopLength_Of(value *string) SubgopLength {
	_init_.Initialize()

	if err := validateSubgopLength_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SubgopLength

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SubgopLength",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SubgopLength_DYNAMIC() SubgopLength {
	_init_.Initialize()
	var returns SubgopLength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SubgopLength",
		"DYNAMIC",
		&returns,
	)
	return returns
}

func SubgopLength_FIXED() SubgopLength {
	_init_.Initialize()
	var returns SubgopLength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SubgopLength",
		"FIXED",
		&returns,
	)
	return returns
}

