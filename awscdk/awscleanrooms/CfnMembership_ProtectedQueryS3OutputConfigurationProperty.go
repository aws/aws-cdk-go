package awscleanrooms


// Contains the configuration to write the query results to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protectedQueryS3OutputConfigurationProperty := &ProtectedQueryS3OutputConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	ResultFormat: jsii.String("resultFormat"),
//
//   	// the properties below are optional
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	SingleFileOutput: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html
//
type CfnMembership_ProtectedQueryS3OutputConfigurationProperty struct {
	// The S3 bucket to unload the protected query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Intended file format of the result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-resultformat
	//
	ResultFormat *string `field:"required" json:"resultFormat" yaml:"resultFormat"`
	// The S3 prefix to unload the protected query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-singlefileoutput
	//
	SingleFileOutput interface{} `field:"optional" json:"singleFileOutput" yaml:"singleFileOutput"`
}

