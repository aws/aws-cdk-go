package awsservicediscovery


// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
//   	name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	name: jsii.String("foo"),
//   	dnsRecordType: servicediscovery.dnsRecordType_CNAME,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   })
//
//   service.registerCnameInstance(jsii.String("CnameInstance"), &cnameInstanceBaseProps{
//   	instanceCname: jsii.String("service.pizza"),
//   })
//
//   app.synth()
//
// Experimental.
type DnsRecordType string

const (
	// An A record.
	// Experimental.
	DnsRecordType_A DnsRecordType = "A"
	// An AAAA record.
	// Experimental.
	DnsRecordType_AAAA DnsRecordType = "AAAA"
	// Both an A and AAAA record.
	// Experimental.
	DnsRecordType_A_AAAA DnsRecordType = "A_AAAA"
	// A Srv record.
	// Experimental.
	DnsRecordType_SRV DnsRecordType = "SRV"
	// A CNAME record.
	// Experimental.
	DnsRecordType_CNAME DnsRecordType = "CNAME"
)

