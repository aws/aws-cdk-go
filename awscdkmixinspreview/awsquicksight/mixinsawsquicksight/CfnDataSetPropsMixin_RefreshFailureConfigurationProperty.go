package mixinsawsquicksight


// The failure configuration of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   refreshFailureConfigurationProperty := &RefreshFailureConfigurationProperty{
//   	EmailAlert: &RefreshFailureEmailAlertProperty{
//   		AlertStatus: jsii.String("alertStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureconfiguration.html
//
type CfnDataSetPropsMixin_RefreshFailureConfigurationProperty struct {
	// The email alert configuration for a dataset refresh failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureconfiguration.html#cfn-quicksight-dataset-refreshfailureconfiguration-emailalert
	//
	EmailAlert interface{} `field:"optional" json:"emailAlert" yaml:"emailAlert"`
}

