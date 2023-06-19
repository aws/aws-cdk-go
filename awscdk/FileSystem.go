package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// File system utilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystem := monocdk.NewFileSystem()
//
// Experimental.
type FileSystem interface {
}

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	_ byte // padding
}

// Experimental.
func NewFileSystem() FileSystem {
	_init_.Initialize()

	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"monocdk.FileSystem",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFileSystem_Override(f FileSystem) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.FileSystem",
		nil, // no parameters
		f,
	)
}

// Copies an entire directory structure.
// Experimental.
func FileSystem_CopyDirectory(srcDir *string, destDir *string, options *CopyOptions, rootDir *string) {
	_init_.Initialize()

	if err := validateFileSystem_CopyDirectoryParameters(srcDir, destDir, options); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.FileSystem",
		"copyDirectory",
		[]interface{}{srcDir, destDir, options, rootDir},
	)
}

// Produces fingerprint based on the contents of a single file or an entire directory tree.
//
// The fingerprint will also include:
// 1. An extra string if defined in `options.extra`.
// 2. The set of exclude patterns, if defined in `options.exclude`
// 3. The symlink follow mode value.
// Experimental.
func FileSystem_Fingerprint(fileOrDirectory *string, options *FingerprintOptions) *string {
	_init_.Initialize()

	if err := validateFileSystem_FingerprintParameters(fileOrDirectory, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.FileSystem",
		"fingerprint",
		[]interface{}{fileOrDirectory, options},
		&returns,
	)

	return returns
}

// Checks whether a directory is empty.
// Experimental.
func FileSystem_IsEmpty(dir *string) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsEmptyParameters(dir); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.FileSystem",
		"isEmpty",
		[]interface{}{dir},
		&returns,
	)

	return returns
}

// Creates a unique temporary directory in the **system temp directory**.
// Experimental.
func FileSystem_Mkdtemp(prefix *string) *string {
	_init_.Initialize()

	if err := validateFileSystem_MkdtempParameters(prefix); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.FileSystem",
		"mkdtemp",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func FileSystem_Tmpdir() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.FileSystem",
		"tmpdir",
		&returns,
	)
	return returns
}

