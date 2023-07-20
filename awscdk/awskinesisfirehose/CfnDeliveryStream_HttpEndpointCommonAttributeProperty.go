package awskinesisfirehose


// Describes the metadata that's delivered to the specified HTTP endpoint destination.
//
// Kinesis Firehose supports any custom HTTP endpoint or HTTP endpoints owned by supported third-party service providers, including Datadog, MongoDB, and New Relic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpEndpointCommonAttributeProperty := &HttpEndpointCommonAttributeProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	AttributeValue: jsii.String("attributeValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html
//
type CfnDeliveryStream_HttpEndpointCommonAttributeProperty struct {
	// The name of the HTTP endpoint common attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html#cfn-kinesisfirehose-deliverystream-httpendpointcommonattribute-attributename
	//
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The value of the HTTP endpoint common attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html#cfn-kinesisfirehose-deliverystream-httpendpointcommonattribute-attributevalue
	//
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

