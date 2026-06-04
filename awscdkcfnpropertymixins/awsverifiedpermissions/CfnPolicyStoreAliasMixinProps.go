package awsverifiedpermissions


// Properties for CfnPolicyStoreAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPolicyStoreAliasMixinProps := &CfnPolicyStoreAliasMixinProps{
//   	AliasName: jsii.String("aliasName"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html
//
type CfnPolicyStoreAliasMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html#cfn-verifiedpermissions-policystorealias-aliasname
	//
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-verifiedpermissions-policystorealias.html#cfn-verifiedpermissions-policystorealias-policystoreid
	//
	PolicyStoreId *string `field:"optional" json:"policyStoreId" yaml:"policyStoreId"`
}

