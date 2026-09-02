package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   rolesItemsProperty := &RolesItemsProperty{
//   	Primary: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-rolesitems.html
//
type CfnUserPropsMixin_RolesItemsProperty struct {
	// Whether this is the primary role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-rolesitems.html#cfn-identitystore-user-rolesitems-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The type of role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-rolesitems.html#cfn-identitystore-user-rolesitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The role name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-rolesitems.html#cfn-identitystore-user-rolesitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

