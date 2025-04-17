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
type CnameInstanceBaseProps struct {
	// Custom attributes of the instance.
	// Default: none.
	//
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	// Default: Automatically generated name.
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// If the service configuration includes a CNAME record, the domain name that you want Route 53 to return in response to DNS queries, for example, example.com. This value is required if the service specified by ServiceId includes settings for an CNAME record.
	InstanceCname *string `field:"required" json:"instanceCname" yaml:"instanceCname"`
}

