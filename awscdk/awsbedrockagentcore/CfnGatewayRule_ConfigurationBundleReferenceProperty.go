package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationBundleReferenceProperty := &ConfigurationBundleReferenceProperty{
//   	BundleArn: jsii.String("bundleArn"),
//   	BundleVersion: jsii.String("bundleVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundlereference.html
//
type CfnGatewayRule_ConfigurationBundleReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundlereference.html#cfn-bedrockagentcore-gatewayrule-configurationbundlereference-bundlearn
	//
	BundleArn *string `field:"required" json:"bundleArn" yaml:"bundleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-configurationbundlereference.html#cfn-bedrockagentcore-gatewayrule-configurationbundlereference-bundleversion
	//
	BundleVersion *string `field:"required" json:"bundleVersion" yaml:"bundleVersion"`
}

