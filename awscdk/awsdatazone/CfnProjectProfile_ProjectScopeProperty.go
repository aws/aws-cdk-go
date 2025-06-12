package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectScopeProperty := &ProjectScopeProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Policy: jsii.String("policy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-projectscope.html
//
type CfnProjectProfile_ProjectScopeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-projectscope.html#cfn-datazone-projectprofile-projectscope-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-projectscope.html#cfn-datazone-projectprofile-projectscope-policy
	//
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

