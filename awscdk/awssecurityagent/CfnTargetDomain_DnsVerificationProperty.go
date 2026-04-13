package awssecurityagent


// Represents DNS TXT verification details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsVerificationProperty := &DnsVerificationProperty{
//   	DnsRecordName: jsii.String("dnsRecordName"),
//   	DnsRecordType: jsii.String("dnsRecordType"),
//   	Token: jsii.String("token"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-dnsverification.html
//
type CfnTargetDomain_DnsVerificationProperty struct {
	// Record name to be added in DNS for target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-dnsverification.html#cfn-securityagent-targetdomain-dnsverification-dnsrecordname
	//
	DnsRecordName *string `field:"optional" json:"dnsRecordName" yaml:"dnsRecordName"`
	// Type of record to be added in DNS for target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-dnsverification.html#cfn-securityagent-targetdomain-dnsverification-dnsrecordtype
	//
	DnsRecordType *string `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// Token used to verify domain ownership.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-dnsverification.html#cfn-securityagent-targetdomain-dnsverification-token
	//
	Token *string `field:"optional" json:"token" yaml:"token"`
}

