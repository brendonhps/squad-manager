package Models


type Dev struct {
	name string `json: name`
	age int `json: age`
}

func InsertDev(dev Dev) (error) {

}

func SearchAllDevs() (devs []Dev, err error) {

}