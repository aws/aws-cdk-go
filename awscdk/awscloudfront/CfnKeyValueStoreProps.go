package awscloudfront


// Properties for defining a `CfnKeyValueStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyValueStoreProps := &CfnKeyValueStoreProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	ImportSource: &ImportSourceProperty{
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceType: jsii.String("sourceType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html
//
type CfnKeyValueStoreProps struct {
	// The name of the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The import source for the key value store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html#cfn-cloudfront-keyvaluestore-importsource
	//
	ImportSource interface{} `field:"optional" json:"importSource" yaml:"importSource"`
}

