package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an OS version for an EC2 Image Builder image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   oSVersion := imagebuilder_alpha.OSVersion_AMAZON_LINUX()
//
// Experimental.
type OSVersion interface {
	// The OS version name.
	// Experimental.
	OsVersion() *string
	// The Platform of the OS version.
	// Experimental.
	Platform() Platform
}

// The jsii proxy struct for OSVersion
type jsiiProxy_OSVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_OSVersion) OsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OSVersion) Platform() Platform {
	var returns Platform
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}


// Experimental.
func NewOSVersion(platform Platform, osVersion *string) OSVersion {
	_init_.Initialize()

	if err := validateNewOSVersionParameters(platform); err != nil {
		panic(err)
	}
	j := jsiiProxy_OSVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		[]interface{}{platform, osVersion},
		&j,
	)

	return &j
}

// Experimental.
func NewOSVersion_Override(o OSVersion, platform Platform, osVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		[]interface{}{platform, osVersion},
		o,
	)
}

// Constructs an OS version with a custom name.
// Experimental.
func OSVersion_Custom(platform Platform, osVersion *string) OSVersion {
	_init_.Initialize()

	if err := validateOSVersion_CustomParameters(platform); err != nil {
		panic(err)
	}
	var returns OSVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"custom",
		[]interface{}{platform, osVersion},
		&returns,
	)

	return returns
}

func OSVersion_AMAZON_LINUX() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"AMAZON_LINUX",
		&returns,
	)
	return returns
}

func OSVersion_AMAZON_LINUX_2() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"AMAZON_LINUX_2",
		&returns,
	)
	return returns
}

func OSVersion_AMAZON_LINUX_2023() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"AMAZON_LINUX_2023",
		&returns,
	)
	return returns
}

func OSVersion_LINUX() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"LINUX",
		&returns,
	)
	return returns
}

func OSVersion_MAC_OS() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"MAC_OS",
		&returns,
	)
	return returns
}

func OSVersion_MAC_OS_14() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"MAC_OS_14",
		&returns,
	)
	return returns
}

func OSVersion_MAC_OS_15() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"MAC_OS_15",
		&returns,
	)
	return returns
}

func OSVersion_REDHAT_ENTERPRISE_LINUX() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"REDHAT_ENTERPRISE_LINUX",
		&returns,
	)
	return returns
}

func OSVersion_REDHAT_ENTERPRISE_LINUX_10() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"REDHAT_ENTERPRISE_LINUX_10",
		&returns,
	)
	return returns
}

func OSVersion_REDHAT_ENTERPRISE_LINUX_8() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"REDHAT_ENTERPRISE_LINUX_8",
		&returns,
	)
	return returns
}

func OSVersion_REDHAT_ENTERPRISE_LINUX_9() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"REDHAT_ENTERPRISE_LINUX_9",
		&returns,
	)
	return returns
}

func OSVersion_SLES() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"SLES",
		&returns,
	)
	return returns
}

func OSVersion_SLES_15() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"SLES_15",
		&returns,
	)
	return returns
}

func OSVersion_UBUNTU() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"UBUNTU",
		&returns,
	)
	return returns
}

func OSVersion_UBUNTU_22_04() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"UBUNTU_22_04",
		&returns,
	)
	return returns
}

func OSVersion_UBUNTU_24_04() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"UBUNTU_24_04",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS_SERVER() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS_SERVER",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS_SERVER_2016() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS_SERVER_2016",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS_SERVER_2019() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS_SERVER_2019",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS_SERVER_2022() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS_SERVER_2022",
		&returns,
	)
	return returns
}

func OSVersion_WINDOWS_SERVER_2025() OSVersion {
	_init_.Initialize()
	var returns OSVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		"WINDOWS_SERVER_2025",
		&returns,
	)
	return returns
}

