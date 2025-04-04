package awsses


// The union type representing the allowed types for the left hand side of a TLS condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressTlsProtocolToEvaluateProperty := &IngressTlsProtocolToEvaluateProperty{
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocoltoevaluate.html
//
type CfnMailManagerTrafficPolicy_IngressTlsProtocolToEvaluateProperty struct {
	// The enum type representing the allowed attribute types for the TLS condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocoltoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingresstlsprotocoltoevaluate-attribute
	//
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}

