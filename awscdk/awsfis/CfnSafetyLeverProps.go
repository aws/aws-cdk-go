package awsfis


// Properties for defining a `CfnSafetyLever`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSafetyLeverProps := &CfnSafetyLeverProps{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-safetylever.html
//
type CfnSafetyLeverProps struct {
	// The ID of the safety lever.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-safetylever.html#cfn-fis-safetylever-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

