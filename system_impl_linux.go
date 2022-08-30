//+build: linux

type SystemImplement struct {
}

func (impl SystemImplement) Init() error {
	return Init()
}

func (impl SystemImplement) Shutdown() error {
	return Shutdown()
}

func (impl SystemImplement) DeviceGetCount() (uint, error) {
	return DeviceGetCount()
}

func (impl SystemImplement) SystemGetDriverVersion() (string, error) {
	return SystemGetDriverVersion()
}

func (impl SystemImplement) SystemGetNVMLVersion() (string, error) {
	return SystemGetNVMLVersion()
}

func (impl SystemImplement) SystemGetProcessName(pid uint) (string, error) {
	return SystemGetProcessName(pid)
}

func (impl SystemImplement) DeviceGetHandleByIndex(idx uint) (GpuHandle, error) {
	return DeviceGetHandleByIndex(idx)
}

func (impl SystemImplement) DeviceGetHandleByPciBusId(pciBusID string) (GpuHandle, error) {
	return DeviceGetHandleByPciBusId(pciBusID)
}

func (impl SystemImplement) DeviceGetHandleBySerial(serial string) (GpuHandle, error) {
	return DeviceGetHandleBySerial(serial)
}

func (impl SystemImplement) DeviceGetHandleByUUID(uuid string) (GpuHandle, error) {
	return DeviceGetHandleByUUID(uuid)
}

func (impl SystemImplement) DeviceGetTopologyCommonAncestor(h1, h2 GpuHandle) (GpuTopologyLevel, error) {
	return DeviceGetTopologyCommonAncestor(h1, h2)
}

func (impl SystemImplement) DeviceOnSameBoard(h1, h2 GpuHandle) (bool, error) {
	return DeviceOnSameBoard(h1, h2)
}

func (impl SystemImplement) SystemGetTopologyGpuSet(cpu int) ([]GpuHandle, error) {
	return SystemGetTopologyGpuSet(cpu)
}

func (impl SystemImplement) SystemGetHicVersion() ([]*HwbcEntry, error) {
	return SystemGetHicVersion()
}

func (impl SystemImplement) EventSetCreate() (*EventSet, error) {
	return EventSetCreate()
}
func (impl SystemImplement) EventSetFree(set *EventSet) error {
	return EventSetFree(set)
}

func (impl SystemImplement) EventSetWait(set EventSet, timeoutMS int) (*EventData, error) {
	return EventSetWait(set, timeoutMS)
}

