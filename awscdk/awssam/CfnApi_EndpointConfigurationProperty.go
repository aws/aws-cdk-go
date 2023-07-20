package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigurationProperty := &EndpointConfigurationProperty{
//   	Type: jsii.String("type"),
//   	VpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-endpointconfiguration.html
//
type CfnApi_EndpointConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-endpointconfiguration.html#cfn-serverless-api-endpointconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-endpointconfiguration.html#cfn-serverless-api-endpointconfiguration-vpcendpointids
	//
	VpcEndpointIds *[]*string `field:"optional" json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

