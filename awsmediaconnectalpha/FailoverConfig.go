package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Source failover configuration for a flow.
//
// MediaConnect supports two failover modes:
// - {@link FailoverConfig.merge} combines two binary-identical sources into a single
//   stream, recovering lost packets from the other source (SMPTE 2022-7).
// - {@link FailoverConfig.failover} switches between a primary and backup source when
//   the active source stops receiving data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   failoverConfig := mediaconnect_alpha.FailoverConfig_Failover(&FailoverFailoverOptions{
//   	PrimarySource: jsii.String("primarySource"),
//   	State: mediaconnect_alpha.State_ENABLED,
//   })
//
// Experimental.
type FailoverConfig interface {
}

// The jsii proxy struct for FailoverConfig
type jsiiProxy_FailoverConfig struct {
	_ byte // padding
}

// Configure switchover-mode failover.
//
// The flow swaps to the backup source when the
// primary source stops receiving data.
// Experimental.
func FailoverConfig_Failover(options *FailoverFailoverOptions) FailoverConfig {
	_init_.Initialize()

	if err := validateFailoverConfig_FailoverParameters(options); err != nil {
		panic(err)
	}
	var returns FailoverConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverConfig",
		"failover",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Configure merge-mode failover.
//
// Requires two binary-identical sources.
// Experimental.
func FailoverConfig_Merge(options *MergeFailoverOptions) FailoverConfig {
	_init_.Initialize()

	if err := validateFailoverConfig_MergeParameters(options); err != nil {
		panic(err)
	}
	var returns FailoverConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.FailoverConfig",
		"merge",
		[]interface{}{options},
		&returns,
	)

	return returns
}

