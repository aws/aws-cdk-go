package awsmediaconnectalpha


// Attributes for importing an existing Bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var bridgeType BridgeType
//
//   bridgeAttributes := &BridgeAttributes{
//   	BridgeArn: jsii.String("bridgeArn"),
//
//   	// the properties below are optional
//   	BridgeName: jsii.String("bridgeName"),
//   	BridgeType: bridgeType,
//   	IsFailoverEnabled: jsii.Boolean(false),
//   }
//
// Experimental.
type BridgeAttributes struct {
	// The ARN of the bridge.
	// Experimental.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the bridge.
	//
	// Not encoded in the bridge ARN, so must be provided explicitly if the imported
	// bridge needs to expose `bridgeName`.
	// Default: - bridgeName is not available on the imported construct.
	//
	// Experimental.
	BridgeName *string `field:"optional" json:"bridgeName" yaml:"bridgeName"`
	// Indicates what type of bridge is imported.
	//
	// Not encoded in the bridge ARN, so must be provided explicitly if the imported
	// bridge is used with `addOutput()` or other methods that need the bridge type.
	// Default: - bridgeType is not available on the imported construct.
	//
	// Experimental.
	BridgeType BridgeType `field:"optional" json:"bridgeType" yaml:"bridgeType"`
	// Failover Configuration for Bridge.
	// Default: false.
	//
	// Experimental.
	IsFailoverEnabled *bool `field:"optional" json:"isFailoverEnabled" yaml:"isFailoverEnabled"`
}

