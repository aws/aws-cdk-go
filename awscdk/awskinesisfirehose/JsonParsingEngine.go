package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The JSON parsing engine for MetadataExtractionProcessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonParsingEngine := awscdk.Aws_kinesisfirehose.JsonParsingEngine_Of(jsii.String("parsingEngine"))
//
type JsonParsingEngine interface {
	// The parsing engine string.
	ParsingEngine() *string
}

// The jsii proxy struct for JsonParsingEngine
type jsiiProxy_JsonParsingEngine struct {
	_ byte // padding
}

func (j *jsiiProxy_JsonParsingEngine) ParsingEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parsingEngine",
		&returns,
	)
	return returns
}


// A custom parsing engine.
func JsonParsingEngine_Of(parsingEngine *string) JsonParsingEngine {
	_init_.Initialize()

	if err := validateJsonParsingEngine_OfParameters(parsingEngine); err != nil {
		panic(err)
	}
	var returns JsonParsingEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.JsonParsingEngine",
		"of",
		[]interface{}{parsingEngine},
		&returns,
	)

	return returns
}

func JsonParsingEngine_JQ_1_6() JsonParsingEngine {
	_init_.Initialize()
	var returns JsonParsingEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.JsonParsingEngine",
		"JQ_1_6",
		&returns,
	)
	return returns
}

