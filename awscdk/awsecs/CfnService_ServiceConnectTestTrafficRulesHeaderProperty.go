package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectTestTrafficRulesHeaderProperty := &ServiceConnectTestTrafficRulesHeaderProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Value: &ServiceConnectTestTrafficRulesHeaderValueProperty{
//   		Exact: jsii.String("exact"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrulesheader.html
//
type CfnService_ServiceConnectTestTrafficRulesHeaderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrulesheader.html#cfn-ecs-service-serviceconnecttesttrafficrulesheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnecttesttrafficrulesheader.html#cfn-ecs-service-serviceconnecttesttrafficrulesheader-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

