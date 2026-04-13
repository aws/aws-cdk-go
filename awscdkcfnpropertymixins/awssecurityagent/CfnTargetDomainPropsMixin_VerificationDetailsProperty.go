package awssecurityagent


// Verification details to verify registered target domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   verificationDetailsProperty := &VerificationDetailsProperty{
//   	DnsTxt: &DnsVerificationProperty{
//   		DnsRecordName: jsii.String("dnsRecordName"),
//   		DnsRecordType: jsii.String("dnsRecordType"),
//   		Token: jsii.String("token"),
//   	},
//   	HttpRoute: &HttpVerificationProperty{
//   		RoutePath: jsii.String("routePath"),
//   		Token: jsii.String("token"),
//   	},
//   	Method: jsii.String("method"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-verificationdetails.html
//
type CfnTargetDomainPropsMixin_VerificationDetailsProperty struct {
	// Represents DNS TXT verification details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-verificationdetails.html#cfn-securityagent-targetdomain-verificationdetails-dnstxt
	//
	DnsTxt interface{} `field:"optional" json:"dnsTxt" yaml:"dnsTxt"`
	// Represents HTTP route verification details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-verificationdetails.html#cfn-securityagent-targetdomain-verificationdetails-httproute
	//
	HttpRoute interface{} `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Type of domain ownership verification method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-targetdomain-verificationdetails.html#cfn-securityagent-targetdomain-verificationdetails-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
}

