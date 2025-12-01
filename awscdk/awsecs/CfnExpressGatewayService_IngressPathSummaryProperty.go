package awsecs


// The entry point into an Express service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressPathSummaryProperty := &IngressPathSummaryProperty{
//   	AccessType: jsii.String("accessType"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html
//
type CfnExpressGatewayService_IngressPathSummaryProperty struct {
	// The type of access to the endpoint for the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html#cfn-ecs-expressgatewayservice-ingresspathsummary-accesstype
	//
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The endpoint for access to the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html#cfn-ecs-expressgatewayservice-ingresspathsummary-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

