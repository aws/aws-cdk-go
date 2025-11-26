package previewawsmediaconnectmixins


// Properties for CfnFlowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   cfnFlowMixinProps := &CfnFlowMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	FlowSize: jsii.String("flowSize"),
//   	Maintenance: &MaintenanceProperty{
//   		MaintenanceDay: jsii.String("maintenanceDay"),
//   		MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   	},
//   	MediaStreams: []interface{}{
//   		&MediaStreamProperty{
//   			Attributes: &MediaStreamAttributesProperty{
//   				Fmtp: &FmtpProperty{
//   					ChannelOrder: jsii.String("channelOrder"),
//   					Colorimetry: jsii.String("colorimetry"),
//   					ExactFramerate: jsii.String("exactFramerate"),
//   					Par: jsii.String("par"),
//   					Range: jsii.String("range"),
//   					ScanMode: jsii.String("scanMode"),
//   					Tcs: jsii.String("tcs"),
//   				},
//   				Lang: jsii.String("lang"),
//   			},
//   			ClockRate: jsii.Number(123),
//   			Description: jsii.String("description"),
//   			Fmt: jsii.Number(123),
//   			MediaStreamId: jsii.Number(123),
//   			MediaStreamName: jsii.String("mediaStreamName"),
//   			MediaStreamType: jsii.String("mediaStreamType"),
//   			VideoFormat: jsii.String("videoFormat"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NdiConfig: &NdiConfigProperty{
//   		MachineName: jsii.String("machineName"),
//   		NdiDiscoveryServers: []interface{}{
//   			&NdiDiscoveryServerConfigProperty{
//   				DiscoveryServerAddress: jsii.String("discoveryServerAddress"),
//   				DiscoveryServerPort: jsii.Number(123),
//   				VpcInterfaceAdapter: jsii.String("vpcInterfaceAdapter"),
//   			},
//   		},
//   		NdiState: jsii.String("ndiState"),
//   	},
//   	Source: &SourceProperty{
//   		Decryption: &EncryptionProperty{
//   			Algorithm: jsii.String("algorithm"),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			DeviceId: jsii.String("deviceId"),
//   			KeyType: jsii.String("keyType"),
//   			Region: jsii.String("region"),
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecretArn: jsii.String("secretArn"),
//   			Url: jsii.String("url"),
//   		},
//   		Description: jsii.String("description"),
//   		EntitlementArn: jsii.String("entitlementArn"),
//   		GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   			BridgeArn: jsii.String("bridgeArn"),
//   			VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   				VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   			},
//   		},
//   		IngestIp: jsii.String("ingestIp"),
//   		IngestPort: jsii.Number(123),
//   		MaxBitrate: jsii.Number(123),
//   		MaxLatency: jsii.Number(123),
//   		MaxSyncBuffer: jsii.Number(123),
//   		MediaStreamSourceConfigurations: []interface{}{
//   			&MediaStreamSourceConfigurationProperty{
//   				EncodingName: jsii.String("encodingName"),
//   				InputConfigurations: []interface{}{
//   					&InputConfigurationProperty{
//   						InputPort: jsii.Number(123),
//   						Interface: &InterfaceProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				MediaStreamName: jsii.String("mediaStreamName"),
//   			},
//   		},
//   		MinLatency: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		Protocol: jsii.String("protocol"),
//   		RouterIntegrationState: jsii.String("routerIntegrationState"),
//   		RouterIntegrationTransitDecryption: &FlowTransitEncryptionProperty{
//   			EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   				Automatic: automatic,
//   				SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   					RoleArn: jsii.String("roleArn"),
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   			},
//   			EncryptionKeyType: jsii.String("encryptionKeyType"),
//   		},
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
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		RecoveryWindow: jsii.Number(123),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	SourceMonitoringConfig: &SourceMonitoringConfigProperty{
//   		AudioMonitoringSettings: []interface{}{
//   			&AudioMonitoringSettingProperty{
//   				SilentAudio: &SilentAudioProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ContentQualityAnalysisState: jsii.String("contentQualityAnalysisState"),
//   		ThumbnailState: jsii.String("thumbnailState"),
//   		VideoMonitoringSettings: []interface{}{
//   			&VideoMonitoringSettingProperty{
//   				BlackFrames: &BlackFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   				FrozenFrames: &FrozenFramesProperty{
//   					State: jsii.String("state"),
//   					ThresholdSeconds: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	VpcInterfaces: []interface{}{
//   		&VpcInterfaceProperty{
//   			Name: jsii.String("name"),
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: jsii.String("networkInterfaceType"),
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html
//
type CfnFlowMixinProps struct {
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Determines the processing capacity and feature set of the flow.
	//
	// Set this optional parameter to LARGE if you want to enable NDI outputs on the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-flowsize
	//
	FlowSize *string `field:"optional" json:"flowSize" yaml:"flowSize"`
	// The maintenance settings you want to use for the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-maintenance
	//
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
	// The media streams that are associated with the flow.
	//
	// After you associate a media stream with a source, you can also associate it with outputs on the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-mediastreams
	//
	MediaStreams interface{} `field:"optional" json:"mediaStreams" yaml:"mediaStreams"`
	// The name of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the configuration settings for NDI outputs.
	//
	// Required when the flow includes NDI outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-ndiconfig
	//
	NdiConfig interface{} `field:"optional" json:"ndiConfig" yaml:"ndiConfig"`
	// The settings for the source that you want to use for the new flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The settings for source failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-sourcefailoverconfig
	//
	SourceFailoverConfig interface{} `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
	// The settings for source monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-sourcemonitoringconfig
	//
	SourceMonitoringConfig interface{} `field:"optional" json:"sourceMonitoringConfig" yaml:"sourceMonitoringConfig"`
	// The VPC Interfaces for this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-vpcinterfaces
	//
	VpcInterfaces interface{} `field:"optional" json:"vpcInterfaces" yaml:"vpcInterfaces"`
}

