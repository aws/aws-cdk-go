package awsses


// Used to associate a configuration set with an email identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationSetAttributesProperty := &configurationSetAttributesProperty{
//   	configurationSetName: jsii.String("configurationSetName"),
//   }
//
type CfnEmailIdentity_ConfigurationSetAttributesProperty struct {
	// The configuration set to associate with an email identity.
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
}

