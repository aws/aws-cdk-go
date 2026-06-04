package awsverifiedpermissions


// Properties for defining a `CfnPolicyStoreAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyStoreAliasProps := &CfnPolicyStoreAliasProps{
//   	AliasName: jsii.String("aliasName"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html
//
type CfnPolicyStoreAliasProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html#cfn-verifiedpermissions-policystorealias-aliasname
	//
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html#cfn-verifiedpermissions-policystorealias-policystoreid
	//
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

