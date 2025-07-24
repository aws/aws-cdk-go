package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrules.html#cfn-ecs-service-serviceconnecttesttrafficrules-header
	//
	Header interface{} `field:"required" json:"header" yaml:"header"`
}

