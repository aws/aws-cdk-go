package mixinsawskinesisanalyticsv2


// Property key-value pairs passed into an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   propertyGroupProperty := &PropertyGroupProperty{
//   	PropertyGroupId: jsii.String("propertyGroupId"),
//   	PropertyMap: map[string]*string{
//   		"propertyMapKey": jsii.String("propertyMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html
//
type CfnApplicationPropsMixin_PropertyGroupProperty struct {
	// Describes the key of an application execution property key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertygroupid
	//
	PropertyGroupId *string `field:"optional" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Describes the value of an application execution property key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertymap
	//
	PropertyMap interface{} `field:"optional" json:"propertyMap" yaml:"propertyMap"`
}

