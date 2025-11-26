package previewawscleanroomsmixins


// Contains the configuration to write the query results to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   protectedQueryS3OutputConfigurationProperty := &ProtectedQueryS3OutputConfigurationProperty{
//   	Bucket: jsii.String("bucket"),
//   	KeyPrefix: jsii.String("keyPrefix"),
//   	ResultFormat: jsii.String("resultFormat"),
//   	SingleFileOutput: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html
//
type CfnMembershipPropsMixin_ProtectedQueryS3OutputConfigurationProperty struct {
	// The S3 bucket to unload the protected query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix to unload the protected query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Intended file format of the result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-resultformat
	//
	ResultFormat *string `field:"optional" json:"resultFormat" yaml:"resultFormat"`
	// Indicates whether files should be output as a single file ( `TRUE` ) or output as multiple files ( `FALSE` ).
	//
	// This parameter is only supported for analyses with the Spark analytics engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedquerys3outputconfiguration.html#cfn-cleanrooms-membership-protectedquerys3outputconfiguration-singlefileoutput
	//
	SingleFileOutput interface{} `field:"optional" json:"singleFileOutput" yaml:"singleFileOutput"`
}

