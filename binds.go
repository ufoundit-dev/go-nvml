//go:build !linux

package nvml

var (
	_sysImpl SystemInterface
)

func SetSystemImplement(impl SystemInterface) {
	_sysImpl = impl
}

func Init() error {
	return _sysImpl.Init()
}
func Shutdown() error {
	return _sysImpl.Shutdown()
}
func DeviceGetCount() (uint, error) {
	return _sysImpl.DeviceGetCount()
}
func SystemGetDriverVersion() (string, error) {
	return _sysImpl.SystemGetDriverVersion()
}
func SystemGetNVMLVersion() (string, error) {
	return _sysImpl.SystemGetNVMLVersion()
}
func SystemGetProcessName(pid uint) (string, error) {
	return _sysImpl.SystemGetProcessName(pid)
}
func DeviceGetHandleByIndex(idx uint) (GpuHandle, error) {
	return _sysImpl.DeviceGetHandleByIndex(idx)
}
func DeviceGetHandleByPciBusId(pciBusID string) (GpuHandle, error) {
	return _sysImpl.DeviceGetHandleByPciBusId(pciBusID)
}
func DeviceGetHandleBySerial(serial string) (GpuHandle, error) {
	return _sysImpl.DeviceGetHandleBySerial(serial)
}
func DeviceGetHandleByUUID(uuid string) (GpuHandle, error) {
	return _sysImpl.DeviceGetHandleByUUID(uuid)
}
func DeviceGetTopologyCommonAncestor(h1, h2 GpuHandle) (GpuTopologyLevel, error) {
	return _sysImpl.DeviceGetTopologyCommonAncestor(h1, h2)
}
func DeviceOnSameBoard(h1, h2 GpuHandle) (bool, error) {
	return _sysImpl.DeviceOnSameBoard(h1, h2)
}
func SystemGetTopologyGpuSet(cpu int) ([]GpuHandle, error) {
	return _sysImpl.SystemGetTopologyGpuSet(cpu)
}
func SystemGetHicVersion() ([]*HwbcEntry, error) {
	return _sysImpl.SystemGetHicVersion()
}
func EventSetCreate() (*EventSet, error) {
	return _sysImpl.EventSetCreate()
}
func EventSetFree(set *EventSet) error {
	return _sysImpl.EventSetFree(set)
}
func EventSetWait(set EventSet, timeoutMS int) (*EventData, error) {
	return _sysImpl.EventSetWait(set, timeoutMS)
}
