package awssecuritylake


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	AwsLogSource: &AwsLogSourceProperty{
//   		SourceName: jsii.String("sourceName"),
//   		SourceVersion: jsii.String("sourceVersion"),
//   	},
//   	CustomLogSource: &CustomLogSourceProperty{
//   		SourceName: jsii.String("sourceName"),
//   		SourceVersion: jsii.String("sourceVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-source.html
//
type CfnSubscriber_SourceProperty struct {
	// Amazon Security Lake supports log and event collection for natively supported AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-source.html#cfn-securitylake-subscriber-source-awslogsource
	//
	AwsLogSource interface{} `field:"optional" json:"awsLogSource" yaml:"awsLogSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-source.html#cfn-securitylake-subscriber-source-customlogsource
	//
	CustomLogSource interface{} `field:"optional" json:"customLogSource" yaml:"customLogSource"`
}

