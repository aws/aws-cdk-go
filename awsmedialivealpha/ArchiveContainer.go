package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The container (transport stream) for an Archive output.
//
// Use the static factory methods to
// select between an MPEG-TS (M2TS) container or a raw container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var m2tsSettings M2tsSettings
//
//   archiveContainer := medialive_alpha.ArchiveContainer_M2ts(m2tsSettings)
//
// Experimental.
type ArchiveContainer interface {
}

// The jsii proxy struct for ArchiveContainer
type jsiiProxy_ArchiveContainer struct {
	_ byte // padding
}

// Experimental.
func NewArchiveContainer_Override(a ArchiveContainer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.ArchiveContainer",
		nil, // no parameters
		a,
	)
}

// An MPEG-TS (M2TS) container, optionally configured via `M2tsSettings`.
// Experimental.
func ArchiveContainer_M2ts(settings M2tsSettings) ArchiveContainer {
	_init_.Initialize()

	var returns ArchiveContainer

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ArchiveContainer",
		"m2ts",
		[]interface{}{settings},
		&returns,
	)

	return returns
}

// A raw container (no transport-stream wrapping).
// Experimental.
func ArchiveContainer_Raw() ArchiveContainer {
	_init_.Initialize()

	var returns ArchiveContainer

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ArchiveContainer",
		"raw",
		nil, // no parameters
		&returns,
	)

	return returns
}

