package awsapplicationsignals


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectionConfigProperty := &SelectionConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html
//
type CfnServiceLevelObjective_SelectionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html#cfn-applicationsignals-servicelevelobjective-selectionconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html#cfn-applicationsignals-servicelevelobjective-selectionconfig-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

