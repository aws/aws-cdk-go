package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   groupDefinitionProperty := &GroupDefinitionProperty{
//   	Key: jsii.String("key"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-groupdefinition.html
//
type CfnDashboardPropsMixin_GroupDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-groupdefinition.html#cfn-bcm-dashboard-groupdefinition-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-groupdefinition.html#cfn-bcm-dashboard-groupdefinition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

