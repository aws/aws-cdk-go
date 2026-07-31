package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Routing tier that determines the maximum bitrate (in Mbps) for a Router Output.
//
// Example:
//   var stack Stack
//   var mediaLiveInput CfnInput
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("MediaLiveOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("medialive-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(15)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_GLOBAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInput(&MediaLiveInputConnectionProps{
//   		MediaLiveInputArn: mediaLiveInput.attrArn,
//   		MediaLivePipelineId: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type RouterOutputTier interface {
	// The tier string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RouterOutputTier
type jsiiProxy_RouterOutputTier struct {
	_ byte // padding
}

func (j *jsiiProxy_RouterOutputTier) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom tier value.
// Experimental.
func RouterOutputTier_Of(value *string) RouterOutputTier {
	_init_.Initialize()

	if err := validateRouterOutputTier_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RouterOutputTier

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputTier",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RouterOutputTier_OUTPUT_100() RouterOutputTier {
	_init_.Initialize()
	var returns RouterOutputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputTier",
		"OUTPUT_100",
		&returns,
	)
	return returns
}

func RouterOutputTier_OUTPUT_20() RouterOutputTier {
	_init_.Initialize()
	var returns RouterOutputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputTier",
		"OUTPUT_20",
		&returns,
	)
	return returns
}

func RouterOutputTier_OUTPUT_50() RouterOutputTier {
	_init_.Initialize()
	var returns RouterOutputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputTier",
		"OUTPUT_50",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RouterOutputTier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

