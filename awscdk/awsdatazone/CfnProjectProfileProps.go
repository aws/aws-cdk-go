package awsdatazone


// Properties for defining a `CfnProjectProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProfileProps := &CfnProjectProfileProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AllowedDesignations: []interface{}{
//   		&DesignationConfigurationProperty{
//   			DesignationId: jsii.String("designationId"),
//   		},
//   	},
//   	ChangeLog: jsii.String("changeLog"),
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	DomainUnitIdentifier: jsii.String("domainUnitIdentifier"),
//   	ProjectScopes: []interface{}{
//   		&ProjectScopeProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Policy: jsii.String("policy"),
//   		},
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html
//
type CfnProjectProfileProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-alloweddesignations
	//
	AllowedDesignations interface{} `field:"optional" json:"allowedDesignations" yaml:"allowedDesignations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-changelog
	//
	ChangeLog *string `field:"optional" json:"changeLog" yaml:"changeLog"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-domainunitidentifier
	//
	DomainUnitIdentifier *string `field:"optional" json:"domainUnitIdentifier" yaml:"domainUnitIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-projectscopes
	//
	ProjectScopes interface{} `field:"optional" json:"projectScopes" yaml:"projectScopes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-projectprofile.html#cfn-datazone-projectprofile-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

