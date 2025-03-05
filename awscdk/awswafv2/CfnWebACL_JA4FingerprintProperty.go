package awswafv2


// Includes the JA4 fingerprint of a web request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jA4FingerprintProperty := &JA4FingerprintProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ja4fingerprint.html
//
type CfnWebACL_JA4FingerprintProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ja4fingerprint.html#cfn-wafv2-webacl-ja4fingerprint-fallbackbehavior
	//
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

