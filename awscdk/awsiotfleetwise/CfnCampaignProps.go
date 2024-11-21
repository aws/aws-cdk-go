package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaignProps := &CfnCampaignProps{
//   	CollectionScheme: &CollectionSchemeProperty{
//   		ConditionBasedCollectionScheme: &ConditionBasedCollectionSchemeProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			ConditionLanguageVersion: jsii.Number(123),
//   			MinimumTriggerIntervalMs: jsii.Number(123),
//   			TriggerMode: jsii.String("triggerMode"),
//   		},
//   		TimeBasedCollectionScheme: &TimeBasedCollectionSchemeProperty{
//   			PeriodMs: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	SignalCatalogArn: jsii.String("signalCatalogArn"),
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	Action: jsii.String("action"),
//   	Compression: jsii.String("compression"),
//   	DataDestinationConfigs: []interface{}{
//   		&DataDestinationConfigProperty{
//   			MqttTopicConfig: &MqttTopicConfigProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				MqttTopicArn: jsii.String("mqttTopicArn"),
//   			},
//   			S3Config: &S3ConfigProperty{
//   				BucketArn: jsii.String("bucketArn"),
//
//   				// the properties below are optional
//   				DataFormat: jsii.String("dataFormat"),
//   				Prefix: jsii.String("prefix"),
//   				StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   			},
//   			TimestreamConfig: &TimestreamConfigProperty{
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				TimestreamTableArn: jsii.String("timestreamTableArn"),
//   			},
//   		},
//   	},
//   	DataExtraDimensions: []*string{
//   		jsii.String("dataExtraDimensions"),
//   	},
//   	DataPartitions: []interface{}{
//   		&DataPartitionProperty{
//   			Id: jsii.String("id"),
//   			StorageOptions: &DataPartitionStorageOptionsProperty{
//   				MaximumSize: &StorageMaximumSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				MinimumTimeToLive: &StorageMinimumTimeToLiveProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				StorageLocation: jsii.String("storageLocation"),
//   			},
//
//   			// the properties below are optional
//   			UploadOptions: &DataPartitionUploadOptionsProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				ConditionLanguageVersion: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DiagnosticsMode: jsii.String("diagnosticsMode"),
//   	ExpiryTime: jsii.String("expiryTime"),
//   	PostTriggerCollectionDuration: jsii.Number(123),
//   	Priority: jsii.Number(123),
//   	SignalsToCollect: []interface{}{
//   		&SignalInformationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DataPartitionId: jsii.String("dataPartitionId"),
//   			MaxSampleCount: jsii.Number(123),
//   			MinimumSamplingIntervalMs: jsii.Number(123),
//   		},
//   	},
//   	SignalsToFetch: []interface{}{
//   		&SignalFetchInformationProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			SignalFetchConfig: &SignalFetchConfigProperty{
//   				ConditionBased: &ConditionBasedSignalFetchConfigProperty{
//   					ConditionExpression: jsii.String("conditionExpression"),
//   					TriggerMode: jsii.String("triggerMode"),
//   				},
//   				TimeBased: &TimeBasedSignalFetchConfigProperty{
//   					ExecutionFrequencyMs: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConditionLanguageVersion: jsii.Number(123),
//   		},
//   	},
//   	SpoolingMode: jsii.String("spoolingMode"),
//   	StartTime: jsii.String("startTime"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html
//
type CfnCampaignProps struct {
	// The data collection scheme associated with the campaign.
	//
	// You can specify a scheme that collects data based on time or an event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-collectionscheme
	//
	CollectionScheme interface{} `field:"required" json:"collectionScheme" yaml:"collectionScheme"`
	// The name of a campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the signal catalog associated with the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-signalcatalogarn
	//
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// The Amazon Resource Name (ARN) of a vehicle or fleet to which the campaign is deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-targetarn
	//
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Specifies how to update a campaign. The action can be one of the following:.
	//
	// - `APPROVE` - To approve delivering a data collection scheme to vehicles.
	// - `SUSPEND` - To suspend collecting signal data. The campaign is deleted from vehicles and all vehicles in the suspended campaign will stop sending data.
	// - `RESUME` - To reactivate the `SUSPEND` campaign. The campaign is redeployed to all vehicles and the vehicles will resume sending data.
	// - `UPDATE` - To update a campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// (Optional) Whether to compress signals before transmitting data to AWS IoT FleetWise .
	//
	// If you don't want to compress the signals, use `OFF` . If it's not specified, `SNAPPY` is used.
	//
	// Default: `SNAPPY`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-compression
	//
	// Default: - "OFF".
	//
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// (Optional) The destination where the campaign sends data.
	//
	// You can choose to send data to be stored in Amazon S3 or Amazon Timestream .
	//
	// Amazon S3 optimizes the cost of data storage and provides additional mechanisms to use vehicle data, such as data lakes, centralized data storage, data processing pipelines, and analytics. AWS IoT FleetWise supports at-least-once file delivery to S3. Your vehicle data is stored on multiple AWS IoT FleetWise servers for redundancy and high availability.
	//
	// You can use Amazon Timestream to access and analyze time series data, and Timestream to query vehicle data so that you can identify trends and patterns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-datadestinationconfigs
	//
	DataDestinationConfigs interface{} `field:"optional" json:"dataDestinationConfigs" yaml:"dataDestinationConfigs"`
	// (Optional) A list of vehicle attributes to associate with a campaign.
	//
	// Enrich the data with specified vehicle attributes. For example, add `make` and `model` to the campaign, and AWS IoT FleetWise will associate the data with those attributes as dimensions in Amazon Timestream . You can then query the data against `make` and `model` .
	//
	// Default: An empty array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-dataextradimensions
	//
	DataExtraDimensions *[]*string `field:"optional" json:"dataExtraDimensions" yaml:"dataExtraDimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-datapartitions
	//
	DataPartitions interface{} `field:"optional" json:"dataPartitions" yaml:"dataPartitions"`
	// (Optional) The description of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// (Optional) Option for a vehicle to send diagnostic trouble codes to AWS IoT FleetWise .
	//
	// If you want to send diagnostic trouble codes, use `SEND_ACTIVE_DTCS` . If it's not specified, `OFF` is used.
	//
	// Default: `OFF`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-diagnosticsmode
	//
	// Default: - "OFF".
	//
	DiagnosticsMode *string `field:"optional" json:"diagnosticsMode" yaml:"diagnosticsMode"`
	// (Optional) The time the campaign expires, in seconds since epoch (January 1, 1970 at midnight UTC time).
	//
	// Vehicle data isn't collected after the campaign expires.
	//
	// Default: 253402214400 (December 31, 9999, 00:00:00 UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-expirytime
	//
	// Default: - "253402214400".
	//
	ExpiryTime *string `field:"optional" json:"expiryTime" yaml:"expiryTime"`
	// (Optional) How long (in milliseconds) to collect raw data after a triggering event initiates the collection.
	//
	// If it's not specified, `0` is used.
	//
	// Default: `0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-posttriggercollectionduration
	//
	// Default: - 0.
	//
	PostTriggerCollectionDuration *float64 `field:"optional" json:"postTriggerCollectionDuration" yaml:"postTriggerCollectionDuration"`
	// (Optional) A number indicating the priority of one campaign over another campaign for a certain vehicle or fleet.
	//
	// A campaign with the lowest value is deployed to vehicles before any other campaigns. If it's not specified, `0` is used.
	//
	// Default: `0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-priority
	//
	// Default: - 0.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// (Optional) A list of information about signals to collect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-signalstocollect
	//
	SignalsToCollect interface{} `field:"optional" json:"signalsToCollect" yaml:"signalsToCollect"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-signalstofetch
	//
	SignalsToFetch interface{} `field:"optional" json:"signalsToFetch" yaml:"signalsToFetch"`
	// (Optional) Whether to store collected data after a vehicle lost a connection with the cloud.
	//
	// After a connection is re-established, the data is automatically forwarded to AWS IoT FleetWise . If you want to store collected data when a vehicle loses connection with the cloud, use `TO_DISK` . If it's not specified, `OFF` is used.
	//
	// Default: `OFF`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-spoolingmode
	//
	// Default: - "OFF".
	//
	SpoolingMode *string `field:"optional" json:"spoolingMode" yaml:"spoolingMode"`
	// (Optional) The time, in milliseconds, to deliver a campaign after it was approved.
	//
	// If it's not specified, `0` is used.
	//
	// Default: `0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-starttime
	//
	// Default: - "0".
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// (Optional) Metadata that can be used to manage the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-campaign.html#cfn-iotfleetwise-campaign-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

