package previewawsinspectorv2mixins


// The CIS targets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cisTargetsProperty := &CisTargetsProperty{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	TargetResourceTags: map[string][]*string{
//   		"targetResourceTagsKey": []*string{
//   			jsii.String("targetResourceTags"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-cistargets.html
//
type CfnCisScanConfigurationPropsMixin_CisTargetsProperty struct {
	// The CIS target account ids.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-cistargets.html#cfn-inspectorv2-cisscanconfiguration-cistargets-accountids
	//
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// The CIS target resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-cisscanconfiguration-cistargets.html#cfn-inspectorv2-cisscanconfiguration-cistargets-targetresourcetags
	//
	TargetResourceTags interface{} `field:"optional" json:"targetResourceTags" yaml:"targetResourceTags"`
}

