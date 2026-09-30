package awssecurityhub


// Configuration for the Network Scanning opt-in feature of Security Hub V2.
//
// Network Scanning is available in the AWS commercial partition only; specifying this property in another partition, such as AWS GovCloud (US) or China, fails. This property is desired state: if you remove it from a stack that previously set it, the feature is disabled. If a stack has never set it, the feature is left as-is, so a stack that does not manage Network Scanning will not disable it. Network Scanning requires Security Hub V2 to be enabled in the same account and Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkScanningProperty := &NetworkScanningProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-hubv2-networkscanning.html
//
type CfnHubV2_NetworkScanningProperty struct {
	// Whether the Network Scanning feature is enabled for this account and Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-hubv2-networkscanning.html#cfn-securityhub-hubv2-networkscanning-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}

