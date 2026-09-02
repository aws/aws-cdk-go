package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The network location of a MediaLive input — the AWS cloud, or an on-premises network for MediaLive Anywhere.
//
// Example:
//   var stack Stack
//
//
//   network := medialive.NewNetwork(stack, jsii.String("Network"), &NetworkProps{
//   	NetworkName: jsii.String("on-prem-network"),
//   	IpPools: []*string{
//   		jsii.String("192.168.1.0/24"),
//   	},
//   })
//
//   medialive.NewInput(stack, jsii.String("OnPremInput"), &InputProps{
//   	InputName: jsii.String("on-prem-rtp"),
//   	InputNetworkLocation: medialive.InputNetworkLocation_ON_PREMISES(),
//   	Input: medialive.InputConfiguration_RtpPush(&PushInputProps{
//   		Destinations: []PushInputDestination{
//   			&PushInputDestination{
//   				Network: *Network,
//   				NetworkRoutes: []NetworkRoute{
//   					&NetworkRoute{
//   						Cidr: jsii.String("10.0.0.0/24"),
//   						Gateway: jsii.String("10.0.0.1"),
//   					},
//   				},
//   				StaticIpAddress: jsii.String("192.168.1.50"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type InputNetworkLocation interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputNetworkLocation
type jsiiProxy_InputNetworkLocation struct {
	_ byte // padding
}

func (j *jsiiProxy_InputNetworkLocation) Value() *string {
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
func InputNetworkLocation_Of(value *string) InputNetworkLocation {
	_init_.Initialize()

	if err := validateInputNetworkLocation_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputNetworkLocation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputNetworkLocation",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputNetworkLocation_AWS() InputNetworkLocation {
	_init_.Initialize()
	var returns InputNetworkLocation
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputNetworkLocation",
		"AWS",
		&returns,
	)
	return returns
}

func InputNetworkLocation_ON_PREMISES() InputNetworkLocation {
	_init_.Initialize()
	var returns InputNetworkLocation
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputNetworkLocation",
		"ON_PREMISES",
		&returns,
	)
	return returns
}

