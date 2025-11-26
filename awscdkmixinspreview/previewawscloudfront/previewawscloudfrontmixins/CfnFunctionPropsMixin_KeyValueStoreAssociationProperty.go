package previewawscloudfrontmixins


// The key value store association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyValueStoreAssociationProperty := &KeyValueStoreAssociationProperty{
//   	KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-keyvaluestoreassociation.html
//
type CfnFunctionPropsMixin_KeyValueStoreAssociationProperty struct {
	// The Amazon Resource Name (ARN) of the key value store association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-function-keyvaluestoreassociation.html#cfn-cloudfront-function-keyvaluestoreassociation-keyvaluestorearn
	//
	KeyValueStoreArn *string `field:"optional" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
}

