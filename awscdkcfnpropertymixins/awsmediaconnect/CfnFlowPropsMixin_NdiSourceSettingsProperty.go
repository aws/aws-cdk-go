package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ndiSourceSettingsProperty := &NdiSourceSettingsProperty{
//   	SourceName: jsii.String("sourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndisourcesettings.html
//
type CfnFlowPropsMixin_NdiSourceSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-ndisourcesettings.html#cfn-mediaconnect-flow-ndisourcesettings-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

