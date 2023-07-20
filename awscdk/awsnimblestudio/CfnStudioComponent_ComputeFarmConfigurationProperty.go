package awsnimblestudio


// The configuration for a render farm that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeFarmConfigurationProperty := &ComputeFarmConfigurationProperty{
//   	ActiveDirectoryUser: jsii.String("activeDirectoryUser"),
//   	Endpoint: jsii.String("endpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-computefarmconfiguration.html
//
type CfnStudioComponent_ComputeFarmConfigurationProperty struct {
	// The name of an Active Directory user that is used on ComputeFarm worker instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-computefarmconfiguration.html#cfn-nimblestudio-studiocomponent-computefarmconfiguration-activedirectoryuser
	//
	ActiveDirectoryUser *string `field:"optional" json:"activeDirectoryUser" yaml:"activeDirectoryUser"`
	// The endpoint of the ComputeFarm that is accessed by the studio component resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-computefarmconfiguration.html#cfn-nimblestudio-studiocomponent-computefarmconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

