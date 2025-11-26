package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressPathSummaryProperty := &IngressPathSummaryProperty{
//   	AccessType: jsii.String("accessType"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html
//
type CfnExpressGatewayServicePropsMixin_IngressPathSummaryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html#cfn-ecs-expressgatewayservice-ingresspathsummary-accesstype
	//
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspathsummary.html#cfn-ecs-expressgatewayservice-ingresspathsummary-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

