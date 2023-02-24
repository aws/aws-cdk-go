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
//   	Action: jsii.String("action"),
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
//   	Compression: jsii.String("compression"),
//   	DataExtraDimensions: []*string{
//   		jsii.String("dataExtraDimensions"),
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
//   			MaxSampleCount: jsii.Number(123),
//   			MinimumSamplingIntervalMs: jsii.Number(123),
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
type CfnCampaignProps struct {
	// `AWS::IoTFleetWise::Campaign.Action`.
	Action *string `field:"required" json:"action" yaml:"action"`
	// `AWS::IoTFleetWise::Campaign.CollectionScheme`.
	CollectionScheme interface{} `field:"required" json:"collectionScheme" yaml:"collectionScheme"`
	// The name of a campaign.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the signal catalog associated with the campaign.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// The ARN of a vehicle or fleet to which the campaign is deployed.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// `AWS::IoTFleetWise::Campaign.Compression`.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// `AWS::IoTFleetWise::Campaign.DataExtraDimensions`.
	DataExtraDimensions *[]*string `field:"optional" json:"dataExtraDimensions" yaml:"dataExtraDimensions"`
	// The description of the campaign.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::IoTFleetWise::Campaign.DiagnosticsMode`.
	DiagnosticsMode *string `field:"optional" json:"diagnosticsMode" yaml:"diagnosticsMode"`
	// `AWS::IoTFleetWise::Campaign.ExpiryTime`.
	ExpiryTime *string `field:"optional" json:"expiryTime" yaml:"expiryTime"`
	// `AWS::IoTFleetWise::Campaign.PostTriggerCollectionDuration`.
	PostTriggerCollectionDuration *float64 `field:"optional" json:"postTriggerCollectionDuration" yaml:"postTriggerCollectionDuration"`
	// `AWS::IoTFleetWise::Campaign.Priority`.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// `AWS::IoTFleetWise::Campaign.SignalsToCollect`.
	SignalsToCollect interface{} `field:"optional" json:"signalsToCollect" yaml:"signalsToCollect"`
	// `AWS::IoTFleetWise::Campaign.SpoolingMode`.
	SpoolingMode *string `field:"optional" json:"spoolingMode" yaml:"spoolingMode"`
	// `AWS::IoTFleetWise::Campaign.StartTime`.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// `AWS::IoTFleetWise::Campaign.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

