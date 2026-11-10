package awssecurityagent


// The configuration for a self-managed private connection, where you manage your own resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedProperty := &SelfManagedProperty{
//   	ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-selfmanaged.html
//
type CfnPrivateConnection_SelfManagedProperty struct {
	// The identifier or ARN of the resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-selfmanaged.html#cfn-securityagent-privateconnection-selfmanaged-resourceconfigurationid
	//
	ResourceConfigurationId *string `field:"required" json:"resourceConfigurationId" yaml:"resourceConfigurationId"`
}

