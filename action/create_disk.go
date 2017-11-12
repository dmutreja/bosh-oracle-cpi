package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"time"
)

const oneGiB int64 = 1024
const minVolumeSize = 50 * oneGiB

// CreateDisk action handles the create_disk method invocation
type CreateDisk struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewCreateDisk creates a CreateDisk instance
func NewCreateDisk(c client.Connector, l boshlog.Logger) CreateDisk {
	return CreateDisk{connector: c, logger: l}
}

// Run creates a persistent block volume of the requested size
// and returns it's ID
func (cd CreateDisk) Run(size int, _ DiskCloudProperties, vmCID VMCID) (DiskCID, error) {

	in, err := newVMFinder(cd.connector, cd.logger).FindInstance(string(vmCID))

	if err != nil {
		return "", bosherr.WrapError(err, "Unable to find VM")
	}

	creator := newDiskCreator(cd.connector, cd.logger,
		resource.NewLocation(in.Location().AvailabilityDomain(),
			cd.connector.CompartmentId()))

	volName := fmt.Sprintf("bosh-volume-%s", time.Now().Format(time.Stamp))
	vol, err := creator.CreateVolume(volName, volumeSize(size))
	if err != nil {
		return "", bosherr.WrapError(err, "Error creating volume")
	}
	return DiskCID(vol.ID()), nil
}

func volumeSize(size int) int64 {
	s := int64(size)
	if s < minVolumeSize {
		return minVolumeSize
	}
	m := s % oneGiB
	if m != 0 {
		s += oneGiB - m
	}
	return s
}
