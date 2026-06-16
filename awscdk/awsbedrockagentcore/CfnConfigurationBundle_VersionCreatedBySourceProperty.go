package awsbedrockagentcore


// The source that created a configuration bundle version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versionCreatedBySourceProperty := &VersionCreatedBySourceProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versioncreatedbysource.html
//
type CfnConfigurationBundle_VersionCreatedBySourceProperty struct {
	// The name of the source (for example, user, optimization-job, or system).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versioncreatedbysource.html#cfn-bedrockagentcore-configurationbundle-versioncreatedbysource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the source, if applicable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-configurationbundle-versioncreatedbysource.html#cfn-bedrockagentcore-configurationbundle-versioncreatedbysource-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

