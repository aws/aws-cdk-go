package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ndiSourceSettingsProperty := &NdiSourceSettingsProperty{
//   	AudioReferenceLevel: jsii.Number(123),
//   	SourceName: jsii.String("sourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndisourcesettings.html
//
type CfnFlowPropsMixin_NdiSourceSettingsProperty struct {
	// The audio reference level of the NDI sender, in dBFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndisourcesettings.html#cfn-mediaconnect-flow-ndisourcesettings-audioreferencelevel
	//
	AudioReferenceLevel *float64 `field:"optional" json:"audioReferenceLevel" yaml:"audioReferenceLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndisourcesettings.html#cfn-mediaconnect-flow-ndisourcesettings-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

