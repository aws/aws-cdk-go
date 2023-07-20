package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   cfnFeatureGroupProps := &CfnFeatureGroupProps{
//   	EventTimeFeatureName: jsii.String("eventTimeFeatureName"),
//   	FeatureDefinitions: []interface{}{
//   		&FeatureDefinitionProperty{
//   			FeatureName: jsii.String("featureName"),
//   			FeatureType: jsii.String("featureType"),
//   		},
//   	},
//   	FeatureGroupName: jsii.String("featureGroupName"),
//   	RecordIdentifierFeatureName: jsii.String("recordIdentifierFeatureName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	OfflineStoreConfig: offlineStoreConfig,
//   	OnlineStoreConfig: onlineStoreConfig,
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html
//
type CfnFeatureGroupProps struct {
	// The name of the feature that stores the `EventTime` of a Record in a `FeatureGroup` .
	//
	// A `EventTime` is point in time when a new event occurs that corresponds to the creation or update of a `Record` in `FeatureGroup` . All `Records` in the `FeatureGroup` must have a corresponding `EventTime` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-eventtimefeaturename
	//
	EventTimeFeatureName *string `field:"required" json:"eventTimeFeatureName" yaml:"eventTimeFeatureName"`
	// A list of `Feature` s. Each `Feature` must include a `FeatureName` and a `FeatureType` .
	//
	// Valid `FeatureType` s are `Integral` , `Fractional` and `String` .
	//
	// `FeatureName` s cannot be any of the following: `is_deleted` , `write_time` , `api_invocation_time` .
	//
	// You can create up to 2,500 `FeatureDefinition` s per `FeatureGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-featuredefinitions
	//
	FeatureDefinitions interface{} `field:"required" json:"featureDefinitions" yaml:"featureDefinitions"`
	// The name of the `FeatureGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-featuregroupname
	//
	FeatureGroupName *string `field:"required" json:"featureGroupName" yaml:"featureGroupName"`
	// The name of the `Feature` whose value uniquely identifies a `Record` defined in the `FeatureGroup` `FeatureDefinitions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-recordidentifierfeaturename
	//
	RecordIdentifierFeatureName *string `field:"required" json:"recordIdentifierFeatureName" yaml:"recordIdentifierFeatureName"`
	// A free form description of a `FeatureGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration of an `OfflineStore` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-offlinestoreconfig
	//
	OfflineStoreConfig interface{} `field:"optional" json:"offlineStoreConfig" yaml:"offlineStoreConfig"`
	// The configuration of an `OnlineStore` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-onlinestoreconfig
	//
	OnlineStoreConfig interface{} `field:"optional" json:"onlineStoreConfig" yaml:"onlineStoreConfig"`
	// The Amazon Resource Name (ARN) of the IAM execution role used to create the feature group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags used to define a `FeatureGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-featuregroup.html#cfn-sagemaker-featuregroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

