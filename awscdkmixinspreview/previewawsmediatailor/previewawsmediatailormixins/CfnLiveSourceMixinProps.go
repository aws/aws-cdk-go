package previewawsmediatailormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLiveSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLiveSourceMixinProps := &CfnLiveSourceMixinProps{
//   	HttpPackageConfigurations: []interface{}{
//   		&HttpPackageConfigurationProperty{
//   			Path: jsii.String("path"),
//   			SourceGroup: jsii.String("sourceGroup"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-livesource.html
//
type CfnLiveSourceMixinProps struct {
	// The HTTP package configurations for the live source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-livesource.html#cfn-mediatailor-livesource-httppackageconfigurations
	//
	HttpPackageConfigurations interface{} `field:"optional" json:"httpPackageConfigurations" yaml:"httpPackageConfigurations"`
	// The name that's used to refer to a live source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-livesource.html#cfn-mediatailor-livesource-livesourcename
	//
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// The name of the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-livesource.html#cfn-mediatailor-livesource-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The tags assigned to the live source.
	//
	// Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-livesource.html#cfn-mediatailor-livesource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

