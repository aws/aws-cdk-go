package awsdatazone


// Indicates the smithy model of the API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelProperty := &ModelProperty{
//   	Smithy: jsii.String("smithy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-formtype-model.html
//
type CfnFormType_ModelProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-formtype-model.html#cfn-datazone-formtype-model-smithy
	//
	Smithy *string `field:"optional" json:"smithy" yaml:"smithy"`
}

