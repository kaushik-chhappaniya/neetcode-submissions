type ParkingSystem struct {
	spaces []int
}
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{spaces: []int{big, medium, small}}
}
func (this *ParkingSystem) AddCar(carType int) bool {
	this.spaces[carType-1]--
	return this.spaces[carType-1]>=0
}
