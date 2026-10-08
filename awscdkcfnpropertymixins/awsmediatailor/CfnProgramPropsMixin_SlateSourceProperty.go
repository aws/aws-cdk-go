package awsmediatailor


// Slate VOD source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   slateSourceProperty := &SlateSourceProperty{
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-slatesource.html
//
type CfnProgramPropsMixin_SlateSourceProperty struct {
	// The name of the source location where the slate VOD source is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-slatesource.html#cfn-mediatailor-program-slatesource-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The slate VOD source name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-slatesource.html#cfn-mediatailor-program-slatesource-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

