package awssynthetics


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserConfigProperty := &BrowserConfigProperty{
//   	BrowserType: jsii.String("browserType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-browserconfig.html
//
type CfnCanary_BrowserConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-browserconfig.html#cfn-synthetics-canary-browserconfig-browsertype
	//
	BrowserType *string `field:"required" json:"browserType" yaml:"browserType"`
}

