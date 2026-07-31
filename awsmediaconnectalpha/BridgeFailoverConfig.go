package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Source failover configuration for a bridge.
//
// Bridges only support `FAILOVER` (switchover) mode — the `MERGE` mode available on
// `Flow.sourceFailoverConfig` is not allowed on bridges by the service. Use
// {@link BridgeFailoverConfig.failover} to construct the configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   bridgeFailoverConfig := mediaconnect_alpha.BridgeFailoverConfig_Failover(&BridgeFailoverOptions{
//   	PrimarySource: jsii.String("primarySource"),
//   	State: mediaconnect_alpha.State_ENABLED,
//   })
//
// Experimental.
type BridgeFailoverConfig interface {
}

// The jsii proxy struct for BridgeFailoverConfig
type jsiiProxy_BridgeFailoverConfig struct {
	_ byte // padding
}

// Configure switchover-mode failover.
//
// The bridge swaps to the backup source when
// the primary source stops receiving data.
// Experimental.
func BridgeFailoverConfig_Failover(options *BridgeFailoverOptions) BridgeFailoverConfig {
	_init_.Initialize()

	if err := validateBridgeFailoverConfig_FailoverParameters(options); err != nil {
		panic(err)
	}
	var returns BridgeFailoverConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeFailoverConfig",
		"failover",
		[]interface{}{options},
		&returns,
	)

	return returns
}

