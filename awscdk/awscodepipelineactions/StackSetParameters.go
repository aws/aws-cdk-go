package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Base parameters for the StackSet.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   parameters := codepipeline_actions.StackSetParameters_FromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
type StackSetParameters interface {
}

// The jsii proxy struct for StackSetParameters
type jsiiProxy_StackSetParameters struct {
	_ byte // padding
}

func NewStackSetParameters_Override(s StackSetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		nil, // no parameters
		s,
	)
}

// Read the parameters from a JSON file from one of the pipeline's artifacts.
//
// The file needs to contain a list of `{ ParameterKey, ParameterValue, UsePreviousValue }` objects, like
// this:
//
// ```
// [
//      {
//          "ParameterKey": "BucketName",
//          "ParameterValue": "my-bucket"
//      },
//      {
//          "ParameterKey": "Asset1",
//          "ParameterValue": "true"
//      },
//      {
//          "ParameterKey": "Asset2",
//          "UsePreviousValue": true
//      }
// ]
// ```
//
// You must specify all template parameters. Parameters you don't specify will revert
// to their `Default` values as specified in the template.
//
// For of parameters you want to retain their existing values
// without specifying what those values are, set `UsePreviousValue: true`.
// Use of this feature is discouraged. CDK is for
// specifying desired-state infrastructure, and use of this feature makes the
// parameter values unmanaged.
func StackSetParameters_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath) StackSetParameters {
	_init_.Initialize()

	if err := validateStackSetParameters_FromArtifactPathParameters(artifactPath); err != nil {
		panic(err)
	}
	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		"fromArtifactPath",
		[]interface{}{artifactPath},
		&returns,
	)

	return returns
}

// A list of template parameters for your stack set.
//
// You must specify all template parameters. Parameters you don't specify will revert
// to their `Default` values as specified in the template.
//
// Specify the names of parameters you want to retain their existing values,
// without specifying what those values are, in an array in the second
// argument to this function. Use of this feature is discouraged. CDK is for
// specifying desired-state infrastructure, and use of this feature makes the
// parameter values unmanaged.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   parameters := codepipeline_actions.StackSetParameters_FromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
func StackSetParameters_FromLiteral(parameters *map[string]*string, usePreviousValues *[]*string) StackSetParameters {
	_init_.Initialize()

	if err := validateStackSetParameters_FromLiteralParameters(parameters); err != nil {
		panic(err)
	}
	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline_actions.StackSetParameters",
		"fromLiteral",
		[]interface{}{parameters, usePreviousValues},
		&returns,
	)

	return returns
}

