package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP cache full behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpCacheFullBehavior := medialive_alpha.RtmpCacheFullBehavior_Of(jsii.String("value"))
//
// Experimental.
type RtmpCacheFullBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpCacheFullBehavior
type jsiiProxy_RtmpCacheFullBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpCacheFullBehavior) Value() *string {
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
func RtmpCacheFullBehavior_Of(value *string) RtmpCacheFullBehavior {
	_init_.Initialize()

	if err := validateRtmpCacheFullBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpCacheFullBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpCacheFullBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpCacheFullBehavior_DISCONNECT_IMMEDIATELY() RtmpCacheFullBehavior {
	_init_.Initialize()
	var returns RtmpCacheFullBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCacheFullBehavior",
		"DISCONNECT_IMMEDIATELY",
		&returns,
	)
	return returns
}

func RtmpCacheFullBehavior_WAIT_FOR_SERVER() RtmpCacheFullBehavior {
	_init_.Initialize()
	var returns RtmpCacheFullBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCacheFullBehavior",
		"WAIT_FOR_SERVER",
		&returns,
	)
	return returns
}

