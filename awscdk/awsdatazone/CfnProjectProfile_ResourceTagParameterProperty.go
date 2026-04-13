package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceTagParameterProperty := &ResourceTagParameterProperty{
//   	IsValueEditable: jsii.Boolean(false),
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-resourcetagparameter.html
//
type CfnProjectProfile_ResourceTagParameterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-resourcetagparameter.html#cfn-datazone-projectprofile-resourcetagparameter-isvalueeditable
	//
	IsValueEditable interface{} `field:"required" json:"isValueEditable" yaml:"isValueEditable"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-resourcetagparameter.html#cfn-datazone-projectprofile-resourcetagparameter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-resourcetagparameter.html#cfn-datazone-projectprofile-resourcetagparameter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

