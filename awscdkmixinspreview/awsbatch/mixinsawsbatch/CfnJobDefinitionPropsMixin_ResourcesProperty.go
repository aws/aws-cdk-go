package mixinsawsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var limits interface{}
//   var requests interface{}
//
//   resourcesProperty := &ResourcesProperty{
//   	Limits: limits,
//   	Requests: requests,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resources.html
//
type CfnJobDefinitionPropsMixin_ResourcesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resources.html#cfn-batch-jobdefinition-resources-limits
	//
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-resources.html#cfn-batch-jobdefinition-resources-requests
	//
	Requests interface{} `field:"optional" json:"requests" yaml:"requests"`
}

