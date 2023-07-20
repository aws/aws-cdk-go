package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emptyDirProperty := &EmptyDirProperty{
//   	Medium: jsii.String("medium"),
//   	SizeLimit: jsii.String("sizeLimit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-emptydir.html
//
type CfnJobDefinition_EmptyDirProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-emptydir.html#cfn-batch-jobdefinition-emptydir-medium
	//
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-emptydir.html#cfn-batch-jobdefinition-emptydir-sizelimit
	//
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

