package previewawscloudfrontmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVpcOriginPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVpcOriginMixinProps := &CfnVpcOriginMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcOriginEndpointConfig: &VpcOriginEndpointConfigProperty{
//   		Arn: jsii.String("arn"),
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   		OriginSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html
//
type CfnVpcOriginMixinProps struct {
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html#cfn-cloudfront-vpcorigin-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig
	//
	VpcOriginEndpointConfig interface{} `field:"optional" json:"vpcOriginEndpointConfig" yaml:"vpcOriginEndpointConfig"`
}

