package awsapplicationsignals


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   selectionConfigProperty := &SelectionConfigProperty{
//   	Pattern: jsii.String("pattern"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html
//
type CfnServiceLevelObjectivePropsMixin_SelectionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html#cfn-applicationsignals-servicelevelobjective-selectionconfig-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-selectionconfig.html#cfn-applicationsignals-servicelevelobjective-selectionconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

