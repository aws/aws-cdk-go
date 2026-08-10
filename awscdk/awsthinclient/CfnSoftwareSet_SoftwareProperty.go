package awsthinclient


// Describes software.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   softwareProperty := &SoftwareProperty{
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-thinclient-softwareset-software.html
//
type CfnSoftwareSet_SoftwareProperty struct {
	// The name of the software component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-thinclient-softwareset-software.html#cfn-thinclient-softwareset-software-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the software component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-thinclient-softwareset-software.html#cfn-thinclient-softwareset-software-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

