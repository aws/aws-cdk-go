package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   staticOverrideProperty := &StaticOverrideProperty{
//   	BundleArn: jsii.String("bundleArn"),
//   	BundleVersion: jsii.String("bundleVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticoverride.html
//
type CfnGatewayRulePropsMixin_StaticOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticoverride.html#cfn-bedrockagentcore-gatewayrule-staticoverride-bundlearn
	//
	BundleArn *string `field:"optional" json:"bundleArn" yaml:"bundleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticoverride.html#cfn-bedrockagentcore-gatewayrule-staticoverride-bundleversion
	//
	BundleVersion *string `field:"optional" json:"bundleVersion" yaml:"bundleVersion"`
}

