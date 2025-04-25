package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressIpv6ToEvaluateProperty := &IngressIpv6ToEvaluateProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6toevaluate.html
//
type CfnMailManagerTrafficPolicy_IngressIpv6ToEvaluateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressipv6toevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressipv6toevaluate-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

