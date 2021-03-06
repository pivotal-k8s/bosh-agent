// +build !windows

package bundlecollection

import (
	"path"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

func (b FileBundle) Install(sourcePath string) (boshsys.FileSystem, string, error) {
	b.logger.Debug(fileBundleLogTag, "Installing %v", b)

	err := b.fs.Chmod(sourcePath, b.fileMode)
	if err != nil {
		return nil, "", bosherr.WrapError(err, "Setting permissions on source directory")
	}

	err = b.fs.Chown(sourcePath, "root:vcap")
	if err != nil {
		return nil, "", bosherr.WrapError(err, "Setting ownership on source directory")
	}

	err = b.fs.MkdirAll(path.Dir(b.installPath), b.fileMode)
	if err != nil {
		return nil, "", bosherr.WrapError(err, "Creating parent installation directory")
	}

	err = b.fs.Chown(path.Dir(b.installPath), "root:vcap")
	if err != nil {
		return nil, "", bosherr.WrapError(err, "Setting ownership on parent installation directory")
	}

	// Rename MUST be the last possibly-failing operation
	// because IsInstalled() relies on installPath presence.
	err = b.fs.Rename(sourcePath, b.installPath)
	if err != nil {
		return nil, "", bosherr.WrapError(err, "Moving to installation directory")
	}

	return b.fs, b.installPath, nil
}

func (b FileBundle) Uninstall() error {
	b.logger.Debug(fileBundleLogTag, "Uninstalling %v", b)

	// RemoveAll MUST be the last possibly-failing operation
	// because IsInstalled() relies on installPath presence.
	return b.fs.RemoveAll(b.installPath)
}
