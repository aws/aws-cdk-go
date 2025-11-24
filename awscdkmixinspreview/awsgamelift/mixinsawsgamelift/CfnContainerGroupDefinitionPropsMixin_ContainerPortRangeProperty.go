package mixinsawsgamelift


// A set of one or more port numbers that can be opened on the container, and the supported network protocol.
//
// *Part of:* [ContainerPortConfiguration](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerPortConfiguration.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerPortRangeProperty := &ContainerPortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerportrange.html
//
type CfnContainerGroupDefinitionPropsMixin_ContainerPortRangeProperty struct {
	// A starting value for the range of allowed port numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerportrange.html#cfn-gamelift-containergroupdefinition-containerportrange-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The network protocol that these ports support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerportrange.html#cfn-gamelift-containergroupdefinition-containerportrange-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// An ending value for the range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than `FromPort` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerportrange.html#cfn-gamelift-containergroupdefinition-containerportrange-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

