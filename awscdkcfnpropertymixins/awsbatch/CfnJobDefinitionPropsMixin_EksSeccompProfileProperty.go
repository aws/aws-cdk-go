package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eksSeccompProfileProperty := &EksSeccompProfileProperty{
//   	LocalhostProfile: jsii.String("localhostProfile"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html
//
type CfnJobDefinitionPropsMixin_EksSeccompProfileProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html#cfn-batch-jobdefinition-eksseccompprofile-localhostprofile
	//
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-eksseccompprofile.html#cfn-batch-jobdefinition-eksseccompprofile-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

