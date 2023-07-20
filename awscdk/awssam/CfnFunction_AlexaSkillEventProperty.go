package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alexaSkillEventProperty := &AlexaSkillEventProperty{
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-alexaskillevent.html
//
type CfnFunction_AlexaSkillEventProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-alexaskillevent.html#cfn-serverless-function-alexaskillevent-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

