package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineStoreSecurityConfigProperty := &onlineStoreSecurityConfigProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFeatureGroup_OnlineStoreSecurityConfigProperty struct {
	// `CfnFeatureGroup.OnlineStoreSecurityConfigProperty.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

