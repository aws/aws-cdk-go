package mixinsawscustomerprofiles


// A unique key map that can be used to map data to the profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyMapProperty := &KeyMapProperty{
//   	Name: jsii.String("name"),
//   	ObjectTypeKeyList: []interface{}{
//   		&ObjectTypeKeyProperty{
//   			FieldNames: []*string{
//   				jsii.String("fieldNames"),
//   			},
//   			StandardIdentifiers: []*string{
//   				jsii.String("standardIdentifiers"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-keymap.html
//
type CfnObjectTypePropsMixin_KeyMapProperty struct {
	// Name of the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-keymap.html#cfn-customerprofiles-objecttype-keymap-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of ObjectTypeKey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-objecttype-keymap.html#cfn-customerprofiles-objecttype-keymap-objecttypekeylist
	//
	ObjectTypeKeyList interface{} `field:"optional" json:"objectTypeKeyList" yaml:"objectTypeKeyList"`
}

