package models

import (
	"errors"

	"code.cloudfoundry.org/bbs/format"
)

func (*VolumePlacement) Version() format.Version {
	return format.V1
}

func (*VolumePlacement) Validate() error {
	return nil
}

// while volume mounts are experimental, we should never persist a "old" volume
// mount to the db layer, so the handler must convert old data models to the new ones
// when volume mounts are no longer experimental, this validation strategy must be reconsidered
func (v *VolumeMount) Validate() error {
	var ve ValidationError
	if v.Driver == "" {
		ve = ve.Append(errors.New("invalid volume_mount driver"))
	}
	if !(v.Mode == "r" || v.Mode == "rw") {
		ve = ve.Append(errors.New("invalid volume_mount mode"))
	}
	if v.Shared != nil && v.Shared.VolumeId == "" {
		ve = ve.Append(errors.New("invalid volume_mount volume id"))
	}

	if !ve.Empty() {
		return ve
	}

	return nil
}
