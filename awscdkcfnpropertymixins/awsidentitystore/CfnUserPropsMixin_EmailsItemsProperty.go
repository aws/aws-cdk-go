package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   emailsItemsProperty := &EmailsItemsProperty{
//   	Primary: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-emailsitems.html
//
type CfnUserPropsMixin_EmailsItemsProperty struct {
	// Whether this is the primary email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-emailsitems.html#cfn-identitystore-user-emailsitems-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-emailsitems.html#cfn-identitystore-user-emailsitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-emailsitems.html#cfn-identitystore-user-emailsitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

