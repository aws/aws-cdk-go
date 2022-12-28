package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFeatureGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var offlineStoreConfig interface{}
//   var onlineStoreConfig interface{}
//
//   cfnFeatureGroupProps := &cfnFeatureGroupProps{
//   	eventTimeFeatureName: jsii.String("eventTimeFeatureName"),
//   	featureDefinitions: []interface{}{
//   		&featureDefinitionProperty{
//   			featureName: jsii.String("featureName"),
//   			featureType: jsii.String("featureType"),
//   		},
//   	},
//   	featureGroupName: jsii.String("featureGroupName"),
//   	recordIdentifierFeatureName: jsii.String("recordIdentifierFeatureName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	offlineStoreConfig: offlineStoreConfig,
//   	onlineStoreConfig: onlineStoreConfig,
//   	roleArn: jsii.String("roleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFeatureGroupProps struct {
	// The name of the feature that stores the `EventTime` of a Record in a `FeatureGroup` .
	//
	// A `EventTime` is point in time when a new event occurs that corresponds to the creation or update of a `Record` in `FeatureGroup` . All `Records` in the `FeatureGroup` must have a corresponding `EventTime` .
	EventTimeFeatureName *string `field:"required" json:"eventTimeFeatureName" yaml:"eventTimeFeatureName"`
	// A list of `Feature` s. Each `Feature` must include a `FeatureName` and a `FeatureType` .
	//
	// Valid `FeatureType` s are `Integral` , `Fractional` and `String` .
	//
	// `FeatureName` s cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	//
	// You can create up to 2,500 `FeatureDefinition` s per `FeatureGroup` .
	FeatureDefinitions interface{} `field:"required" json:"featureDefinitions" yaml:"featureDefinitions"`
	// The name of the `FeatureGroup` .
	FeatureGroupName *string `field:"required" json:"featureGroupName" yaml:"featureGroupName"`
	// The name of the `Feature` whose value uniquely identifies a `Record` defined in the `FeatureGroup` `FeatureDefinitions` .
	RecordIdentifierFeatureName *string `field:"required" json:"recordIdentifierFeatureName" yaml:"recordIdentifierFeatureName"`
	// A free form description of a `FeatureGroup` .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration of an `OfflineStore` .
	OfflineStoreConfig interface{} `field:"optional" json:"offlineStoreConfig" yaml:"offlineStoreConfig"`
	// The configuration of an `OnlineStore` .
	OnlineStoreConfig interface{} `field:"optional" json:"onlineStoreConfig" yaml:"onlineStoreConfig"`
	// The Amazon Resource Name (ARN) of the IAM execution role used to create the feature group.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags used to define a `FeatureGroup` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

