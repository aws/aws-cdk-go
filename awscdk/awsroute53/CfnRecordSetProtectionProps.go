package awsroute53


// Properties for defining a `CfnRecordSetProtection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordSetProtectionProps := &CfnRecordSetProtectionProps{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	RecordName: jsii.String("recordName"),
//   	RecordType: jsii.String("recordType"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html
//
type CfnRecordSetProtectionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-hostedzoneid
	//
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-recordname
	//
	RecordName *string `field:"required" json:"recordName" yaml:"recordName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-recordtype
	//
	RecordType *string `field:"required" json:"recordType" yaml:"recordType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

