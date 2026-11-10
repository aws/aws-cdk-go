package awslambda


// The compute resource type a network connector is provisioned for.
//
// A connector serves exactly one type; a WebFunction can only use a connector created for
// `WEB_FUNCTION`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//   var code WebFunctionCode
//
//
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("EgressSecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   connector := lambda.NewNetworkConnector(this, jsii.String("Egress"), &NetworkConnectorProps{
//   	Configuration: lambda.NetworkConnectorConfiguration_VpcEgress(&VpcEgressConfiguration{
//   		ComputeType: lambda.NetworkConnectorComputeType_WEB_FUNCTION,
//   		Vpc: *Vpc,
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   		},
//   		SecurityGroups: []ISecurityGroupRef{
//   			securityGroup,
//   		},
//   	}),
//   })
//
//   lambda.NewWebFunction(this, jsii.String("MyWebFunction"), &WebFunctionProps{
//   	Code: Code,
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	EgressNetworkConnector: connector,
//   })
//
type NetworkConnectorComputeType string

const (
	// Lambda WebFunction.
	NetworkConnectorComputeType_WEB_FUNCTION NetworkConnectorComputeType = "WEB_FUNCTION"
	// Lambda MicroVM.
	NetworkConnectorComputeType_MICRO_VM NetworkConnectorComputeType = "MICRO_VM"
)

