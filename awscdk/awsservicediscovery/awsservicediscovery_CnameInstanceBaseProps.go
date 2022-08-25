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
type CnameInstanceBaseProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// If the service configuration includes a CNAME record, the domain name that you want Route 53 to return in response to DNS queries, for example, example.com. This value is required if the service specified by ServiceId includes settings for an CNAME record.
	InstanceCname *string `field:"required" json:"instanceCname" yaml:"instanceCname"`
}

