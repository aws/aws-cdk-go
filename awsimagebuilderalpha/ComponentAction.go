package awsimagebuilderalpha


// The action for a step within the component document.
//
// Example:
//   component := imagebuilder.NewComponent(this, jsii.String("MyComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
//   		"phases": []interface{}{
//   			map[string]interface{}{
//   				"name": imagebuilder.ComponentPhaseName_BUILD,
//   				"steps": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("install-app"),
//   						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
//   						"inputs": map[string][]*string{
//   							"commands": []*string{
//   								jsii.String("echo \"Installing my application...\""),
//   								jsii.String("yum update -y"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type ComponentAction string

const (
	// The AppendFile action adds the provided content to the pre-existing content of a file.
	// Experimental.
	ComponentAction_APPEND_FILE ComponentAction = "APPEND_FILE"
	// The Assert action performs value with comparison/logic operators, and succeeds/fails the step based on the outcome of the assertion.
	// Experimental.
	ComponentAction_ASSERT ComponentAction = "ASSERT"
	// The CopyFile action copies files from a source file to a destination.
	// Experimental.
	ComponentAction_COPY_FILE ComponentAction = "COPY_FILE"
	// The CopyFolder action copies folders from a source to a destination.
	// Experimental.
	ComponentAction_COPY_FOLDER ComponentAction = "COPY_FOLDER"
	// The CreateFile action creates a file in the provided location.
	// Experimental.
	ComponentAction_CREATE_FILE ComponentAction = "CREATE_FILE"
	// The CreateFolder action creates a folder in the provided location.
	// Experimental.
	ComponentAction_CREATE_FOLDER ComponentAction = "CREATE_FOLDER"
	// The CreateSymlink action creates symbolic links from a given path to a target.
	// Experimental.
	ComponentAction_CREATE_SYMLINK ComponentAction = "CREATE_SYMLINK"
	// The DeleteFile action deletes file(s) in the provided location.
	// Experimental.
	ComponentAction_DELETE_FILE ComponentAction = "DELETE_FILE"
	// The DeleteFolder action deletes folders in the provided location.
	// Experimental.
	ComponentAction_DELETE_FOLDER ComponentAction = "DELETE_FOLDER"
	// The ExecuteBash action runs bash scripts with inline bash commands.
	// Experimental.
	ComponentAction_EXECUTE_BASH ComponentAction = "EXECUTE_BASH"
	// The ExecuteBinary action runs a provided binary executable.
	// Experimental.
	ComponentAction_EXECUTE_BINARY ComponentAction = "EXECUTE_BINARY"
	// The ExecuteDocument action allows running other component documents from within a given component.
	// Experimental.
	ComponentAction_EXECUTE_DOCUMENT ComponentAction = "EXECUTE_DOCUMENT"
	// The ExecutePowershell action runs PowerShell scripts with inline PowerShell commands.
	// Experimental.
	ComponentAction_EXECUTE_POWERSHELL ComponentAction = "EXECUTE_POWERSHELL"
	// The InstallMSI action runs a Windows application with the provided MSI file.
	// Experimental.
	ComponentAction_INSTALL_MSI ComponentAction = "INSTALL_MSI"
	// The ListFiles action lists files in the provided folder.
	// Experimental.
	ComponentAction_LIST_FILES ComponentAction = "LIST_FILES"
	// The MoveFile action moves files from a source to a destination.
	// Experimental.
	ComponentAction_MOVE_FILE ComponentAction = "MOVE_FILE"
	// The MoveFolder action moves folders from a source to a destination.
	// Experimental.
	ComponentAction_MOVE_FOLDER ComponentAction = "MOVE_FOLDER"
	// The ReadFile action reads the content of a text file.
	// Experimental.
	ComponentAction_READ_FILE ComponentAction = "READ_FILE"
	// The Reboot action reboots the instance.
	// Experimental.
	ComponentAction_REBOOT ComponentAction = "REBOOT"
	// The SetFileEncoding action modifies the encoding property of an existing file.
	// Experimental.
	ComponentAction_SET_FILE_ENCODING ComponentAction = "SET_FILE_ENCODING"
	// The SetFileOwner action modifies the owner and group ownership properties of an existing file.
	// Experimental.
	ComponentAction_SET_FILE_OWNER ComponentAction = "SET_FILE_OWNER"
	// The SetFolderOwner action recursively modifies the owner and group ownership properties of an existing folder.
	// Experimental.
	ComponentAction_SET_FOLDER_OWNER ComponentAction = "SET_FOLDER_OWNER"
	// The SetFilePermissions action modifies the permission of an existing file.
	// Experimental.
	ComponentAction_SET_FILE_PERMISSIONS ComponentAction = "SET_FILE_PERMISSIONS"
	// The SetFolderPermissions action recursively modifies the permissions of an existing folder and all of its subfiles and subfolders.
	// Experimental.
	ComponentAction_SET_FOLDER_PERMISSIONS ComponentAction = "SET_FOLDER_PERMISSIONS"
	// The SetRegistry action sets the value for the specified Windows registry key.
	// Experimental.
	ComponentAction_SET_REGISTRY ComponentAction = "SET_REGISTRY"
	// The S3Download action downloads an Amazon S3 object/folder to a local destination.
	// Experimental.
	ComponentAction_S3_DOWNLOAD ComponentAction = "S3_DOWNLOAD"
	// The S3Upload action uploads a file or folder to an Amazon S3 location.
	// Experimental.
	ComponentAction_S3_UPLOAD ComponentAction = "S3_UPLOAD"
	// The UninstallMSI action removes a Windows application using an MSI file.
	// Experimental.
	ComponentAction_UNINSTALL_MSI ComponentAction = "UNINSTALL_MSI"
	// The UpdateOS action installs OS updates.
	// Experimental.
	ComponentAction_UPDATE_OS ComponentAction = "UPDATE_OS"
	// The WebDownload action downloads files from a URL to a local destination.
	// Experimental.
	ComponentAction_WEB_DOWNLOAD ComponentAction = "WEB_DOWNLOAD"
)

