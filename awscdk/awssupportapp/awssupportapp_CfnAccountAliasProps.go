package awssupportapp


// Properties for defining a `CfnAccountAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountAliasProps := &cfnAccountAliasProps{
//   	accountAlias: jsii.String("accountAlias"),
//   }
//
type CfnAccountAliasProps struct {
	// `AWS::SupportApp::AccountAlias.AccountAlias`.
	AccountAlias *string `field:"required" json:"accountAlias" yaml:"accountAlias"`
}

