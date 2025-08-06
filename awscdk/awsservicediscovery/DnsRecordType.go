package awsservicediscovery


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &PublicDnsNamespaceProps{
//   	Name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
//   	Name: jsii.String("foo"),
//   	DnsRecordType: servicediscovery.DnsRecordType_CNAME,
//   	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
//   })
//
//   service.RegisterCnameInstance(jsii.String("CnameInstance"), &CnameInstanceBaseProps{
//   	InstanceCname: jsii.String("service.pizza"),
//   })
//
//   app.Synth()
//
type DnsRecordType string

const (
	// An A record.
	DnsRecordType_A DnsRecordType = "A"
	// An AAAA record.
	DnsRecordType_AAAA DnsRecordType = "AAAA"
	// Both an A and AAAA record.
	DnsRecordType_A_AAAA DnsRecordType = "A_AAAA"
	// A Srv record.
	DnsRecordType_SRV DnsRecordType = "SRV"
	// A CNAME record.
	DnsRecordType_CNAME DnsRecordType = "CNAME"
)

