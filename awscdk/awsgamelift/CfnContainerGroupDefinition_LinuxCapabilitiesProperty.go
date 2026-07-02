package awsgamelift


// A set of Linux capabilities that are added to a container's default Docker configuration.
//
// For more detailed information, see the capabilities(7) Linux manual page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxCapabilitiesProperty := &LinuxCapabilitiesProperty{
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-linuxcapabilities.html
//
type CfnContainerGroupDefinition_LinuxCapabilitiesProperty struct {
	// The list of Linux capabilities to add to the container's default configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-linuxcapabilities.html#cfn-gamelift-containergroupdefinition-linuxcapabilities-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

