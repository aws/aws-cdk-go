package awscertificatemanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsItemsProperty := &TagsItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-tagsitems.html
//
type CfnAcmeEndpoint_TagsItemsProperty struct {
	// The key name of the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-tagsitems.html#cfn-certificatemanager-acmeendpoint-tagsitems-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-tagsitems.html#cfn-certificatemanager-acmeendpoint-tagsitems-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

