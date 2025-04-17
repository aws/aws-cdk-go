package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfigurationProperty := &FirelensConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html
//
type CfnJobDefinition_FirelensConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html#cfn-batch-jobdefinition-firelensconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-firelensconfiguration.html#cfn-batch-jobdefinition-firelensconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

