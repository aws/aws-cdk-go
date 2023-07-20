package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostPathProperty := &HostPathProperty{
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-hostpath.html
//
type CfnJobDefinition_HostPathProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-hostpath.html#cfn-batch-jobdefinition-hostpath-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

