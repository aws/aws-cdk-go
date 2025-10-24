package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDelivery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryProps := &CfnDeliveryProps{
//   	DeliveryDestinationArn: jsii.String("deliveryDestinationArn"),
//   	DeliverySourceName: jsii.String("deliverySourceName"),
//
//   	// the properties below are optional
//   	FieldDelimiter: jsii.String("fieldDelimiter"),
//   	RecordFields: []*string{
//   		jsii.String("recordFields"),
//   	},
//   	S3EnableHiveCompatiblePath: jsii.Boolean(false),
//   	S3SuffixPath: jsii.String("s3SuffixPath"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html
//
type CfnDeliveryProps struct {
	// The ARN of the delivery destination that is associated with this delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-deliverydestinationarn
	//
	DeliveryDestinationArn *string `field:"required" json:"deliveryDestinationArn" yaml:"deliveryDestinationArn"`
	// The name of the delivery source that is associated with this delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-deliverysourcename
	//
	DeliverySourceName *string `field:"required" json:"deliverySourceName" yaml:"deliverySourceName"`
	// The field delimiter that is used between record fields when the final output format of a delivery is in `Plain` , `W3C` , or `Raw` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-fielddelimiter
	//
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// The list of record fields to be delivered to the destination, in order.
	//
	// If the delivery's log source has mandatory fields, they must be included in this list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-recordfields
	//
	RecordFields *[]*string `field:"optional" json:"recordFields" yaml:"recordFields"`
	// Use this parameter to cause the S3 objects that contain delivered logs to use a prefix structure that allows for integration with Apache Hive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-s3enablehivecompatiblepath
	//
	S3EnableHiveCompatiblePath interface{} `field:"optional" json:"s3EnableHiveCompatiblePath" yaml:"s3EnableHiveCompatiblePath"`
	// Use this to reconfigure the S3 object prefix to contain either static or variable sections.
	//
	// The valid variables to use in the suffix path will vary by each log source. To find the values supported for the suffix path for each log source, use the [DescribeConfigurationTemplates](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeConfigurationTemplates.html) operation and check the `allowedSuffixPathFields` field in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-s3suffixpath
	//
	S3SuffixPath *string `field:"optional" json:"s3SuffixPath" yaml:"s3SuffixPath"`
	// An array of key-value pairs to apply to the delivery.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

