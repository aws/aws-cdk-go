package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   photosItemsProperty := &PhotosItemsProperty{
//   	Display: jsii.String("display"),
//   	Primary: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-photositems.html
//
type CfnUserPropsMixin_PhotosItemsProperty struct {
	// A display name for the photo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-photositems.html#cfn-identitystore-user-photositems-display
	//
	Display *string `field:"optional" json:"display" yaml:"display"`
	// Whether this is the primary photo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-photositems.html#cfn-identitystore-user-photositems-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of photo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-photositems.html#cfn-identitystore-user-photositems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The photo data or URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-photositems.html#cfn-identitystore-user-photositems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

