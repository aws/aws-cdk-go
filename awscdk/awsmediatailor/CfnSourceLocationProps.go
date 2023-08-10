package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSourceLocation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSourceLocationProps := &CfnSourceLocationProps{
//   	HttpConfiguration: &HttpConfigurationProperty{
//   		BaseUrl: jsii.String("baseUrl"),
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//
//   	// the properties below are optional
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
//   	SegmentDeliveryConfigurations: []interface{}{
//   		&SegmentDeliveryConfigurationProperty{
//   			BaseUrl: jsii.String("baseUrl"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html
//
type CfnSourceLocationProps struct {
	// <p>The HTTP configuration for the source location.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-httpconfiguration
	//
	HttpConfiguration interface{} `field:"required" json:"httpConfiguration" yaml:"httpConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-sourcelocationname
	//
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// <p>Access configuration parameters.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-accessconfiguration
	//
	AccessConfiguration interface{} `field:"optional" json:"accessConfiguration" yaml:"accessConfiguration"`
	// <p>The optional configuration for a server that serves segments.
	//
	// Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration
	//
	DefaultSegmentDeliveryConfiguration interface{} `field:"optional" json:"defaultSegmentDeliveryConfiguration" yaml:"defaultSegmentDeliveryConfiguration"`
	// <p>A list of the segment delivery configurations associated with this resource.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-segmentdeliveryconfigurations
	//
	SegmentDeliveryConfigurations interface{} `field:"optional" json:"segmentDeliveryConfigurations" yaml:"segmentDeliveryConfigurations"`
	// The tags to assign to the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-sourcelocation.html#cfn-mediatailor-sourcelocation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

