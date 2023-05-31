package awsec2


// Construction properties for a GatewayVpcEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gatewayVpcEndpointService iGatewayVpcEndpointService
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   gatewayVpcEndpointProps := &GatewayVpcEndpointProps{
//   	Service: gatewayVpcEndpointService,
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	Subnets: []subnetSelection{
//   		&subnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			SubnetName: jsii.String("subnetName"),
//   			Subnets: []iSubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   		},
//   	},
//   }
//
// Experimental.
type GatewayVpcEndpointProps struct {
	// The service to use for this gateway VPC endpoint.
	// Experimental.
	Service IGatewayVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Where to add endpoint routing.
	//
	// By default, this endpoint will be routable from all subnets in the VPC.
	// Specify a list of subnet selection objects here to be more specific.
	//
	// Example:
	//   declare const vpc: ec2.Vpc;
	//
	//   vpc.addGatewayEndpoint('DynamoDbEndpoint', {
	//     service: ec2.GatewayVpcEndpointAwsService.DYNAMODB,
	//     // Add only to ISOLATED subnets
	//     subnets: [
	//       { subnetType: ec2.SubnetType.PRIVATE_ISOLATED }
	//     ]
	//   });
	//
	// Experimental.
	Subnets *[]*SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The VPC network in which the gateway endpoint will be used.
	// Experimental.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

