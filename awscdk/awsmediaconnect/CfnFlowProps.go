package awsmediaconnect


// Properties for defining a `CfnFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowProps := &CfnFlowProps{
//   	Name: jsii.String("name"),
//   	Source: &SourceProperty{
//   		Decryption: &EncryptionProperty{
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			Algorithm: jsii.String("algorithm"),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			DeviceId: jsii.String("deviceId"),
//   			KeyType: jsii.String("keyType"),
//   			Region: jsii.String("region"),
//   			ResourceId: jsii.String("resourceId"),
//   			SecretArn: jsii.String("secretArn"),
//   			Url: jsii.String("url"),
//   		},
//   		Description: jsii.String("description"),
//   		EntitlementArn: jsii.String("entitlementArn"),
//   		IngestIp: jsii.String("ingestIp"),
//   		IngestPort: jsii.Number(123),
//   		MaxBitrate: jsii.Number(123),
//   		MaxLatency: jsii.Number(123),
//   		MinLatency: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		Protocol: jsii.String("protocol"),
//   		SenderControlPort: jsii.Number(123),
//   		SenderIpAddress: jsii.String("senderIpAddress"),
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceIngestPort: jsii.String("sourceIngestPort"),
//   		SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   		SourceListenerPort: jsii.Number(123),
//   		StreamId: jsii.String("streamId"),
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		WhitelistCidr: jsii.String("whitelistCidr"),
//   	},
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		RecoveryWindow: jsii.Number(123),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
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

