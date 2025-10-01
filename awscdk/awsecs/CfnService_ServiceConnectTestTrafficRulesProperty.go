package awsecs


// The test traffic routing configuration for Amazon ECS blue/green deployments.
//
// This configuration allows you to define rules for routing specific traffic to the new service revision during the deployment process, allowing for safe testing before full production traffic shift.
//
// For more information, see [Service Connect for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect-blue-green.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectTestTrafficRulesProperty := &ServiceConnectTestTrafficRulesProperty{
//   	Header: &ServiceConnectTestTrafficRulesHeaderProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Value: &ServiceConnectTestTrafficRulesHeaderValueProperty{
//   			Exact: jsii.String("exact"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrules.html
//
type CfnService_ServiceConnectTestTrafficRulesProperty struct {
	// The HTTP header-based routing rules that determine which requests should be routed to the new service version during blue/green deployment testing.
	//
	// These rules provide fine-grained control over test traffic routing based on request headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrules.html#cfn-ecs-service-serviceconnecttesttrafficrules-header
	//
	Header interface{} `field:"required" json:"header" yaml:"header"`
}

