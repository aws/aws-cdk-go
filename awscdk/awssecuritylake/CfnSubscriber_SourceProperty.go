package awssecuritylake


// Sources are logs and events generated from a single system that match a specific event class in the Open Cybersecurity Schema Framework (OCSF) schema.
//
// Amazon Security Lake can collect logs and events from a variety of sources, including natively supported AWS services and third-party custom sources.
//
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
	// The natively supported AWS service which is used a Amazon Security Lake source to collect logs and events from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-source.html#cfn-securitylake-subscriber-source-awslogsource
	//
	AwsLogSource interface{} `field:"optional" json:"awsLogSource" yaml:"awsLogSource"`
	// The custom log source AWS which is used a Amazon Security Lake source to collect logs and events from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-subscriber-source.html#cfn-securitylake-subscriber-source-customlogsource
	//
	CustomLogSource interface{} `field:"optional" json:"customLogSource" yaml:"customLogSource"`
}

