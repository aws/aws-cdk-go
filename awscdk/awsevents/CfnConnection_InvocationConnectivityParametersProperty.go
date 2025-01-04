package awsevents


// For connections to private APIs, the parameters to use for invoking the API.
//
// For more information, see [Connecting to private APIs](https://docs.aws.amazon.com/eventbridge/latest/userguide/connection-private.html) in the **Amazon EventBridge User Guide** .
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
	// The parameters for EventBridge to use when invoking the resource endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-invocationconnectivityparameters.html#cfn-events-connection-invocationconnectivityparameters-resourceparameters
	//
	ResourceParameters interface{} `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

