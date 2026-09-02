package awsmgn


// Configuration for a migration source environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	SourceEnvironment: jsii.String("sourceEnvironment"),
//   	SourceS3Configuration: &SourceS3ConfigurationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sourceconfiguration.html
//
type CfnNetworkMigrationDefinition_SourceConfigurationProperty struct {
	// The source environment type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sourceconfiguration.html#cfn-mgn-networkmigrationdefinition-sourceconfiguration-sourceenvironment
	//
	SourceEnvironment *string `field:"required" json:"sourceEnvironment" yaml:"sourceEnvironment"`
	// S3 configuration for source network data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-networkmigrationdefinition-sourceconfiguration.html#cfn-mgn-networkmigrationdefinition-sourceconfiguration-sources3configuration
	//
	SourceS3Configuration interface{} `field:"required" json:"sourceS3Configuration" yaml:"sourceS3Configuration"`
}

