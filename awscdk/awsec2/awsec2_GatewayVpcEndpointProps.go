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
//   gatewayVpcEndpointProps := &gatewayVpcEndpointProps{
//   	service: gatewayVpcEndpointService,
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	subnets: []subnetSelection{
//   		&subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type GatewayVpcEndpointProps struct {
	// The service to use for this gateway VPC endpoint.
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
	Subnets *[]*SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The VPC network in which the gateway endpoint will be used.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

