package previewawscodebuildevents


// Type definition for Project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   project := &Project{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	Artifacts: &Artifacts1{
//   		EncryptionDisabled: []*string{
//   			jsii.String("encryptionDisabled"),
//   		},
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		NamespaceType: []*string{
//   			jsii.String("namespaceType"),
//   		},
//   		Packaging: []*string{
//   			jsii.String("packaging"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Badge: &Badge{
//   		BadgeEnabled: []*string{
//   			jsii.String("badgeEnabled"),
//   		},
//   		BadgeRequestUrl: []*string{
//   			jsii.String("badgeRequestUrl"),
//   		},
//   	},
//   	Cache: &Cache{
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	Created: []*string{
//   		jsii.String("created"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	EncryptionKey: []*string{
//   		jsii.String("encryptionKey"),
//   	},
//   	Environment: &Environment1{
//   		ComputeType: []*string{
//   			jsii.String("computeType"),
//   		},
//   		EnvironmentVariables: []EnvironmentItem{
//   			&EnvironmentItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Type: []*string{
//   					jsii.String("type"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		Image: []*string{
//   			jsii.String("image"),
//   		},
//   		ImagePullCredentialsType: []*string{
//   			jsii.String("imagePullCredentialsType"),
//   		},
//   		PrivilegedMode: []*string{
//   			jsii.String("privilegedMode"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	LastModified: []*string{
//   		jsii.String("lastModified"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	QueuedTimeoutInMinutes: []*string{
//   		jsii.String("queuedTimeoutInMinutes"),
//   	},
//   	ServiceRole: []*string{
//   		jsii.String("serviceRole"),
//   	},
//   	Source: &Source1{
//   		Auth: &Auth{
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   		Buildspec: []*string{
//   			jsii.String("buildspec"),
//   		},
//   		GitCloneDepth: []*string{
//   			jsii.String("gitCloneDepth"),
//   		},
//   		GitSubmodulesConfig: &GitSubmodulesConfig{
//   			FetchSubmodules: []*string{
//   				jsii.String("fetchSubmodules"),
//   			},
//   		},
//   		InsecureSsl: []*string{
//   			jsii.String("insecureSsl"),
//   		},
//   		Location: []*string{
//   			jsii.String("location"),
//   		},
//   		ReportBuildStatus: []*string{
//   			jsii.String("reportBuildStatus"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   	SourceVersion: []*string{
//   		jsii.String("sourceVersion"),
//   	},
//   	Tags: []ProjectItem{
//   		&ProjectItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	TimeoutInMinutes: []*string{
//   		jsii.String("timeoutInMinutes"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Project struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// artifacts property.
	//
	// Specify an array of string values to match this event if the actual value of artifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Artifacts *AWSAPICallViaCloudTrail_Artifacts1 `field:"optional" json:"artifacts" yaml:"artifacts"`
	// badge property.
	//
	// Specify an array of string values to match this event if the actual value of badge is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Badge *AWSAPICallViaCloudTrail_Badge `field:"optional" json:"badge" yaml:"badge"`
	// cache property.
	//
	// Specify an array of string values to match this event if the actual value of cache is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cache *AWSAPICallViaCloudTrail_Cache `field:"optional" json:"cache" yaml:"cache"`
	// created property.
	//
	// Specify an array of string values to match this event if the actual value of created is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Created *[]*string `field:"optional" json:"created" yaml:"created"`
	// description property.
	//
	// Specify an array of string values to match this event if the actual value of description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// encryptionKey property.
	//
	// Specify an array of string values to match this event if the actual value of encryptionKey is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EncryptionKey *[]*string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *AWSAPICallViaCloudTrail_Environment1 `field:"optional" json:"environment" yaml:"environment"`
	// lastModified property.
	//
	// Specify an array of string values to match this event if the actual value of lastModified is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModified *[]*string `field:"optional" json:"lastModified" yaml:"lastModified"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// queuedTimeoutInMinutes property.
	//
	// Specify an array of string values to match this event if the actual value of queuedTimeoutInMinutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueuedTimeoutInMinutes *[]*string `field:"optional" json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// serviceRole property.
	//
	// Specify an array of string values to match this event if the actual value of serviceRole is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceRole *[]*string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// source property.
	//
	// Specify an array of string values to match this event if the actual value of source is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Source *AWSAPICallViaCloudTrail_Source1 `field:"optional" json:"source" yaml:"source"`
	// sourceVersion property.
	//
	// Specify an array of string values to match this event if the actual value of sourceVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceVersion *[]*string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*AWSAPICallViaCloudTrail_ProjectItem `field:"optional" json:"tags" yaml:"tags"`
	// timeoutInMinutes property.
	//
	// Specify an array of string values to match this event if the actual value of timeoutInMinutes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TimeoutInMinutes *[]*string `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

