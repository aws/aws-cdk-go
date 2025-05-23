package awsnimblestudio


// Initialization scripts for studio components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioComponentInitializationScriptProperty := &StudioComponentInitializationScriptProperty{
//   	LaunchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   	Platform: jsii.String("platform"),
//   	RunContext: jsii.String("runContext"),
//   	Script: jsii.String("script"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentinitializationscript.html
//
type CfnStudioComponent_StudioComponentInitializationScriptProperty struct {
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentinitializationscript.html#cfn-nimblestudio-studiocomponent-studiocomponentinitializationscript-launchprofileprotocolversion
	//
	LaunchProfileProtocolVersion *string `field:"optional" json:"launchProfileProtocolVersion" yaml:"launchProfileProtocolVersion"`
	// The platform of the initialization script, either Windows or Linux.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentinitializationscript.html#cfn-nimblestudio-studiocomponent-studiocomponentinitializationscript-platform
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The method to use when running the initialization script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentinitializationscript.html#cfn-nimblestudio-studiocomponent-studiocomponentinitializationscript-runcontext
	//
	RunContext *string `field:"optional" json:"runContext" yaml:"runContext"`
	// The initialization script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentinitializationscript.html#cfn-nimblestudio-studiocomponent-studiocomponentinitializationscript-script
	//
	Script *string `field:"optional" json:"script" yaml:"script"`
}

