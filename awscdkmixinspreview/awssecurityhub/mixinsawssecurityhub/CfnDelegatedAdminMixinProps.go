package mixinsawssecurityhub


// Properties for CfnDelegatedAdminPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDelegatedAdminMixinProps := &CfnDelegatedAdminMixinProps{
//   	AdminAccountId: jsii.String("adminAccountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-delegatedadmin.html
//
type CfnDelegatedAdminMixinProps struct {
	// The AWS account identifier of the account to designate as the Security Hub administrator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-delegatedadmin.html#cfn-securityhub-delegatedadmin-adminaccountid
	//
	AdminAccountId *string `field:"optional" json:"adminAccountId" yaml:"adminAccountId"`
}

