package previewawsdatazonemixins


// The user policy grant principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var allUsersGrantFilter interface{}
//
//   userPolicyGrantPrincipalProperty := &UserPolicyGrantPrincipalProperty{
//   	AllUsersGrantFilter: allUsersGrantFilter,
//   	UserIdentifier: jsii.String("userIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-userpolicygrantprincipal.html
//
type CfnPolicyGrantPropsMixin_UserPolicyGrantPrincipalProperty struct {
	// The all users grant filter of the user policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-userpolicygrantprincipal.html#cfn-datazone-policygrant-userpolicygrantprincipal-allusersgrantfilter
	//
	AllUsersGrantFilter interface{} `field:"optional" json:"allUsersGrantFilter" yaml:"allUsersGrantFilter"`
	// The user ID of the user policy grant principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-policygrant-userpolicygrantprincipal.html#cfn-datazone-policygrant-userpolicygrantprincipal-useridentifier
	//
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

