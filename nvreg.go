package main

import (
	"errors"
	"golang.org/x/sys/windows/registry"
	"strings"
)

const nvModesName = "NV_Modes"

func addResolution(res string) (err error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SYSTEM\\ControlSet001\\Control\\Class\\{4d36e968-e325-11ce-bfc1-08002be10318}\\0000", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, key.Close())
	}()

	modeGroups, _, err := key.GetStringsValue(nvModesName)
	if err != nil {
		return err
	}
	err = updateModeGroupsAddStandardResolution(modeGroups, res)
	if err != nil {
		return err
	}

	// Write back
	err = key.SetStringsValue(nvModesName, modeGroups)
	return err
}

// updateModeGroupsAddStandardResolution updates the supplied mode groups, adding a given resolution to the standard ones
func updateModeGroupsAddStandardResolution(modeGroups []string, res string) error {
	for i, mode := range modeGroups {
		// Docs, not for NV_Modes specifically but terminology seems to be the same:
		// https://download.nvidia.com/Windows/43.45/NV_Compress_Modes_Users_Guide_2.1.pdf
		// {*} 	= all PCI IDs. Seems to be optional in NV_Modes
		// S	= Standard mode
		// H 	= Horizontal spanning mode
		// V	= Vertical spanning mode
		if strings.HasPrefix(mode, "S ") {
			modeGroups[i] = addStandardResolution(mode, res)
			return nil
		}
	}
	return errors.New("no standard mode group found to add resolution to")
}

func addStandardResolution(mode string, res string) string {
	// Idempotency - do not re-add the res if already present
	if strings.Contains(mode, res) {
		return mode
	}
	return mode + " " + res
}
