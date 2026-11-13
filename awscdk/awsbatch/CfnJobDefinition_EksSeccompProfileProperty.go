package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksSeccompProfileProperty := &EksSeccompProfileProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	LocalhostProfile: jsii.String("localhostProfile"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html
//
type CfnJobDefinition_EksSeccompProfileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html#cfn-batch-jobdefinition-eksseccompprofile-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html#cfn-batch-jobdefinition-eksseccompprofile-localhostprofile
	//
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
}

