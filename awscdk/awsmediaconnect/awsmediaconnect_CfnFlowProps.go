package awsmediaconnect


// Properties for defining a `CfnFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowProps := &cfnFlowProps{
//   	name: jsii.String("name"),
//   	source: &sourceProperty{
//   		decryption: &encryptionProperty{
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			algorithm: jsii.String("algorithm"),
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			deviceId: jsii.String("deviceId"),
//   			keyType: jsii.String("keyType"),
//   			region: jsii.String("region"),
//   			resourceId: jsii.String("resourceId"),
//   			secretArn: jsii.String("secretArn"),
//   			url: jsii.String("url"),
//   		},
//   		description: jsii.String("description"),
//   		entitlementArn: jsii.String("entitlementArn"),
//   		ingestIp: jsii.String("ingestIp"),
//   		ingestPort: jsii.Number(123),
//   		maxBitrate: jsii.Number(123),
//   		maxLatency: jsii.Number(123),
//   		minLatency: jsii.Number(123),
//   		name: jsii.String("name"),
//   		protocol: jsii.String("protocol"),
//   		senderControlPort: jsii.Number(123),
//   		senderIpAddress: jsii.String("senderIpAddress"),
//   		sourceArn: jsii.String("sourceArn"),
//   		sourceIngestPort: jsii.String("sourceIngestPort"),
//   		sourceListenerAddress: jsii.String("sourceListenerAddress"),
//   		sourceListenerPort: jsii.Number(123),
//   		streamId: jsii.String("streamId"),
//   		vpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		whitelistCidr: jsii.String("whitelistCidr"),
//   	},
//
//   	// the properties below are optional
//   	availabilityZone: jsii.String("availabilityZone"),
//   	sourceFailoverConfig: &failoverConfigProperty{
//   		failoverMode: jsii.String("failoverMode"),
//   		recoveryWindow: jsii.Number(123),
//   		sourcePriority: &sourcePriorityProperty{
//   			primarySource: jsii.String("primarySource"),
//   		},
//   		state: jsii.String("state"),
//   	},
//   }
//
type CfnFlowProps struct {
	// The name of the flow.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The settings for the source that you want to use for the new flow.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS Region.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The settings for source failover.
	SourceFailoverConfig interface{} `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

