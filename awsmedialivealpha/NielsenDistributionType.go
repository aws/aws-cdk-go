package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Nielsen watermark distribution type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   nielsenDistributionType := medialive_alpha.NielsenDistributionType_Of(jsii.String("value"))
//
// Experimental.
type NielsenDistributionType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NielsenDistributionType
type jsiiProxy_NielsenDistributionType struct {
	_ byte // padding
}

func (j *jsiiProxy_NielsenDistributionType) Value() *string {
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
func NielsenDistributionType_Of(value *string) NielsenDistributionType {
	_init_.Initialize()

	if err := validateNielsenDistributionType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NielsenDistributionType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NielsenDistributionType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NielsenDistributionType_FINAL_DISTRIBUTOR() NielsenDistributionType {
	_init_.Initialize()
	var returns NielsenDistributionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenDistributionType",
		"FINAL_DISTRIBUTOR",
		&returns,
	)
	return returns
}

func NielsenDistributionType_PROGRAM_CONTENT() NielsenDistributionType {
	_init_.Initialize()
	var returns NielsenDistributionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenDistributionType",
		"PROGRAM_CONTENT",
		&returns,
	)
	return returns
}

