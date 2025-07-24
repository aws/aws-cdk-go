package awssupportapp


// Properties for defining a `CfnAccountAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountAliasProps := &CfnAccountAliasProps{
//   	AccountAlias: jsii.String("accountAlias"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-accountalias.html
//
type CfnAccountAliasProps struct {
	// An alias or short name for an AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-accountalias.html#cfn-supportapp-accountalias-accountalias
	//
	AccountAlias *string `field:"required" json:"accountAlias" yaml:"accountAlias"`
}

