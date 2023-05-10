package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// Base parameters for the StackSet.
//
// Example:
//   parameters := codepipeline_actions.StackSetParameters_FromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
// Experimental.
type StackSetParameters interface {
}

// The jsii proxy struct for StackSetParameters
type jsiiProxy_StackSetParameters struct {
	_ byte // padding
}

// Experimental.
func NewStackSetParameters_Override(s StackSetParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline_actions.StackSetParameters",
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
// Experimental.
func StackSetParameters_FromArtifactPath(artifactPath awscodepipeline.ArtifactPath) StackSetParameters {
	_init_.Initialize()

	if err := validateStackSetParameters_FromArtifactPathParameters(artifactPath); err != nil {
		panic(err)
	}
	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline_actions.StackSetParameters",
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
//   parameters := codepipeline_actions.StackSetParameters_FromLiteral(map[string]*string{
//   	"BucketName": jsii.String("my-bucket"),
//   	"Asset1": jsii.String("true"),
//   })
//
// Experimental.
func StackSetParameters_FromLiteral(parameters *map[string]*string, usePreviousValues *[]*string) StackSetParameters {
	_init_.Initialize()

	if err := validateStackSetParameters_FromLiteralParameters(parameters); err != nil {
		panic(err)
	}
	var returns StackSetParameters

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline_actions.StackSetParameters",
		"fromLiteral",
		[]interface{}{parameters, usePreviousValues},
		&returns,
	)

	return returns
}

