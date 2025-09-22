package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   vpc := ec2.NewVpc(stack, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(2),
//   })
//
//   namespace := servicediscovery.NewPrivateDnsNamespace(stack, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
//   	Name: jsii.String("boobar.com"),
//   	Vpc: Vpc,
//   })
//
//   service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
//   	DnsRecordType: servicediscovery.DnsRecordType_A_AAAA,
//   	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
//   	LoadBalancer: jsii.Boolean(true),
//   })
//
//   loadbalancer := elbv2.NewApplicationLoadBalancer(stack, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   service.RegisterLoadBalancer(jsii.String("Loadbalancer"), loadbalancer)
//
//   arnService := namespace.CreateService(jsii.String("ArnService"), &DnsServiceProps{
//   	DiscoveryType: servicediscovery.DiscoveryType_API,
//   })
//
//   arnService.RegisterNonIpInstance(jsii.String("NonIpInstance"), &NonIpInstanceBaseProps{
//   	CustomAttributes: map[string]*string{
//   		"arn": jsii.String("arn://"),
//   	},
//   })
//
//   app.Synth()
//
type PrivateDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon VPC that you want to associate the namespace with.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

