package previewawskinesisanalyticsv2mixins


// The information required to specify a Maven reference.
//
// You can use Maven references to specify dependency JAR files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mavenReferenceProperty := &MavenReferenceProperty{
//   	ArtifactId: jsii.String("artifactId"),
//   	GroupId: jsii.String("groupId"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mavenreference.html
//
type CfnApplicationPropsMixin_MavenReferenceProperty struct {
	// The artifact ID of the Maven reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mavenreference.html#cfn-kinesisanalyticsv2-application-mavenreference-artifactid
	//
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// The group ID of the Maven reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mavenreference.html#cfn-kinesisanalyticsv2-application-mavenreference-groupid
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The version of the Maven reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-mavenreference.html#cfn-kinesisanalyticsv2-application-mavenreference-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

