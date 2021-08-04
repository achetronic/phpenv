package version

type VersionManager interface {
	GetLocalVersions() []string
	GetRemoteVersions() []string
	InstallVersion(v string) (bool, string)
	UninstallVersion(v string) (bool, string)
	UseVersion(v string) (bool, string)
}
