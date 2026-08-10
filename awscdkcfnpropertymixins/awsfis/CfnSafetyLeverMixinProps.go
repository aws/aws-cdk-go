package awsfis


// Properties for CfnSafetyLeverPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSafetyLeverMixinProps := &CfnSafetyLeverMixinProps{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-safetylever.html
//
type CfnSafetyLeverMixinProps struct {
	// The ID of the safety lever.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-safetylever.html#cfn-fis-safetylever-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

