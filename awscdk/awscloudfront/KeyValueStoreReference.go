package awscloudfront


// A reference to a KeyValueStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValueStoreReference := &KeyValueStoreReference{
//   	KeyValueStoreArn: jsii.String("keyValueStoreArn"),
//   	KeyValueStoreName: jsii.String("keyValueStoreName"),
//   }
//
type KeyValueStoreReference struct {
	// The ARN of the KeyValueStore resource.
	KeyValueStoreArn *string `field:"required" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
	// The Name of the KeyValueStore resource.
	KeyValueStoreName *string `field:"required" json:"keyValueStoreName" yaml:"keyValueStoreName"`
}

