package mixinsawskendra


// Specifies access control list files for the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessControlListConfigurationProperty := &AccessControlListConfigurationProperty{
//   	KeyPath: jsii.String("keyPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-accesscontrollistconfiguration.html
//
type CfnDataSourcePropsMixin_AccessControlListConfigurationProperty struct {
	// Path to the AWS S3 bucket that contains the access control list files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-accesscontrollistconfiguration.html#cfn-kendra-datasource-accesscontrollistconfiguration-keypath
	//
	KeyPath *string `field:"optional" json:"keyPath" yaml:"keyPath"`
}

