package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 treeblock size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265TreeblockSize := medialive_alpha.H265TreeblockSize_Of(jsii.String("value"))
//
// Experimental.
type H265TreeblockSize interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265TreeblockSize
type jsiiProxy_H265TreeblockSize struct {
	_ byte // padding
}

func (j *jsiiProxy_H265TreeblockSize) Value() *string {
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
func H265TreeblockSize_Of(value *string) H265TreeblockSize {
	_init_.Initialize()

	if err := validateH265TreeblockSize_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265TreeblockSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265TreeblockSize",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265TreeblockSize_AUTO() H265TreeblockSize {
	_init_.Initialize()
	var returns H265TreeblockSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265TreeblockSize",
		"AUTO",
		&returns,
	)
	return returns
}

func H265TreeblockSize_TREE_SIZE_32X32() H265TreeblockSize {
	_init_.Initialize()
	var returns H265TreeblockSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265TreeblockSize",
		"TREE_SIZE_32X32",
		&returns,
	)
	return returns
}

