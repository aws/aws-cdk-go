package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether MediaPackage sets a MediaPackage V2 audio rendition as default / auto-select in the HLS manifest.
//
// Across all renditions: at most one may be `YES`; not all may be `NO`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   mediaPackageV2HlsSetting := medialive_alpha.MediaPackageV2HlsSetting_NO()
//
// Experimental.
type MediaPackageV2HlsSetting interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MediaPackageV2HlsSetting
type jsiiProxy_MediaPackageV2HlsSetting struct {
	_ byte // padding
}

func (j *jsiiProxy_MediaPackageV2HlsSetting) Value() *string {
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
func MediaPackageV2HlsSetting_Of(value *string) MediaPackageV2HlsSetting {
	_init_.Initialize()

	if err := validateMediaPackageV2HlsSetting_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MediaPackageV2HlsSetting

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2HlsSetting",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MediaPackageV2HlsSetting_NO() MediaPackageV2HlsSetting {
	_init_.Initialize()
	var returns MediaPackageV2HlsSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2HlsSetting",
		"NO",
		&returns,
	)
	return returns
}

func MediaPackageV2HlsSetting_OMIT() MediaPackageV2HlsSetting {
	_init_.Initialize()
	var returns MediaPackageV2HlsSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2HlsSetting",
		"OMIT",
		&returns,
	)
	return returns
}

func MediaPackageV2HlsSetting_YES() MediaPackageV2HlsSetting {
	_init_.Initialize()
	var returns MediaPackageV2HlsSetting
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2HlsSetting",
		"YES",
		&returns,
	)
	return returns
}

