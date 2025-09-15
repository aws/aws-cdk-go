package cloudassemblyschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interoperable representation of a deployable cloud application.
//
// The external and interoperable contract for a Cloud Assembly is
// a directory containing a valid Cloud Assembly.
//
// Implementations should use the directory to load the Cloud Assembly from disk.
// It is recommended that implementations validate loaded manifest files using
// the provided functionality from this package.
// Within an implementation, it may be prudent to keep (parts of) the Cloud Assembly
// in memory during execution and use an implementation-specific contract.
// However when an implementation is providing an external contract,
// this interface should be used.
type ICloudAssembly interface {
	// The directory of the cloud assembly.
	//
	// This directory will be used to read the Cloud Assembly from.
	// Its contents (in particular `manifest.json`) must comply with the schema defined in this package.
	Directory() *string
}

// The jsii proxy for ICloudAssembly
type jsiiProxy_ICloudAssembly struct {
	_ byte // padding
}

func (j *jsiiProxy_ICloudAssembly) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

