package interfacesawspcaconnectorad


// A reference to a TemplateGroupAccessControlEntry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateGroupAccessControlEntryReference := &TemplateGroupAccessControlEntryReference{
//   	GroupSecurityIdentifier: jsii.String("groupSecurityIdentifier"),
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
type TemplateGroupAccessControlEntryReference struct {
	// The GroupSecurityIdentifier of the TemplateGroupAccessControlEntry resource.
	GroupSecurityIdentifier *string `field:"required" json:"groupSecurityIdentifier" yaml:"groupSecurityIdentifier"`
	// The TemplateArn of the TemplateGroupAccessControlEntry resource.
	TemplateArn *string `field:"required" json:"templateArn" yaml:"templateArn"`
}

