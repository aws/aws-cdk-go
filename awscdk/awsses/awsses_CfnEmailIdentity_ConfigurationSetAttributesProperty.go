package awsses


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
	// `CfnEmailIdentity.ConfigurationSetAttributesProperty.ConfigurationSetName`.
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
}

