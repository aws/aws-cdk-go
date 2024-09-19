package awsec2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsec2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The gateway or endpoint targeted by the route.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
//   })
//
//   igw := awsec2alpha.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
//   	Vpc: myVpc,
//   })
//   awsec2alpha.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": igw,
//   	},
//   })
//
// Experimental.
type RouteTargetType interface {
	// The endpoint route target.
	//
	// This is used for targets such as
	// VPC endpoints.
	// Default: - target is not set to an endpoint, in this case a gateway is needed.
	//
	// Experimental.
	Endpoint() awsec2.IVpcEndpoint
	// The gateway route target.
	//
	// This is used for targets such as
	// egress-only internet gateway or VPC peering connection.
	// Default: - target is not set to a gateway, in this case an endpoint is needed.
	//
	// Experimental.
	Gateway() IRouteTarget
}

// The jsii proxy struct for RouteTargetType
type jsiiProxy_RouteTargetType struct {
	_ byte // padding
}

func (j *jsiiProxy_RouteTargetType) Endpoint() awsec2.IVpcEndpoint {
	var returns awsec2.IVpcEndpoint
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteTargetType) Gateway() IRouteTarget {
	var returns IRouteTarget
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}


// Experimental.
func NewRouteTargetType(props *RouteTargetProps) RouteTargetType {
	_init_.Initialize()

	if err := validateNewRouteTargetTypeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RouteTargetType{}

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.RouteTargetType",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewRouteTargetType_Override(r RouteTargetType, props *RouteTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.RouteTargetType",
		[]interface{}{props},
		r,
	)
}

