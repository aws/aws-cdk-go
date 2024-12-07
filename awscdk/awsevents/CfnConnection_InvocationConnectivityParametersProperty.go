package awsevents


// The private resource the HTTP request will be sent to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   invocationConnectivityParametersProperty := &InvocationConnectivityParametersProperty{
//   	ResourceParameters: &ResourceParametersProperty{
//   		ResourceConfigurationArn: jsii.String("resourceConfigurationArn"),
//
//   		// the properties below are optional
//   		ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-invocationconnectivityparameters.html
//
type CfnConnection_InvocationConnectivityParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-invocationconnectivityparameters.html#cfn-events-connection-invocationconnectivityparameters-resourceparameters
	//
	ResourceParameters interface{} `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

