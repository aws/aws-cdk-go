package awsathena


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclConfigurationProperty := &AclConfigurationProperty{
//   	S3AclOption: jsii.String("s3AclOption"),
//   }
//
type CfnWorkGroup_AclConfigurationProperty struct {
	// `CfnWorkGroup.AclConfigurationProperty.S3AclOption`.
	S3AclOption *string `field:"required" json:"s3AclOption" yaml:"s3AclOption"`
}

