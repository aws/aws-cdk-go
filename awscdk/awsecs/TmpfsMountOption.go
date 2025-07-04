package awsecs


// The supported options for a tmpfs mount for a container.
type TmpfsMountOption string

const (
	TmpfsMountOption_DEFAULTS TmpfsMountOption = "DEFAULTS"
	TmpfsMountOption_RO TmpfsMountOption = "RO"
	TmpfsMountOption_RW TmpfsMountOption = "RW"
	TmpfsMountOption_SUID TmpfsMountOption = "SUID"
	TmpfsMountOption_NOSUID TmpfsMountOption = "NOSUID"
	TmpfsMountOption_DEV TmpfsMountOption = "DEV"
	TmpfsMountOption_NODEV TmpfsMountOption = "NODEV"
	TmpfsMountOption_EXEC TmpfsMountOption = "EXEC"
	TmpfsMountOption_NOEXEC TmpfsMountOption = "NOEXEC"
	TmpfsMountOption_SYNC TmpfsMountOption = "SYNC"
	TmpfsMountOption_ASYNC TmpfsMountOption = "ASYNC"
	TmpfsMountOption_DIRSYNC TmpfsMountOption = "DIRSYNC"
	TmpfsMountOption_REMOUNT TmpfsMountOption = "REMOUNT"
	TmpfsMountOption_MAND TmpfsMountOption = "MAND"
	TmpfsMountOption_NOMAND TmpfsMountOption = "NOMAND"
	TmpfsMountOption_ATIME TmpfsMountOption = "ATIME"
	TmpfsMountOption_NOATIME TmpfsMountOption = "NOATIME"
	TmpfsMountOption_DIRATIME TmpfsMountOption = "DIRATIME"
	TmpfsMountOption_NODIRATIME TmpfsMountOption = "NODIRATIME"
	TmpfsMountOption_BIND TmpfsMountOption = "BIND"
	TmpfsMountOption_RBIND TmpfsMountOption = "RBIND"
	TmpfsMountOption_UNBINDABLE TmpfsMountOption = "UNBINDABLE"
	TmpfsMountOption_RUNBINDABLE TmpfsMountOption = "RUNBINDABLE"
	TmpfsMountOption_PRIVATE TmpfsMountOption = "PRIVATE"
	TmpfsMountOption_RPRIVATE TmpfsMountOption = "RPRIVATE"
	TmpfsMountOption_SHARED TmpfsMountOption = "SHARED"
	TmpfsMountOption_RSHARED TmpfsMountOption = "RSHARED"
	TmpfsMountOption_SLAVE TmpfsMountOption = "SLAVE"
	TmpfsMountOption_RSLAVE TmpfsMountOption = "RSLAVE"
	TmpfsMountOption_RELATIME TmpfsMountOption = "RELATIME"
	TmpfsMountOption_NORELATIME TmpfsMountOption = "NORELATIME"
	TmpfsMountOption_STRICTATIME TmpfsMountOption = "STRICTATIME"
	TmpfsMountOption_NOSTRICTATIME TmpfsMountOption = "NOSTRICTATIME"
	TmpfsMountOption_MODE TmpfsMountOption = "MODE"
	TmpfsMountOption_UID TmpfsMountOption = "UID"
	TmpfsMountOption_GID TmpfsMountOption = "GID"
	TmpfsMountOption_NR_INODES TmpfsMountOption = "NR_INODES"
	TmpfsMountOption_NR_BLOCKS TmpfsMountOption = "NR_BLOCKS"
	TmpfsMountOption_MPOL TmpfsMountOption = "MPOL"
)

