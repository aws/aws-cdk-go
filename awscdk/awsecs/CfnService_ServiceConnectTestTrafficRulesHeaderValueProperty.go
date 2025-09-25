package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectTestTrafficRulesHeaderValueProperty := &ServiceConnectTestTrafficRulesHeaderValueProperty{
//   	Exact: jsii.String("exact"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrulesheadervalue.html
//
type CfnService_ServiceConnectTestTrafficRulesHeaderValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrulesheadervalue.html#cfn-ecs-service-serviceconnecttesttrafficrulesheadervalue-exact
	//
	Exact *string `field:"required" json:"exact" yaml:"exact"`
}

