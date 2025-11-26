package previewawsiotsitewisemixins


// Contains information for a user identity in an access policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userProperty := &UserProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-user.html
//
type CfnAccessPolicyPropsMixin_UserProperty struct {
	// The IAM Identity Center ID of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-user.html#cfn-iotsitewise-accesspolicy-user-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

