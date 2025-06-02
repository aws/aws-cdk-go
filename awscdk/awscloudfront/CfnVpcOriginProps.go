package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVpcOrigin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcOriginProps := &CfnVpcOriginProps{
//   	VpcOriginEndpointConfig: &VpcOriginEndpointConfigProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HttpPort: jsii.Number(123),
//   		HttpsPort: jsii.Number(123),
//   		OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   		OriginSslProtocols: []*string{
//   			jsii.String("originSslProtocols"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html
//
type CfnVpcOriginProps struct {
	// The VPC origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html#cfn-cloudfront-vpcorigin-vpcoriginendpointconfig
	//
	VpcOriginEndpointConfig interface{} `field:"required" json:"vpcOriginEndpointConfig" yaml:"vpcOriginEndpointConfig"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-vpcorigin.html#cfn-cloudfront-vpcorigin-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

