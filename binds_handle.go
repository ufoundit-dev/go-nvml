//go:build !linux

package nvml

import "time"

type NewHandleFunction func(h GpuHandle) GpuHandleInterface

var (
	_newFunction NewHandleFunction
)

func SewHandleFunction(f NewHandleFunction) {
	_newFunction = f
}

func (h GpuHandle) DeviceClearCpuAffinity() error {
	return _newFunction(h).DeviceClearCpuAffinity()
}

func (h GpuHandle) DeviceGetAPIRestriction(apiType RestrictedAPI) (bool, error) {
	return _newFunction(h).DeviceGetAPIRestriction(apiType)
}
func (h GpuHandle) DeviceGetApplicationsClock(clockType ClockType) (uint, error) {
	return _newFunction(h).DeviceGetApplicationsClock(clockType)
}
func (h GpuHandle) DeviceGetAutoBoostedClocksEnabled() (curState bool, defaultState bool, err error) {
	return _newFunction(h).DeviceGetAutoBoostedClocksEnabled()
}
func (h GpuHandle) DeviceGetBAR1MemoryInfo() (free uint64, used uint64, total uint64, err error) {
	return _newFunction(h).DeviceGetBAR1MemoryInfo()
}
func (h GpuHandle) DeviceGetBoardId() (uint, error) {
	return _newFunction(h).DeviceGetBoardId()
}
func (h GpuHandle) DeviceGetBrand() (BrandType, error) {
	return _newFunction(h).DeviceGetBrand()
}
func (h GpuHandle) DeviceGetBridgeChipInfo() ([]*BridgeChipInfo, error) {
	return _newFunction(h).DeviceGetBridgeChipInfo()
}
func (h GpuHandle) DeviceGetClockInfo(clockType ClockType) (uint, error) {
	return _newFunction(h).DeviceGetClockInfo(clockType)
}
func (h GpuHandle) DeviceGetComputeMode() (ComputeMode, error) {
	return _newFunction(h).DeviceGetComputeMode()
}
func (h GpuHandle) DeviceGetComputeRunningProcesses(size int) ([]*ProcessInfo, error) {
	return _newFunction(h).DeviceGetComputeRunningProcesses(size)
}
func (h GpuHandle) DeviceGetCpuAffinity(size uint) ([]uint, error) {
	return _newFunction(h).DeviceGetCpuAffinity(size)
}
func (h GpuHandle) DeviceGetCurrPcieLinkGeneration() (uint, error) {
	return _newFunction(h).DeviceGetCurrPcieLinkGeneration()
}
func (h GpuHandle) DeviceGetCurrPcieLinkWidth() (uint, error) {
	return _newFunction(h).DeviceGetCurrPcieLinkWidth()
}
func (h GpuHandle) DeviceGetCurrentClocksThrottleReasons() ([]ClocksThrottleReasons, error) {
	return _newFunction(h).DeviceGetCurrentClocksThrottleReasons()
}
func (h GpuHandle) DeviceGetDecoderUtilization() (util uint, samplePeriod uint, err error) {
	return _newFunction(h).DeviceGetDecoderUtilization()
}
func (h GpuHandle) DeviceGetDefaultApplicationsClock(clockType ClockType) (uint, error) {
	return _newFunction(h).DeviceGetDefaultApplicationsClock(clockType)
}
func (h GpuHandle) DeviceGetDetailedEccErrors(mt MemoryErrorType, ec EccCounterType) (*EccErrorCounts, error) {
	return _newFunction(h).DeviceGetDetailedEccErrors(mt, ec)
}
func (h GpuHandle) DeviceGetDisplayActive() (bool, error) {
	return _newFunction(h).DeviceGetDisplayActive()
}
func (h GpuHandle) DeviceGetDisplayMode() (bool, error) {
	return _newFunction(h).DeviceGetDisplayMode()
}
func (h GpuHandle) DeviceGetEccMode() (curMode bool, pendingMode bool, err error) {
	return _newFunction(h).DeviceGetEccMode()
}
func (h GpuHandle) DeviceGetEncoderUtilization() (util uint, samplePeriod uint, err error) {
	return _newFunction(h).DeviceGetEncoderUtilization()
}
func (h GpuHandle) DeviceGetEnforcedPowerLimit() (uint, error) {
	return _newFunction(h).DeviceGetEnforcedPowerLimit()
}
func (h GpuHandle) DeviceGetFanSpeed() (uint, error) {
	return _newFunction(h).DeviceGetFanSpeed()
}
func (h GpuHandle) DeviceGetGpuOperationMode() (curMode GpuOperationMode, pendingMode GpuOperationMode, err error) {
	return _newFunction(h).DeviceGetGpuOperationMode()
}
func (h GpuHandle) GetGraphicsRunningProcesses(size int) ([]*ProcessInfo, error) {
	return _newFunction(h).GetGraphicsRunningProcesses(size)
}
func (h GpuHandle) DeviceGetIndex() (uint, error) {
	return _newFunction(h).DeviceGetIndex()
}
func (h GpuHandle) DeviceGetInforomConfigurationChecksum() (uint, error) {
	return _newFunction(h).DeviceGetInforomConfigurationChecksum()
}
func (h GpuHandle) DeviceGetInforomImageVersion() (string, error) {
	return _newFunction(h).DeviceGetInforomImageVersion()
}
func (h GpuHandle) DeviceGetInforomVersion(object InforomObject) (string, error) {
	return _newFunction(h).DeviceGetInforomVersion(object)
}
func (h GpuHandle) DeviceGetMaxClockInfo(clockType ClockType) (uint, error) {
	return _newFunction(h).DeviceGetMaxClockInfo(clockType)
}
func (h GpuHandle) DeviceGetMaxPcieLinkGeneration() (uint, error) {
	return _newFunction(h).DeviceGetMaxPcieLinkGeneration()
}
func (h GpuHandle) DeviceGetMaxPcieLinkWidth() (uint, error) {
	return _newFunction(h).DeviceGetMaxPcieLinkWidth()
}
func (h GpuHandle) DeviceGetMemoryErrorCounter(mt MemoryErrorType, ec EccCounterType, loc MemoryLocation) (uint64, error) {
	return _newFunction(h).DeviceGetMemoryErrorCounter(mt, ec, loc)
}
func (h GpuHandle) DeviceGetMemoryInfo() (free uint64, used uint64, total uint64, err error) {
	return _newFunction(h).DeviceGetMemoryInfo()
}
func (h GpuHandle) DeviceGetMinorNumber() (uint, error) {
	return _newFunction(h).DeviceGetMinorNumber()
}
func (h GpuHandle) DeviceGetMultiGpuBoard() (uint, error) {
	return _newFunction(h).DeviceGetMultiGpuBoard()
}
func (h GpuHandle) DeviceGetName() (string, error) {
	return _newFunction(h).DeviceGetName()
}
func (h GpuHandle) DeviceGetPciInfo() (*PciInfo, error) {
	return _newFunction(h).DeviceGetPciInfo()
}
func (h GpuHandle) DeviceGetPcieReplayCounter() (uint, error) {
	return _newFunction(h).DeviceGetPcieReplayCounter()
}
func (h GpuHandle) DeviceGetPcieThroughput(counterType PcieUtilCounter) (uint, error) {
	return _newFunction(h).DeviceGetPcieThroughput(counterType)
}
func (h GpuHandle) DeviceGetPerformanceState() (uint, error) {
	return _newFunction(h).DeviceGetPerformanceState()
}
func (h GpuHandle) DeviceGetPersistenceMode() (bool, error) {
	return _newFunction(h).DeviceGetPersistenceMode()
}
func (h GpuHandle) DeviceGetPowerManagementDefaultLimit() (uint, error) {
	return _newFunction(h).DeviceGetPowerManagementDefaultLimit()
}
func (h GpuHandle) DeviceGetPowerManagementLimit() (uint, error) {
	return _newFunction(h).DeviceGetPowerManagementLimit()
}
func (h GpuHandle) DeviceGetPowerManagementLimitConstraints() (min uint, max uint, err error) {
	return _newFunction(h).DeviceGetPowerManagementLimitConstraints()
}
func (h GpuHandle) DeviceGetPowerManagementMode() (bool, error) {
	return _newFunction(h).DeviceGetPowerManagementMode()
}
func (h GpuHandle) DeviceGetPowerState() (uint, error) {
	return _newFunction(h).DeviceGetPowerState()
}
func (h GpuHandle) DeviceGetPowerUsage() (uint, error) {
	return _newFunction(h).DeviceGetPowerUsage()
}
func (h GpuHandle) DeviceGetRetiredPages(cause PageRetirementCause) ([]uint64, error) {
	return _newFunction(h).DeviceGetRetiredPages(cause)
}
func (h GpuHandle) DeviceGetRetiredPagesPendingStatus() (bool, error) {
	return _newFunction(h).DeviceGetRetiredPagesPendingStatus()
}
func (h GpuHandle) DeviceGetSerial() (string, error) {
	return _newFunction(h).DeviceGetSerial()
}
func (h GpuHandle) DeviceGetSupportedClocksThrottleReasons() (uint64, error) {
	return _newFunction(h).DeviceGetSupportedClocksThrottleReasons()
}
func (h GpuHandle) DeviceGetSupportedGraphicsClocks(memoryClockMHz uint) ([]uint, error) {
	return _newFunction(h).DeviceGetSupportedGraphicsClocks(memoryClockMHz)
}
func (h GpuHandle) DeviceGetSupportedMemoryClocks() ([]uint, error) {
	return _newFunction(h).DeviceGetSupportedMemoryClocks()
}
func (h GpuHandle) DeviceGetTemperature() (uint, error) {
	return _newFunction(h).DeviceGetTemperature()
}
func (h GpuHandle) DeviceGetTemperatureThreshold(threshold TemperatureThresholds) (uint, error) {
	return _newFunction(h).DeviceGetTemperatureThreshold(threshold)
}
func (h GpuHandle) DeviceGetTopologyNearestGpus(level GpuTopologyLevel) ([]GpuHandle, error) {
	return _newFunction(h).DeviceGetTopologyNearestGpus(level)
}
func (h GpuHandle) DeviceGetTotalEccErrors(mt MemoryErrorType, et EccCounterType) (uint64, error) {
	return _newFunction(h).DeviceGetTotalEccErrors(mt, et)
}
func (h GpuHandle) DeviceGetUUID() (string, error) {
	return _newFunction(h).DeviceGetUUID()
}
func (h GpuHandle) DeviceGetUtilizationRates() (*Utilization, error) {
	return _newFunction(h).DeviceGetUtilizationRates()
}
func (h GpuHandle) DeviceGetVbiosVersion() (string, error) {
	return _newFunction(h).DeviceGetVbiosVersion()
}
func (h GpuHandle) DeviceGetViolationStatus(policy PerfPolicy) (*ViolationTime, error) {
	return _newFunction(h).DeviceGetViolationStatus(policy)
}
func (h GpuHandle) DeviceResetApplicationsClocks() error {
	return _newFunction(h).DeviceResetApplicationsClocks()
}
func (h GpuHandle) DeviceSetAutoBoostedClocksEnabled() (bool, error) {
	return _newFunction(h).DeviceSetAutoBoostedClocksEnabled()
}
func (h GpuHandle) DeviceSetCpuAffinity() error {
	return _newFunction(h).DeviceSetCpuAffinity()
}
func (h GpuHandle) DeviceSetDefaultAutoBoostedClocksEnabled(enabled bool) error {
	return _newFunction(h).DeviceSetDefaultAutoBoostedClocksEnabled(enabled)
}
func (h GpuHandle) DeviceValidateInforom() error {
	return _newFunction(h).DeviceValidateInforom()
}
func (h GpuHandle) DeviceClearEccErrorCounts(counterType EccCounterType) error {
	return _newFunction(h).DeviceClearEccErrorCounts(counterType)
}
func (h GpuHandle) DeviceSetAPIRestriction(api RestrictedAPI, isEnabled bool) error {
	return _newFunction(h).DeviceSetAPIRestriction(api, isEnabled)
}
func (h GpuHandle) DeviceSetApplicationsClocks(memClocksMHz, clocksMHz uint) error {
	return _newFunction(h).DeviceSetApplicationsClocks(memClocksMHz, clocksMHz)
}
func (h GpuHandle) DeviceSetComputeMode(mode ComputeMode) error {
	return _newFunction(h).DeviceSetComputeMode(mode)
}
func (h GpuHandle) DeviceSetEccMode(isEnabled bool) error {
	return _newFunction(h).DeviceSetEccMode(isEnabled)
}
func (h GpuHandle) DeviceSetGpuOperationMode(mode GpuOperationMode) error {
	return _newFunction(h).DeviceSetGpuOperationMode(mode)
}
func (h GpuHandle) DeviceSetPersistenceMode(isEnabled bool) error {
	return _newFunction(h).DeviceSetPersistenceMode(isEnabled)
}
func (h GpuHandle) DeviceSetPowerManagementLimit(limit uint) error {
	return _newFunction(h).DeviceSetPowerManagementLimit(limit)
}
func (h GpuHandle) DeviceGetAverageGPUUsage(since time.Duration) (uint, error) {
	return _newFunction(h).DeviceGetAverageGPUUsage(since)
}
func (h GpuHandle) DeviceGetProcessUtilization(maxProcess int, since time.Duration) ([]*ProcessUtilizationSample, error) {
	return _newFunction(h).DeviceGetProcessUtilization(maxProcess, since)
}
func (h GpuHandle) DeviceGetSupportedEventTypes() ([]EventType, error) {
	return _newFunction(h).DeviceGetSupportedEventTypes()
}
func (h GpuHandle) DeviceRegisterEvents(evtType EventType, set EventSet) error {
	return _newFunction(h).DeviceRegisterEvents(evtType, set)
}
