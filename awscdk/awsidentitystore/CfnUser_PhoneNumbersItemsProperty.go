package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phoneNumbersItemsProperty := &PhoneNumbersItemsProperty{
//   	Primary: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-phonenumbersitems.html
//
type CfnUser_PhoneNumbersItemsProperty struct {
	// Whether this is the primary phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-phonenumbersitems.html#cfn-identitystore-user-phonenumbersitems-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-phonenumbersitems.html#cfn-identitystore-user-phonenumbersitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-phonenumbersitems.html#cfn-identitystore-user-phonenumbersitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

