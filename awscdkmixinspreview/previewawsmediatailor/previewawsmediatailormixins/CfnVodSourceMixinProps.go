package previewawsmediatailormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVodSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVodSourceMixinProps := &CfnVodSourceMixinProps{
//   	HttpPackageConfigurations: []interface{}{
//   		&HttpPackageConfigurationProperty{
//   			Path: jsii.String("path"),
//   			SourceGroup: jsii.String("sourceGroup"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html
//
type CfnVodSourceMixinProps struct {
	// The HTTP package configurations for the VOD source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-httppackageconfigurations
	//
	HttpPackageConfigurations interface{} `field:"optional" json:"httpPackageConfigurations" yaml:"httpPackageConfigurations"`
	// The name of the source location that the VOD source is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The tags assigned to the VOD source.
	//
	// Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the VOD source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-vodsource.html#cfn-mediatailor-vodsource-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

