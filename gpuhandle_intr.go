package nvml

import "time"

type GpuHandleInterface interface {
	DeviceClearCpuAffinity() error
	DeviceGetAPIRestriction(apiType RestrictedAPI) (bool, error)
	DeviceGetApplicationsClock(clockType ClockType) (uint, error)
	DeviceGetAutoBoostedClocksEnabled() (curState bool, defaultState bool, err error)
	DeviceGetBAR1MemoryInfo() (free uint64, used uint64, total uint64, err error)
	DeviceGetBoardId() (uint, error)
	DeviceGetBrand() (BrandType, error)
	DeviceGetBridgeChipInfo() ([]*BridgeChipInfo, error)
	DeviceGetClockInfo(clockType ClockType) (uint, error)
	DeviceGetComputeMode() (ComputeMode, error)
	DeviceGetComputeRunningProcesses(size int) ([]*ProcessInfo, error)
	DeviceGetCpuAffinity(size uint) ([]uint, error)
	DeviceGetCurrPcieLinkGeneration() (uint, error)
	DeviceGetCurrPcieLinkWidth() (uint, error)
	DeviceGetCurrentClocksThrottleReasons() ([]ClocksThrottleReasons, error)
	DeviceGetDecoderUtilization() (util uint, samplePeriod uint, err error)
	DeviceGetDefaultApplicationsClock(clockType ClockType) (uint, error)
	DeviceGetDetailedEccErrors(mt MemoryErrorType, ec EccCounterType) (*EccErrorCounts, error)
	DeviceGetDisplayActive() (bool, error)
	DeviceGetDisplayMode() (bool, error)
	DeviceGetEccMode() (curMode bool, pendingMode bool, err error)
	DeviceGetEncoderUtilization() (util uint, samplePeriod uint, err error)
	DeviceGetEnforcedPowerLimit() (uint, error)
	DeviceGetFanSpeed() (uint, error)
	DeviceGetGpuOperationMode() (curMode GpuOperationMode, pendingMode GpuOperationMode, err error)
	GetGraphicsRunningProcesses(size int) ([]*ProcessInfo, error)
	DeviceGetIndex() (uint, error)
	DeviceGetInforomConfigurationChecksum() (uint, error)
	DeviceGetInforomImageVersion() (string, error)
	DeviceGetInforomVersion(object InforomObject) (string, error)
	DeviceGetMaxClockInfo(clockType ClockType) (uint, error)
	DeviceGetMaxPcieLinkGeneration() (uint, error)
	DeviceGetMaxPcieLinkWidth() (uint, error)
	DeviceGetMemoryErrorCounter(mt MemoryErrorType, ec EccCounterType, loc MemoryLocation) (uint64, error)
	DeviceGetMemoryInfo() (free uint64, used uint64, total uint64, err error)
	DeviceGetMinorNumber() (uint, error)
	DeviceGetMultiGpuBoard() (uint, error)
	DeviceGetName() (string, error)
	DeviceGetPciInfo() (*PciInfo, error)
	DeviceGetPcieReplayCounter() (uint, error)
	DeviceGetPcieThroughput(counterType PcieUtilCounter) (uint, error)
	DeviceGetPerformanceState() (uint, error)
	DeviceGetPersistenceMode() (bool, error)
	DeviceGetPowerManagementDefaultLimit() (uint, error)
	DeviceGetPowerManagementLimit() (uint, error)
	DeviceGetPowerManagementLimitConstraints() (min uint, max uint, err error)
	DeviceGetPowerManagementMode() (bool, error)
	DeviceGetPowerState() (uint, error)
	DeviceGetPowerUsage() (uint, error)
	DeviceGetRetiredPages(cause PageRetirementCause) ([]uint64, error)
	DeviceGetRetiredPagesPendingStatus() (bool, error)
	DeviceGetSerial() (string, error)
	DeviceGetSupportedClocksThrottleReasons() (uint64, error)
	DeviceGetSupportedGraphicsClocks(memoryClockMHz uint) ([]uint, error)
	DeviceGetSupportedMemoryClocks() ([]uint, error)
	DeviceGetTemperature() (uint, error)
	DeviceGetTemperatureThreshold(threshold TemperatureThresholds) (uint, error)
	DeviceGetTopologyNearestGpus(level GpuTopologyLevel) ([]GpuHandle, error)
	DeviceGetTotalEccErrors(mt MemoryErrorType, et EccCounterType) (uint64, error)
	DeviceGetUUID() (string, error)
	DeviceGetUtilizationRates() (*Utilization, error)
	DeviceGetVbiosVersion() (string, error)
	DeviceGetViolationStatus(policy PerfPolicy) (*ViolationTime, error)
	DeviceResetApplicationsClocks() error
	DeviceSetAutoBoostedClocksEnabled() (bool, error)
	DeviceSetCpuAffinity() error
	DeviceSetDefaultAutoBoostedClocksEnabled(enabled bool) error
	DeviceValidateInforom() error
	DeviceClearEccErrorCounts(counterType EccCounterType) error
	DeviceSetAPIRestriction(api RestrictedAPI, isEnabled bool) error
	DeviceSetApplicationsClocks(memClocksMHz, clocksMHz uint) error
	DeviceSetComputeMode(mode ComputeMode) error
	DeviceSetEccMode(isEnabled bool) error
	DeviceSetGpuOperationMode(mode GpuOperationMode) error
	DeviceSetPersistenceMode(isEnabled bool) error
	DeviceSetPowerManagementLimit(limit uint) error
	DeviceGetAverageGPUUsage(since time.Duration) (uint, error)
	DeviceGetProcessUtilization(maxProcess int, since time.Duration) ([]*ProcessUtilizationSample, error)
	DeviceGetSupportedEventTypes() ([]EventType, error)
	DeviceRegisterEvents(evtType EventType, set EventSet) error
}
