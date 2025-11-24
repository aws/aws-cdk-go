package mixinsawsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSourceLocationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSourceLocationMixinProps := &CfnSourceLocationMixinProps{
//   	AccessConfiguration: &AccessConfigurationProperty{
//   		AccessType: jsii.String("accessType"),
//   		SecretsManagerAccessTokenConfiguration: &SecretsManagerAccessTokenConfigurationProperty{
//   			HeaderName: jsii.String("headerName"),
//   			SecretArn: jsii.String("secretArn"),
//   			SecretStringKey: jsii.String("secretStringKey"),
//   		},
//   	},
//   	DefaultSegmentDeliveryConfiguration: &DefaultSegmentDeliveryConfigurationProperty{
//   		BaseUrl: jsii.String("baseUrl"),
//   	},
//   	HttpConfiguration: &HttpConfigurationProperty{
//   		BaseUrl: jsii.String("baseUrl"),
//   	},
//   	SegmentDeliveryConfigurations: []interface{}{
//   		&SegmentDeliveryConfigurationProperty{
//   			BaseUrl: jsii.String("baseUrl"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html
//
type CfnSourceLocationMixinProps struct {
	// The access configuration for the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-accessconfiguration
	//
	AccessConfiguration interface{} `field:"optional" json:"accessConfiguration" yaml:"accessConfiguration"`
	// The default segment delivery configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration
	//
	DefaultSegmentDeliveryConfiguration interface{} `field:"optional" json:"defaultSegmentDeliveryConfiguration" yaml:"defaultSegmentDeliveryConfiguration"`
	// The HTTP configuration for the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-httpconfiguration
	//
	HttpConfiguration interface{} `field:"optional" json:"httpConfiguration" yaml:"httpConfiguration"`
	// The segment delivery configurations for the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-segmentdeliveryconfigurations
	//
	SegmentDeliveryConfigurations interface{} `field:"optional" json:"segmentDeliveryConfigurations" yaml:"segmentDeliveryConfigurations"`
	// The name of the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The tags assigned to the source location.
	//
	// Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

