package awskendra


// Specifies access control list files for the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlListConfigurationProperty := &accessControlListConfigurationProperty{
//   	keyPath: jsii.String("keyPath"),
//   }
//
type CfnDataSource_AccessControlListConfigurationProperty struct {
	// Path to the AWS S3 bucket that contains the access control list files.
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
}

