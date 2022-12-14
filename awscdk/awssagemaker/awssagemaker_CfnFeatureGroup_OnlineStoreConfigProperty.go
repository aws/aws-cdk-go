package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineStoreConfigProperty := &onlineStoreConfigProperty{
//   	enableOnlineStore: jsii.Boolean(false),
//   	securityConfig: &onlineStoreSecurityConfigProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
type CfnFeatureGroup_OnlineStoreConfigProperty struct {
	// `CfnFeatureGroup.OnlineStoreConfigProperty.EnableOnlineStore`.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// `CfnFeatureGroup.OnlineStoreConfigProperty.SecurityConfig`.
	SecurityConfig interface{} `field:"optional" json:"securityConfig" yaml:"securityConfig"`
}

