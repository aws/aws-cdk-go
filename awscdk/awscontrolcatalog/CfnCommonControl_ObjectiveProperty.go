package awscontrolcatalog


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectiveProperty := &ObjectiveProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-commoncontrol-objective.html
//
type CfnCommonControl_ObjectiveProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-commoncontrol-objective.html#cfn-controlcatalog-commoncontrol-objective-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-commoncontrol-objective.html#cfn-controlcatalog-commoncontrol-objective-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

