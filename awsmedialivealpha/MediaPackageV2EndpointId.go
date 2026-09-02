package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The pipeline endpoint for a MediaPackage V2 destination.
//
// Example:
//   var primary IChannel
//   var hdVideo EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_MediaPackageV2PerPipeline(&MediaPackageV2PerPipelineOutputGroupProps{
//   	Name: jsii.String("emp"),
//   	Destinations: []MediaPackageV2Destination{
//   		medialive.MediaPackageV2Destination_Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_2()),
//   		medialive.MediaPackageV2Destination_*Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_1()),
//   	},
//   	Outputs: []MediaPackageV2OutputDefinition{
//   		&MediaPackageV2OutputDefinition{
//   			Encode: hdVideo,
//   			OutputName: jsii.String("hd"),
//   		},
//   	},
//   })
//
// Experimental.
type MediaPackageV2EndpointId interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MediaPackageV2EndpointId
type jsiiProxy_MediaPackageV2EndpointId struct {
	_ byte // padding
}

func (j *jsiiProxy_MediaPackageV2EndpointId) Value() *string {
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
func MediaPackageV2EndpointId_Of(value *string) MediaPackageV2EndpointId {
	_init_.Initialize()

	if err := validateMediaPackageV2EndpointId_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MediaPackageV2EndpointId

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2EndpointId",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MediaPackageV2EndpointId_ENDPOINT_1() MediaPackageV2EndpointId {
	_init_.Initialize()
	var returns MediaPackageV2EndpointId
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2EndpointId",
		"ENDPOINT_1",
		&returns,
	)
	return returns
}

func MediaPackageV2EndpointId_ENDPOINT_2() MediaPackageV2EndpointId {
	_init_.Initialize()
	var returns MediaPackageV2EndpointId
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MediaPackageV2EndpointId",
		"ENDPOINT_2",
		&returns,
	)
	return returns
}

