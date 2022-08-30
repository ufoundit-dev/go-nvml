//+build: linux
package nvml

import "time"

type GpuHandleImplement struct {
	handle GpuHandle
}

func NewGpuHandleImplement(h GpuHandle) *GpuHandleImplement {
	return &GpuHandleImplement{
		handle: h,
	}
}

func (impl *GpuHandleImplement) DeviceClearCpuAffinity() error {
	return impl.handle.DeviceClearCpuAffinity()
}

func (impl *GpuHandleImplement) DeviceGetAPIRestriction(apiType RestrictedAPI) (bool, error) {
	return impl.handle.DeviceGetAPIRestriction(apiType)
}

func (impl *GpuHandleImplement) DeviceGetApplicationsClock(clockType ClockType) (uint, error) {
	return impl.handle.DeviceGetApplicationsClock(clockType)
}
func (impl *GpuHandleImplement) DeviceGetAutoBoostedClocksEnabled() (curState bool, defaultState bool, err error) {
	return impl.handle.DeviceGetAutoBoostedClocksEnabled()
}
func (impl *GpuHandleImplement) DeviceGetBAR1MemoryInfo() (free uint64, used uint64, total uint64, err error) {
	return impl.handle.DeviceGetBAR1MemoryInfo()
}
func (impl *GpuHandleImplement) DeviceGetBoardId() (uint, error) {
	return impl.handle.DeviceGetBoardId()
}
func (impl *GpuHandleImplement) DeviceGetBrand() (BrandType, error) {
	return impl.handle.DeviceGetBrand()
}
func (impl *GpuHandleImplement) DeviceGetBridgeChipInfo() ([]*BridgeChipInfo, error) {
	return impl.handle.DeviceGetBridgeChipInfo()
}
func (impl *GpuHandleImplement) DeviceGetClockInfo(clockType ClockType) (uint, error) {
	return impl.handle.DeviceGetClockInfo(clockType)
}
func (impl *GpuHandleImplement) DeviceGetComputeMode() (ComputeMode, error) {
	return impl.handle.DeviceGetComputeMode()
}
func (impl *GpuHandleImplement) DeviceGetComputeRunningProcesses(size int) ([]*ProcessInfo, error) {
	return impl.handle.DeviceGetComputeRunningProcesses(size)
}
func (impl *GpuHandleImplement) DeviceGetCpuAffinity(size uint) ([]uint, error) {
	return impl.handle.DeviceGetCpuAffinity(size)
}
func (impl *GpuHandleImplement) DeviceGetCurrPcieLinkGeneration() (uint, error) {
	return impl.handle.DeviceGetCurrPcieLinkGeneration()
}
func (impl *GpuHandleImplement) DeviceGetCurrPcieLinkWidth() (uint, error) {
	return impl.handle.DeviceGetCurrPcieLinkWidth()
}
func (impl *GpuHandleImplement) DeviceGetCurrentClocksThrottleReasons() ([]ClocksThrottleReasons, error) {
	return impl.handle.DeviceGetCurrentClocksThrottleReasons()
}
func (impl *GpuHandleImplement) DeviceGetDecoderUtilization() (util uint, samplePeriod uint, err error) {
	return impl.handle.DeviceGetDecoderUtilization()
}
func (impl *GpuHandleImplement) DeviceGetDefaultApplicationsClock(clockType ClockType) (uint, error) {
	return impl.handle.DeviceGetDefaultApplicationsClock(clockType)
}
func (impl *GpuHandleImplement) DeviceGetDetailedEccErrors(mt MemoryErrorType, ec EccCounterType) (*EccErrorCounts, error) {
	return impl.handle.DeviceGetDetailedEccErrors(mt, ec)
}
func (impl *GpuHandleImplement) DeviceGetDisplayActive() (bool, error) {
	return impl.handle.DeviceGetDisplayActive()
}
func (impl *GpuHandleImplement) DeviceGetDisplayMode() (bool, error) {
	return impl.handle.DeviceGetDisplayMode()
}
func (impl *GpuHandleImplement) DeviceGetEccMode() (curMode bool, pendingMode bool, err error) {
	return impl.handle.DeviceGetEccMode()
}
func (impl *GpuHandleImplement) DeviceGetEncoderUtilization() (util uint, samplePeriod uint, err error) {
	return impl.handle.DeviceGetEncoderUtilization()
}
func (impl *GpuHandleImplement) DeviceGetEnforcedPowerLimit() (uint, error) {
	return impl.handle.DeviceGetEnforcedPowerLimit()
}
func (impl *GpuHandleImplement) DeviceGetFanSpeed() (uint, error) {
	return impl.handle.DeviceGetFanSpeed()
}
func (impl *GpuHandleImplement) DeviceGetGpuOperationMode() (curMode GpuOperationMode, pendingMode GpuOperationMode, err error) {
	return impl.handle.DeviceGetGpuOperationMode()
}
func (impl *GpuHandleImplement) GetGraphicsRunningProcesses(size int) ([]*ProcessInfo, error) {
	return impl.handle.GetGraphicsRunningProcesses(size)
}
func (impl *GpuHandleImplement) DeviceGetIndex() (uint, error) {
	return impl.handle.DeviceGetIndex()
}
func (impl *GpuHandleImplement) DeviceGetInforomConfigurationChecksum() (uint, error) {
	return impl.handle.DeviceGetInforomConfigurationChecksum()
}
func (impl *GpuHandleImplement) DeviceGetInforomImageVersion() (string, error) {
	return impl.handle.DeviceGetInforomImageVersion()
}
func (impl *GpuHandleImplement) DeviceGetInforomVersion(object InforomObject) (string, error) {
	return impl.handle.DeviceGetInforomVersion(object)
}
func (impl *GpuHandleImplement) DeviceGetMaxClockInfo(clockType ClockType) (uint, error) {
	return impl.handle.DeviceGetMaxClockInfo(clockType)
}
func (impl *GpuHandleImplement) DeviceGetMaxPcieLinkGeneration() (uint, error) {
	return impl.handle.DeviceGetMaxPcieLinkGeneration()
}
func (impl *GpuHandleImplement) DeviceGetMaxPcieLinkWidth() (uint, error) {
	return impl.handle.DeviceGetMaxPcieLinkWidth()
}
func (impl *GpuHandleImplement) DeviceGetMemoryErrorCounter(mt MemoryErrorType, ec EccCounterType, loc MemoryLocation) (uint64, error) {
	return impl.handle.DeviceGetMemoryErrorCounter(mt, ec, loc)
}
func (impl *GpuHandleImplement) DeviceGetMemoryInfo() (free uint64, used uint64, total uint64, err error) {
	return impl.handle.DeviceGetMemoryInfo()
}
func (impl *GpuHandleImplement) DeviceGetMinorNumber() (uint, error) {
	return impl.handle.DeviceGetMinorNumber()
}
func (impl *GpuHandleImplement) DeviceGetMultiGpuBoard() (uint, error) {
	return impl.handle.DeviceGetMultiGpuBoard()
}
func (impl *GpuHandleImplement) DeviceGetName() (string, error) {
	return impl.handle.DeviceGetName()
}
func (impl *GpuHandleImplement) DeviceGetPciInfo() (*PciInfo, error) {
	return impl.handle.DeviceGetPciInfo()
}
func (impl *GpuHandleImplement) DeviceGetPcieReplayCounter() (uint, error) {
	return impl.handle.DeviceGetPcieReplayCounter()
}
func (impl *GpuHandleImplement) DeviceGetPcieThroughput(counterType PcieUtilCounter) (uint, error) {
	return impl.handle.DeviceGetPcieThroughput(counterType)
}
func (impl *GpuHandleImplement) DeviceGetPerformanceState() (uint, error) {
	return impl.handle.DeviceGetPerformanceState()
}
func (impl *GpuHandleImplement) DeviceGetPersistenceMode() (bool, error) {
	return impl.handle.DeviceGetPersistenceMode()
}
func (impl *GpuHandleImplement) DeviceGetPowerManagementDefaultLimit() (uint, error) {
	return impl.handle.DeviceGetPowerManagementDefaultLimit()
}
func (impl *GpuHandleImplement) DeviceGetPowerManagementLimit() (uint, error) {
	return impl.handle.DeviceGetPowerManagementLimit()
}
func (impl *GpuHandleImplement) DeviceGetPowerManagementLimitConstraints() (min uint, max uint, err error) {
	return impl.handle.DeviceGetPowerManagementLimitConstraints()
}
func (impl *GpuHandleImplement) DeviceGetPowerManagementMode() (bool, error) {
	return impl.handle.DeviceGetPowerManagementMode()
}
func (impl *GpuHandleImplement) DeviceGetPowerState() (uint, error) {
	return impl.handle.DeviceGetPowerState()
}
func (impl *GpuHandleImplement) DeviceGetPowerUsage() (uint, error) {
	return impl.handle.DeviceGetPowerUsage()
}
func (impl *GpuHandleImplement) DeviceGetRetiredPages(cause PageRetirementCause) ([]uint64, error) {
	return impl.handle.DeviceGetRetiredPages(cause)
}
func (impl *GpuHandleImplement) DeviceGetRetiredPagesPendingStatus() (bool, error) {
	return impl.handle.DeviceGetRetiredPagesPendingStatus()
}
func (impl *GpuHandleImplement) DeviceGetSerial() (string, error) {
	return impl.handle.DeviceGetSerial()
}
func (impl *GpuHandleImplement) DeviceGetSupportedClocksThrottleReasons() (uint64, error) {
	return impl.handle.DeviceGetSupportedClocksThrottleReasons()
}
func (impl *GpuHandleImplement) DeviceGetSupportedGraphicsClocks(memoryClockMHz uint) ([]uint, error) {
	return impl.handle.DeviceGetSupportedGraphicsClocks(memoryClockMHz)
}
func (impl *GpuHandleImplement) DeviceGetSupportedMemoryClocks() ([]uint, error) {
	return impl.handle.DeviceGetSupportedMemoryClocks()
}
func (impl *GpuHandleImplement) DeviceGetTemperature() (uint, error) {
	return impl.handle.DeviceGetTemperature()
}
func (impl *GpuHandleImplement) DeviceGetTemperatureThreshold(threshold TemperatureThresholds) (uint, error) {
	return impl.handle.DeviceGetTemperatureThreshold(threshold)
}
func (impl *GpuHandleImplement) DeviceGetTopologyNearestGpus(level GpuTopologyLevel) ([]GpuHandle, error) {
	return impl.handle.DeviceGetTopologyNearestGpus(level)
}
func (impl *GpuHandleImplement) DeviceGetTotalEccErrors(mt MemoryErrorType, et EccCounterType) (uint64, error) {
	return impl.handle.DeviceGetTotalEccErrors(mt, et)
}
func (impl *GpuHandleImplement) DeviceGetUUID() (string, error) {
	return impl.handle.DeviceGetUUID()
}
func (impl *GpuHandleImplement) DeviceGetUtilizationRates() (*Utilization, error) {
	return impl.handle.DeviceGetUtilizationRates()
}
func (impl *GpuHandleImplement) DeviceGetVbiosVersion() (string, error) {
	return impl.handle.DeviceGetVbiosVersion()
}
func (impl *GpuHandleImplement) DeviceGetViolationStatus(policy PerfPolicy) (*ViolationTime, error) {
	return impl.handle.DeviceGetViolationStatus(policy)
}
func (impl *GpuHandleImplement) DeviceResetApplicationsClocks() error {
	return impl.handle.DeviceResetApplicationsClocks()
}
func (impl *GpuHandleImplement) DeviceSetAutoBoostedClocksEnabled() (bool, error) {
	return impl.handle.DeviceSetAutoBoostedClocksEnabled()
}
func (impl *GpuHandleImplement) DeviceSetCpuAffinity() error {
	return impl.handle.DeviceSetCpuAffinity()
}
func (impl *GpuHandleImplement) DeviceSetDefaultAutoBoostedClocksEnabled(enabled bool) error {
	return impl.handle.DeviceSetDefaultAutoBoostedClocksEnabled(enabled)
}
func (impl *GpuHandleImplement) DeviceValidateInforom() error {
	return impl.handle.DeviceValidateInforom()
}
func (impl *GpuHandleImplement) DeviceClearEccErrorCounts(counterType EccCounterType) error {
	return impl.handle.DeviceClearEccErrorCounts(counterType)
}
func (impl *GpuHandleImplement) DeviceSetAPIRestriction(api RestrictedAPI, isEnabled bool) error {
	return impl.handle.DeviceSetAPIRestriction(api, isEnabled)
}
func (impl *GpuHandleImplement) DeviceSetApplicationsClocks(memClocksMHz, clocksMHz uint) error {
	return impl.handle.DeviceSetApplicationsClocks(memClocksMHz, clocksMHz)
}
func (impl *GpuHandleImplement) DeviceSetComputeMode(mode ComputeMode) error {
	return impl.handle.DeviceSetComputeMode(mode)
}
func (impl *GpuHandleImplement) DeviceSetEccMode(isEnabled bool) error {
	return impl.handle.DeviceSetEccMode(isEnabled)
}
func (impl *GpuHandleImplement) DeviceSetGpuOperationMode(mode GpuOperationMode) error {
	return impl.handle.DeviceSetGpuOperationMode(mode)
}
func (impl *GpuHandleImplement) DeviceSetPersistenceMode(isEnabled bool) error {
	return impl.handle.DeviceSetPersistenceMode(isEnabled)
}
func (impl *GpuHandleImplement) DeviceSetPowerManagementLimit(limit uint) error {
	return impl.handle.DeviceSetPowerManagementLimit(limit)
}
func (impl *GpuHandleImplement) DeviceGetAverageGPUUsage(since time.Duration) (uint, error) {
	return impl.handle.DeviceGetAverageGPUUsage(since)
}
func (impl *GpuHandleImplement) DeviceGetProcessUtilization(maxProcess int, since time.Duration) ([]*ProcessUtilizationSample, error) {
	return impl.handle.DeviceGetProcessUtilization(maxProcess, since)
}
func (impl *GpuHandleImplement) DeviceGetSupportedEventTypes() ([]EventType, error) {
	return impl.handle.DeviceGetSupportedEventTypes()
}
func (impl *GpuHandleImplement) DeviceRegisterEvents(evtType EventType, set EventSet) error {
	return impl.handle.DeviceRegisterEvents(evtType, set)
}
