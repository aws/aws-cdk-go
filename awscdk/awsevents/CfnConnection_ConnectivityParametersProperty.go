package awsevents


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectivityParametersProperty := &ConnectivityParametersProperty{
//   	ResourceParameters: &ResourceParametersProperty{
//   		ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   		// the properties below are optional
//   		ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-connectivityparameters.html
//
type CfnConnection_ConnectivityParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-connectivityparameters.html#cfn-events-connection-connectivityparameters-resourceparameters
	//
	ResourceParameters interface{} `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

