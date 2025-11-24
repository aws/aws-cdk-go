package mixinsawsredshift


// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingPropertiesProperty := &LoggingPropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//   	LogDestinationType: jsii.String("logDestinationType"),
//   	LogExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
//
type CfnClusterPropsMixin_LoggingPropertiesProperty struct {
	// The name of an existing S3 bucket where the log files are to be stored.
	//
	// Constraints:
	//
	// - Must be in the same region as the cluster
	// - The cluster must have read bucket and put object permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The log destination type.
	//
	// An enum with possible values of `s3` and `cloudwatch` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-logdestinationtype
	//
	LogDestinationType *string `field:"optional" json:"logDestinationType" yaml:"logDestinationType"`
	// The collection of exported log types.
	//
	// Possible values are `connectionlog` , `useractivitylog` , and `userlog` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-logexports
	//
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// The prefix applied to the log file names.
	//
	// Valid characters are any letter from any language, any whitespace character, any numeric character, and the following characters: underscore ( `_` ), period ( `.` ), colon ( `:` ), slash ( `/` ), equal ( `=` ), plus ( `+` ), backslash ( `\` ), hyphen ( `-` ), at symbol ( `@` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

