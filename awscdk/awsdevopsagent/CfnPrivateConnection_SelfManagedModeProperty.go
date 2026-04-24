package awsdevopsagent


// Configuration for a self-managed Private Connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedModeProperty := &SelfManagedModeProperty{
//   	ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-selfmanagedmode.html
//
type CfnPrivateConnection_SelfManagedModeProperty struct {
	// The ARN of the Resource Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-selfmanagedmode.html#cfn-devopsagent-privateconnection-selfmanagedmode-resourceconfigurationid
	//
	ResourceConfigurationId *string `field:"required" json:"resourceConfigurationId" yaml:"resourceConfigurationId"`
}

