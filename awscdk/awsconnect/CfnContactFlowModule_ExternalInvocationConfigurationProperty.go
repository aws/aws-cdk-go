package awsconnect


// The external invocation configuration for the flow module.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalInvocationConfigurationProperty := &ExternalInvocationConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-contactflowmodule-externalinvocationconfiguration.html
//
type CfnContactFlowModule_ExternalInvocationConfigurationProperty struct {
	// Enable external invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-contactflowmodule-externalinvocationconfiguration.html#cfn-connect-contactflowmodule-externalinvocationconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

