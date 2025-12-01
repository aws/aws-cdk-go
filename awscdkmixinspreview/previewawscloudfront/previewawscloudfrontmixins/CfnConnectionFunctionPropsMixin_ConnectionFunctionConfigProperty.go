package previewawscloudfrontmixins


// Contains configuration information about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectionFunctionConfigProperty := &ConnectionFunctionConfigProperty{
//   	Comment: jsii.String("comment"),
//   	KeyValueStoreAssociations: []interface{}{
//   		&KeyValueStoreAssociationProperty{
//   			KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   		},
//   	},
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html
//
type CfnConnectionFunctionPropsMixin_ConnectionFunctionConfigProperty struct {
	// A comment to describe the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The configuration for the key value store associations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-keyvaluestoreassociations
	//
	KeyValueStoreAssociations interface{} `field:"optional" json:"keyValueStoreAssociations" yaml:"keyValueStoreAssociations"`
	// The function's runtime environment version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-connectionfunction-connectionfunctionconfig.html#cfn-cloudfront-connectionfunction-connectionfunctionconfig-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

