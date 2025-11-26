package previewawscloudfrontmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConnectionFunctionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectionFunctionMixinProps := &CfnConnectionFunctionMixinProps{
//   	AutoPublish: jsii.Boolean(false),
//   	ConnectionFunctionCode: jsii.String("connectionFunctionCode"),
//   	ConnectionFunctionConfig: &ConnectionFunctionConfigProperty{
//   		Comment: jsii.String("comment"),
//   		KeyValueStoreAssociations: []interface{}{
//   			&KeyValueStoreAssociationProperty{
//   				KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   			},
//   		},
//   		Runtime: jsii.String("runtime"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html
//
type CfnConnectionFunctionMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html#cfn-cloudfront-connectionfunction-autopublish
	//
	// Default: - false.
	//
	AutoPublish interface{} `field:"optional" json:"autoPublish" yaml:"autoPublish"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html#cfn-cloudfront-connectionfunction-connectionfunctioncode
	//
	ConnectionFunctionCode *string `field:"optional" json:"connectionFunctionCode" yaml:"connectionFunctionCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig
	//
	ConnectionFunctionConfig interface{} `field:"optional" json:"connectionFunctionConfig" yaml:"connectionFunctionConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html#cfn-cloudfront-connectionfunction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-connectionfunction.html#cfn-cloudfront-connectionfunction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

