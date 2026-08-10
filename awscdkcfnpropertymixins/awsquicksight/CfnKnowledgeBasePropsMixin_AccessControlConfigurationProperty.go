package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   accessControlConfigurationProperty := &AccessControlConfigurationProperty{
//   	IsAclEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-accesscontrolconfiguration.html
//
type CfnKnowledgeBasePropsMixin_AccessControlConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-knowledgebase-accesscontrolconfiguration.html#cfn-quicksight-knowledgebase-accesscontrolconfiguration-isaclenabled
	//
	IsAclEnabled interface{} `field:"optional" json:"isAclEnabled" yaml:"isAclEnabled"`
}

