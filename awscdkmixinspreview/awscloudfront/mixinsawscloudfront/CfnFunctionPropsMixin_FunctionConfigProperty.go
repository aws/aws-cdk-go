package mixinsawscloudfront


// Contains configuration information about a CloudFront function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   functionConfigProperty := &FunctionConfigProperty{
//   	Comment: jsii.String("comment"),
//   	KeyValueStoreAssociations: []interface{}{
//   		&KeyValueStoreAssociationProperty{
//   			KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   		},
//   	},
//   	Runtime: jsii.String("runtime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html
//
type CfnFunctionPropsMixin_FunctionConfigProperty struct {
	// A comment to describe the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html#cfn-cloudfront-function-functionconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The configuration for the key value store associations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html#cfn-cloudfront-function-functionconfig-keyvaluestoreassociations
	//
	KeyValueStoreAssociations interface{} `field:"optional" json:"keyValueStoreAssociations" yaml:"keyValueStoreAssociations"`
	// The function's runtime environment version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-functionconfig.html#cfn-cloudfront-function-functionconfig-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

