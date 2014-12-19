package devicepathresolver_test

import (
	"errors"
	"os"
	"time"

	boshsettings "github.com/cloudfoundry/bosh-agent/settings"
	fakesys "github.com/cloudfoundry/bosh-agent/system/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-agent/infrastructure/devicepathresolver"
)

var _ = Describe("IDDevicePathResolver", func() {
	var (
		fs           *fakesys.FakeFileSystem
		cmdRunner    *fakesys.FakeCmdRunner
		diskSettings boshsettings.DiskSettings
		pathResolver DevicePathResolver
	)

	BeforeEach(func() {
		fs = fakesys.NewFakeFileSystem()
		cmdRunner = fakesys.NewFakeCmdRunner()
		pathResolver = NewIDDevicePathResolver(time.Second, cmdRunner, fs)
		diskSettings = boshsettings.DiskSettings{
			ID: "fake-disk-id-include-truncate",
		}
	})

	Describe("GetRealDevicePath", func() {
		It("refreshes udev", func() {
			pathResolver.GetRealDevicePath(diskSettings)
			Expect(cmdRunner.RunCommands).To(ContainElement([]string{"udevadm", "trigger"}))
		})

		Context("when path exists", func() {
			BeforeEach(func() {
				err := fs.MkdirAll("fake-device-path", os.FileMode(0750))
				Expect(err).ToNot(HaveOccurred())

				err = fs.Symlink("fake-device-path", "/dev/disk/by-id/virtio-fake-disk-id-include")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns the path ", func() {
				path, timeout, err := pathResolver.GetRealDevicePath(diskSettings)
				Expect(err).ToNot(HaveOccurred())

				Expect(path).To(Equal("fake-device-path"))
				Expect(timeout).To(BeFalse())
			})
		})

		Context("when path does not exist", func() {
			BeforeEach(func() {
				err := fs.Symlink("fake-device-path", "/dev/disk/by-id/virtio-fake-disk-id-include")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns an error", func() {
				_, _, err := pathResolver.GetRealDevicePath(diskSettings)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when symlink does not exist", func() {
			It("returns an error", func() {
				_, _, err := pathResolver.GetRealDevicePath(diskSettings)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when no matching device is found the first time", func() {
			Context("when the timeout has not expired", func() {
				BeforeEach(func() {
					time.AfterFunc(500*time.Millisecond, func() {
						err := fs.MkdirAll("fake-device-path", os.FileMode(0750))
						Expect(err).ToNot(HaveOccurred())

						err = fs.Symlink("fake-device-path", "/dev/disk/by-id/virtio-fake-disk-id-include")
						Expect(err).ToNot(HaveOccurred())
					})
				})

				It("returns the real path", func() {
					path, timeout, err := pathResolver.GetRealDevicePath(diskSettings)
					Expect(err).ToNot(HaveOccurred())

					Expect(path).To(Equal("fake-device-path"))
					Expect(timeout).To(BeFalse())
				})
			})
		})

		Context("when refreshing udev fails", func() {
			BeforeEach(func() {
				cmdRunner.AddCmdResult("udevadm trigger", fakesys.FakeCmdResult{
					Error: errors.New("fake-udevadm-error"),
				})
			})

			It("returns an error", func() {
				_, timeout, err := pathResolver.GetRealDevicePath(diskSettings)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-udevadm-error"))
				Expect(timeout).To(BeFalse())
			})
		})

		Context("when id is empty", func() {
			BeforeEach(func() {
				diskSettings = boshsettings.DiskSettings{}
			})

			It("returns an error", func() {
				_, timeout, err := pathResolver.GetRealDevicePath(diskSettings)
				Expect(err).To(HaveOccurred())
				Expect(timeout).To(BeFalse())
			})
		})
	})
})