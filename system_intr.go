package nvml

type SystemInterface interface {
	Init() error
	Shutdown() error
	DeviceGetCount() (uint, error)
	SystemGetDriverVersion() (string, error)
	SystemGetNVMLVersion() (string, error)
	SystemGetProcessName(pid uint) (string, error)
	DeviceGetHandleByIndex(idx uint) (GpuHandle, error)
	DeviceGetHandleByPciBusId(pciBusID string) (GpuHandle, error)
	DeviceGetHandleBySerial(serial string) (GpuHandle, error)
	DeviceGetHandleByUUID(uuid string) (GpuHandle, error)
	DeviceGetTopologyCommonAncestor(h1, h2 GpuHandle) (GpuTopologyLevel, error)
	DeviceOnSameBoard(h1, h2 GpuHandle) (bool, error)
	SystemGetTopologyGpuSet(cpu int) ([]GpuHandle, error)
	SystemGetHicVersion() ([]*HwbcEntry, error)
	EventSetCreate() (*EventSet, error)
	EventSetFree(set *EventSet) error
	EventSetWait(set EventSet, timeoutMS int) (*EventData, error)
}
