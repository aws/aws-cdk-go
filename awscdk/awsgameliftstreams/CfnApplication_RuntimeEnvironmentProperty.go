package awsgameliftstreams


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeEnvironmentProperty := &RuntimeEnvironmentProperty{
//   	Type: jsii.String("type"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html
//
type CfnApplication_RuntimeEnvironmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html#cfn-gameliftstreams-application-runtimeenvironment-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gameliftstreams-application-runtimeenvironment.html#cfn-gameliftstreams-application-runtimeenvironment-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
}

