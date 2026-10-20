package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAdConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAdConfigurationProps := &CfnAdConfigurationProps{
//   	MediaTailorPlaybackConfigurations: []interface{}{
//   		&MediaTailorPlaybackConfigurationProperty{
//   			PlaybackConfigurationArn: jsii.String("playbackConfigurationArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	PostRollConfiguration: &PostRollConfigurationProperty{
//   		DurationSeconds: jsii.Number(123),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-adconfiguration.html
//
type CfnAdConfigurationProps struct {
	// List of integration configurations with MediaTailor resources.
	//
	// The first item in the list is the default playback configuration used for the ad configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-adconfiguration.html#cfn-ivs-adconfiguration-mediatailorplaybackconfigurations
	//
	MediaTailorPlaybackConfigurations interface{} `field:"required" json:"mediaTailorPlaybackConfigurations" yaml:"mediaTailorPlaybackConfigurations"`
	// Ad configuration name.
	//
	// The value does not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-adconfiguration.html#cfn-ivs-adconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configuration for the post-roll ad break to use for this ad configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-adconfiguration.html#cfn-ivs-adconfiguration-postrollconfiguration
	//
	PostRollConfiguration interface{} `field:"optional" json:"postRollConfiguration" yaml:"postRollConfiguration"`
	// Tags attached to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-adconfiguration.html#cfn-ivs-adconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

