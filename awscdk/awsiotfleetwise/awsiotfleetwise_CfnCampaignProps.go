package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaignProps := &cfnCampaignProps{
//   	action: jsii.String("action"),
//   	collectionScheme: &collectionSchemeProperty{
//   		conditionBasedCollectionScheme: &conditionBasedCollectionSchemeProperty{
//   			expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			conditionLanguageVersion: jsii.Number(123),
//   			minimumTriggerIntervalMs: jsii.Number(123),
//   			triggerMode: jsii.String("triggerMode"),
//   		},
//   		timeBasedCollectionScheme: &timeBasedCollectionSchemeProperty{
//   			periodMs: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	signalCatalogArn: jsii.String("signalCatalogArn"),
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	compression: jsii.String("compression"),
//   	dataExtraDimensions: []*string{
//   		jsii.String("dataExtraDimensions"),
//   	},
//   	description: jsii.String("description"),
//   	diagnosticsMode: jsii.String("diagnosticsMode"),
//   	expiryTime: jsii.String("expiryTime"),
//   	postTriggerCollectionDuration: jsii.Number(123),
//   	priority: jsii.Number(123),
//   	signalsToCollect: []interface{}{
//   		&signalInformationProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			maxSampleCount: jsii.Number(123),
//   			minimumSamplingIntervalMs: jsii.Number(123),
//   		},
//   	},
//   	spoolingMode: jsii.String("spoolingMode"),
//   	startTime: jsii.String("startTime"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCampaignProps struct {
	// `AWS::IoTFleetWise::Campaign.Action`.
	Action *string `field:"required" json:"action" yaml:"action"`
	// `AWS::IoTFleetWise::Campaign.CollectionScheme`.
	CollectionScheme interface{} `field:"required" json:"collectionScheme" yaml:"collectionScheme"`
	// `AWS::IoTFleetWise::Campaign.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::IoTFleetWise::Campaign.SignalCatalogArn`.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// `AWS::IoTFleetWise::Campaign.TargetArn`.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// `AWS::IoTFleetWise::Campaign.Compression`.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// `AWS::IoTFleetWise::Campaign.DataExtraDimensions`.
	DataExtraDimensions *[]*string `field:"optional" json:"dataExtraDimensions" yaml:"dataExtraDimensions"`
	// `AWS::IoTFleetWise::Campaign.Description`.
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

