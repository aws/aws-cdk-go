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
//   		GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   			BridgeArn: jsii.String("bridgeArn"),
//
//   			// the properties below are optional
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
//   				MediaStreamName: jsii.String("mediaStreamName"),
//
//   				// the properties below are optional
//   				InputConfigurations: []interface{}{
//   					&InputConfigurationProperty{
//   						InputPort: jsii.Number(123),
//   						Interface: &InterfaceProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
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
//   	Maintenance: &MaintenanceProperty{
//   		MaintenanceDay: jsii.String("maintenanceDay"),
//   		MaintenanceStartHour: jsii.String("maintenanceStartHour"),
//   	},
//   	MediaStreams: []interface{}{
//   		&MediaStreamProperty{
//   			MediaStreamId: jsii.Number(123),
//   			MediaStreamName: jsii.String("mediaStreamName"),
//   			MediaStreamType: jsii.String("mediaStreamType"),
//
//   			// the properties below are optional
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
//   			VideoFormat: jsii.String("videoFormat"),
//   		},
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
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			NetworkInterfaceIds: []*string{
//   				jsii.String("networkInterfaceIds"),
//   			},
//   			NetworkInterfaceType: jsii.String("networkInterfaceType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html
//
type CfnFlowProps struct {
	// The name of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The settings for the source that you want to use for the new flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The maintenance settings you want to use for the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-maintenance
	//
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
	// The media streams associated with the flow.
	//
	// You can associate any of these media streams with sources and outputs on the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-mediastreams
	//
	MediaStreams interface{} `field:"optional" json:"mediaStreams" yaml:"mediaStreams"`
	// The settings for source failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-sourcefailoverconfig
	//
	SourceFailoverConfig interface{} `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
	// The settings for source monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-sourcemonitoringconfig
	//
	SourceMonitoringConfig interface{} `field:"optional" json:"sourceMonitoringConfig" yaml:"sourceMonitoringConfig"`
	// The VPC interfaces that you added to this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flow.html#cfn-mediaconnect-flow-vpcinterfaces
	//
	VpcInterfaces interface{} `field:"optional" json:"vpcInterfaces" yaml:"vpcInterfaces"`
}

