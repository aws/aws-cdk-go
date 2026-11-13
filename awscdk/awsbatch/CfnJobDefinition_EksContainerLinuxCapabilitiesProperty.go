package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksContainerLinuxCapabilitiesProperty := &EksContainerLinuxCapabilitiesProperty{
//   	Add: []*string{
//   		jsii.String("add"),
//   	},
//   	Drop: []*string{
//   		jsii.String("drop"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainerlinuxcapabilities.html
//
type CfnJobDefinition_EksContainerLinuxCapabilitiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainerlinuxcapabilities.html#cfn-batch-jobdefinition-ekscontainerlinuxcapabilities-add
	//
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ekscontainerlinuxcapabilities.html#cfn-batch-jobdefinition-ekscontainerlinuxcapabilities-drop
	//
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

