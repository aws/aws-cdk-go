package awsroute53


// Properties for CfnRecordSetProtectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRecordSetProtectionMixinProps := &CfnRecordSetProtectionMixinProps{
//   	Comment: jsii.String("comment"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	RecordName: jsii.String("recordName"),
//   	RecordType: jsii.String("recordType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html
//
type CfnRecordSetProtectionMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-recordname
	//
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html#cfn-route53-recordsetprotection-recordtype
	//
	RecordType *string `field:"optional" json:"recordType" yaml:"recordType"`
}

