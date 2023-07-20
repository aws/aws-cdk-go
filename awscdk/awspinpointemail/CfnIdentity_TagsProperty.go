package awspinpointemail


// An object that defines the tags (keys and values) that you want to associate with the identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsProperty := &TagsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-identity-tags.html
//
type CfnIdentity_TagsProperty struct {
	// One part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag key is 128 characters. The minimum length is 1 character.
	//
	// If you specify tags for the identity, then this value is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-identity-tags.html#cfn-pinpointemail-identity-tags-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The optional part of a key-value pair that defines a tag.
	//
	// The maximum length of a tag value is 256 characters. The minimum length is 0 characters. If you don’t want a resource to have a specific tag value, don’t specify a value for this parameter. Amazon Pinpoint will set the value to an empty string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-identity-tags.html#cfn-pinpointemail-identity-tags-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

