package awscdk


// Example:
//   type myConstruct struct {
//   	resource
//   	tags
//   }
//
//   func newMyConstruct(scope construct, id *string) *myConstruct {
//   	this := &myConstruct{}
//   	newResource_Override(this, scope, id)
//
//   	awscdk.NewCfnResource(this, jsii.String("Resource"), &cfnResourceProps{
//   		Type: jsii.String("Whatever::The::Type"),
//   		Properties: map[string]interface{}{
//   			// ...
//   			"Tags": this.tags.renderedTags,
//   		},
//   	})
//   	return this
//   }
//
type TagType string

const (
	TagType_STANDARD TagType = "STANDARD"
	TagType_AUTOSCALING_GROUP TagType = "AUTOSCALING_GROUP"
	TagType_MAP TagType = "MAP"
	TagType_KEY_VALUE TagType = "KEY_VALUE"
	TagType_NOT_TAGGABLE TagType = "NOT_TAGGABLE"
)

