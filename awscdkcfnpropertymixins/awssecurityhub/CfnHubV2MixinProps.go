package awssecurityhub


// Properties for CfnHubV2PropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHubV2MixinProps := &CfnHubV2MixinProps{
//   	NetworkScanning: &NetworkScanningProperty{
//   		Status: jsii.String("status"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hubv2.html
//
type CfnHubV2MixinProps struct {
	// Configuration for the Network Scanning opt-in feature of Security Hub V2.
	//
	// Network Scanning is available in the AWS commercial partition only; specifying this property in another partition, such as AWS GovCloud (US) or China, fails. This property is desired state: if you remove it from a stack that previously set it, the feature is disabled. If a stack has never set it, the feature is left as-is, so a stack that does not manage Network Scanning will not disable it. Network Scanning requires Security Hub V2 to be enabled in the same account and Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hubv2.html#cfn-securityhub-hubv2-networkscanning
	//
	NetworkScanning interface{} `field:"optional" json:"networkScanning" yaml:"networkScanning"`
	// The tags to add to the hub V2 resource when you enable Security Hub CSPM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hubv2.html#cfn-securityhub-hubv2-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

