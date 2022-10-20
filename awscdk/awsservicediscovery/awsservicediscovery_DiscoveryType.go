package awsservicediscovery


// Specifies information about the discovery type of a service.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   vpc := ec2.NewVpc(stack, jsii.String("Vpc"), &vpcProps{
//   	maxAzs: jsii.Number(2),
//   })
//
//   namespace := servicediscovery.NewPrivateDnsNamespace(stack, jsii.String("Namespace"), &privateDnsNamespaceProps{
//   	name: jsii.String("boobar.com"),
//   	vpc: vpc,
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	dnsRecordType: servicediscovery.dnsRecordType_A_AAAA,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   	loadBalancer: jsii.Boolean(true),
//   })
//
//   loadbalancer := elbv2.NewApplicationLoadBalancer(stack, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   service.registerLoadBalancer(jsii.String("Loadbalancer"), loadbalancer)
//
//   arnService := namespace.createService(jsii.String("ArnService"), &dnsServiceProps{
//   	discoveryType: servicediscovery.discoveryType_API,
//   })
//
//   arnService.registerNonIpInstance(jsii.String("NonIpInstance"), &nonIpInstanceBaseProps{
//   	customAttributes: map[string]*string{
//   		"arn": jsii.String("arn://"),
//   	},
//   })
//
//   app.synth()
//
type DiscoveryType string

const (
	// Instances are discoverable via API only.
	DiscoveryType_API DiscoveryType = "API"
	// Instances are discoverable via DNS or API.
	DiscoveryType_DNS_AND_API DiscoveryType = "DNS_AND_API"
)

