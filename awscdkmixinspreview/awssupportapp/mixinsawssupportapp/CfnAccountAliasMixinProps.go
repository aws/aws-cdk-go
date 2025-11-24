package mixinsawssupportapp


// Properties for CfnAccountAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccountAliasMixinProps := &CfnAccountAliasMixinProps{
//   	AccountAlias: jsii.String("accountAlias"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-accountalias.html
//
type CfnAccountAliasMixinProps struct {
	// An alias or short name for an AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-accountalias.html#cfn-supportapp-accountalias-accountalias
	//
	AccountAlias *string `field:"optional" json:"accountAlias" yaml:"accountAlias"`
}

